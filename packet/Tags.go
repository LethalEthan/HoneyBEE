package packet

type TagsFormat struct {
	Type     Identifier
	TagArray []Tag
}

type Tag struct {
	Length  int32
	TagName []Identifier
	Count   int32
	Entries []int32
}
