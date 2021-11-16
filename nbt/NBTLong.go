package nbt

import "HoneyBEE/utils"

type Long struct {
	Name  string
	Value int64
}

type LongArray struct {
	Name  string
	Value []int64
}

func CreateLongTag(Name string, Val int64) Long {
	return Long{Name, Val}
}
func CreateLongArrayTag(Name string, Val []int64) LongArray {
	return LongArray{Name, Val}
}

func (NBTE *NBTEncoder) EncodeLong(Name string, val int64) {
	NBTE.EncodeTag(TagLong, Name)
	NBTE.data = append(NBTE.data, utils.Int64ToByteArray(val)...)
}

func (NBTE *NBTEncoder) EncodeLongArray(Name string, Value []int64) {
	NBTE.EncodeTag(TagLongArray, Name)
	NBTE.data = append(NBTE.data, utils.Int32ToByteArray(int32(len(Value)))...)
	for i := range Value {
		NBTE.data = append(NBTE.data, utils.Int64ToByteArray(Value[i])...)
	}
}
