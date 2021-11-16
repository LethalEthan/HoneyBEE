package nbt

import "HoneyBEE/utils"

type Int struct {
	Name  string
	Value int32
}

type IntArray struct {
	Name  string
	Value []int32
}

func CreateIntTag(Name string, Val int32) Int {
	return Int{Name, Val}
}
func CreateIntArrayTag(Name string, Val []int32) IntArray {
	return IntArray{Name, Val}
}

func (NBTE *NBTEncoder) EncodeInt(Name string, Value int32) {
	NBTE.EncodeTag(TagInt, Name)
	NBTE.data = append(NBTE.data, utils.Int32ToByteArray(Value)...)
}

func (NBTE *NBTEncoder) EncodeIntArray(Name string, Value []int32) {
	NBTE.EncodeTag(TagInt, Name)
	NBTE.data = append(NBTE.data, utils.Int32ToByteArray(int32(len(Value)))...)
	for i := range Value {
		NBTE.data = append(NBTE.data, utils.Int32ToByteArray(Value[i])...)
	}
}
