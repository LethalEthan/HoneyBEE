package nbt

type TEnd struct{}

type TByte struct {
	Name  string
	Value byte
}

type TShort struct {
	Name  string
	Value int16
}

type TInt struct {
	Name  string
	Value int32
}

type TLong struct {
	Name  string
	Value int64
}

type TFloat struct {
	Name  string
	Value float32
}

type TDouble struct {
	Name  string
	Value float64
}

type TByteArray struct {
	Name   string
	Value  []byte
	length int
	Index  int
}

type TString struct {
	Name  string
	Value string
	//length int
}

type TList struct {
	Name   string
	Value  []interface{}
	Type   byte
	length int
}

type TCompound struct {
	Name         string
	Value        []interface{}
	TagsIndex    map[string]int
	TagsIDIndex  map[int]string
	NumEntries   int
	PreviousTag  *TCompound
	PreviousTags []*TCompound
	NextTags     []*CompoundTag
	Index        int
}

type TIntArray struct {
	Name   string
	Value  []int32
	length int
	Index  int
}

type TLongArray struct {
	Name   string
	Value  []int64
	length int
	Index  int
}

type TNone struct{}
