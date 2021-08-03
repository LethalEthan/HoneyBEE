package nbt

//CreateCompoundTag - Create compound tag and sets it in NBTWriter so objects are written to it
func (NBTW *NBTWriter) AddCompoundTag(Name string) {
	TC := new(TCompound)
	TC.Name = Name
	TC.Value = make([]interface{}, 0, 8)
	TC.PreviousTag = NBTW.CurrentTag
	TC.PreviousTags = NBTW.CurrentTag.PreviousTags
	NBTW.CurrentTag.NextTags = append(NBTW.CurrentTag.NextTags, TC.NextTags...) //Finish me
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
func CreateCompoundTag(Name string, Capacity int) TCompound {
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
		NBTW.CurrentTag.AddTag(TEnd{})
		NBTW.CurrentTag.NumEntries = len(NBTW.CurrentTag.Value) //NumEntries is updated every time when AddTag is called, this is just for fool-proofing
		NBTW.CurrentTag.PreviousTag.AddTag(*NBTW.CurrentTag)
		NBTW.CurrentTag = NBTW.CurrentTag.PreviousTag
		NBTW.totalNumEntries += NBTW.CurrentTag.NumEntries
		return
	}
	NBTW.CurrentTag.NumEntries = len(NBTW.CurrentTag.Value)
	NBTW.CurrentTag.AddTag(TEnd{})
}

func (NBTW *NBTWriter) writeCompoundTag(Name string) {
	NBTW.writeTag(TagCompound, Name)
}

func (TC *TCompound) AddTag(val interface{}) {
	//TC.Index++
	TC.Value = append(TC.Value, val)
	// TC.NumEntries = len(TC.Value)
	// TC.TagsIDIndex[TC.NumEntries] = Name
	// TC.TagsIndex[Name] = TC.NumEntries
}

func (TC *TCompound) EndTag() {
	TC.Value = append(TC.Value, TEnd{})
}

func (TC *TCompound) AddMultipleTags(val []interface{}) {
	for _, v := range val {
		TC.Value = append(TC.Value, v)
	}
}
