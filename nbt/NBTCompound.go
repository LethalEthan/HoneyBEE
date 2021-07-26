package nbt

//CreateCompoundTag - Create compound tag and sets it in NBTWriter so objects are written to it
func (NBTW *NBTWriter) CreateCompoundTag(Name string) {
	TC := new(TCompound)
	TC.Name = Name
	TC.Value = make([]interface{}, 0, 8)
	//TC.init()
	TC.PreviousTag = NBTW.CurrentTag
	// NBTW.CurrentTag.TagsIndex[Name] = NBTW.CurrentTag.Index
	// NBTW.CurrentTag.TagsIDIndex[NBTW.CurrentTag.Index] = Name
	NBTW.currentTagIndex++
	//NBTW.CurrentTag.Index++
	NBTW.CurrentTag = TC
	NBTW.numCompounds++
}

/*CreateCompoundTagObject - Creates Compound tag and returns it, this does not change anything in writer/reader
compound tag will have to be added to NBTWriter with NBTW.AddTag(TCompound)
Capacity is for the []interface{} slice that stores the objects, the capacity is how many objects can be appended without needing a new slice to be re-allocated*/
func CreateCompoundTagObject(Name string, Capacity int) TCompound {
	TC := new(TCompound)
	TC.Name = Name
	if Capacity > 0 {
		TC.Value = make([]interface{}, 0, Capacity)
	} else {
		TC.Value = make([]interface{}, 0, 8)
	}
	return *TC
}

func (NBTW *NBTWriter) EndCompoundTag() {
	if NBTW.CurrentTag.PreviousTag != nil {
		NBTW.CurrentTag.AddTag("", TEnd{})
		NBTW.CurrentTag.NumEntries = len(NBTW.CurrentTag.Value) //NumEntries is updated every time when AddTag is called, this is just for fool-proofing
		NBTW.CurrentTag.PreviousTag.AddTag(NBTW.CurrentTag.Name, *NBTW.CurrentTag)
		NBTW.CurrentTag = NBTW.CurrentTag.PreviousTag
		return
	}
	NBTW.CurrentTag.NumEntries = len(NBTW.CurrentTag.Value)
	NBTW.CurrentTag.AddTag("", TEnd{})
}

func (NBTW *NBTWriter) writeCompoundTag(Name string) {
	NBTW.writeTag(TagCompound, Name)
}

// func (TC *TCompound) init() {
// 	// TC.TagsIndex = make(map[string]int)
// 	// TC.TagsIDIndex = make(map[int]string)
// 	TC.Value = make([]interface{}, 0, 16) //Make length 0 and capacity 16 helps with lowering allocations until it reaches 16
// }

func (TC *TCompound) AddTag(Name string, val interface{}) {
	//TC.Index++
	TC.Value = append(TC.Value, val)
	// TC.NumEntries = len(TC.Value)
	// TC.TagsIDIndex[TC.NumEntries] = Name
	// TC.TagsIndex[Name] = TC.NumEntries
}

func (TC *TCompound) AddMultipleTags(val []interface{}) {
	for _, v := range val {
		TC.Value = append(TC.Value, v)
	}
}

// func (TC *TCompound) FindTagByName(Name string) (interface{}, error) {
// 	T, ok := TC.TagsIndex[Name]
// 	if ok == true {
// 		return &TC.Value[T], nil
// 	}
// 	return nil, NBTRead404
// }
