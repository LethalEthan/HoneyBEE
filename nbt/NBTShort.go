package nbt

import "HoneyBEE/utils"

type Short struct {
	Name  string
	Value int16
}

func (NBTE *NBTEncoder) EncodeShort(Name string, Value int16) {
	NBTE.EncodeTag(TagShort, Name)
	NBTE.data = append(NBTE.data, utils.Int16ToByteArray(Value)...)
}
