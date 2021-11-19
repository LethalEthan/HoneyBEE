package packet

type TagsArray struct {
	Length   int
	Type     Identifier
	TagArray []Tag
}

type Tag struct {
	Length  int32
	TagName Identifier
	Count   int
	Entries []int32
}

func CreateTag(Length int32, TagName Identifier, Entries []int32) Tag {
	// T := *new(Tag)
	// T.TagName = TagName
	// T.Entries = append(T.Entries, Entries...)
	return Tag{Length, TagName, len(Entries), Entries}
}

func CreateTagsArray(Type Identifier, TagArray []Tag) TagsArray {
	// TA := *new(TagsArray)
	// TA.Type = Type
	// TA.TagArray = TagArray
	// TA.Length = len(TagArray)
	return TagsArray{len(TagArray), Type, TagArray}
}

func (pr *PacketReader) ReadTagArray() []TagsArray {
	ArrayLength, _, err := pr.ReadVarInt()
	if err != nil || ArrayLength < 0 {
		panic(err)
	}
	TA := make([]TagsArray, ArrayLength)
	var Type Identifier
	for i := 0; i < int(ArrayLength); i++ {
		Type, err = pr.ReadIdentifier()
		if err != nil {
			panic(err)
		}
		//Log.Debug("Type: ", Type, "Seek: ", pr.GetSeeker())
		Tags := pr.ReadTags()
		TA[i] = CreateTagsArray(Type, Tags)
	}
	return TA
}

func (pr *PacketReader) ReadTags() []Tag {
	Length, _, err := pr.ReadVarInt()
	if err != nil {
		panic(err)
	}
	T := make([]Tag, Length)
	for i := 0; i < int(Length); i++ {
		TagName, err := pr.ReadIdentifier()
		if err != nil {
			panic(err)
		}
		Count, _, err := pr.ReadVarInt()
		if err != nil {
			panic(err)
		}
		Entries, err := pr.ReadVarIntArray(int(Count))
		if err != nil {
			panic(err)
		}
		T[i] = CreateTag(Length, TagName, Entries)
	}
	return T
}
