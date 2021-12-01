package nbt

import "HoneyBEE/utils"

type String struct {
	Name  string
	Value string
}

func CreateStringTag(Name string, Value string) String {
	return String{Name, Value}
}

func (NBTE *NBTEncoder) EncodeString(Name string, Value string) {
	// NBTE.EncodeTag(TagString, Name)
	if Name != "" {
		NBTE.data = append(NBTE.data, TagString)
		NBTE.data = append(NBTE.data, utils.Int16ToByteArray(int16(len(Name)))...)
		NBTE.data = append(NBTE.data, Name...)
	} else {
		NBTE.data = append(NBTE.data, []byte{0, 0}...)
	}
	if Value != "" {
		NBTE.data = append(NBTE.data, utils.Int16ToByteArray(int16(len(Value)))...)
		NBTE.data = append(NBTE.data, []byte(Value)...)
	} else {
		NBTE.data = append(NBTE.data, []byte{0, 0}...)
	}
}
