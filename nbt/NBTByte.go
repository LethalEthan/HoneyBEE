package nbt

import "HoneyBEE/utils"

type Byte struct {
	Name  string
	Value byte
}

type ByteArray struct {
	Name  string
	Value []byte
}

func CreateByteTag(Name string, Val byte) Byte {
	return Byte{Name, Val}
}

func (NBTE *NBTEncoder) EncodeByte(Name string, Value byte) {
	NBTE.EncodeTag(TagByte, Name)
	NBTE.data = append(NBTE.data, Value)
}

func CreateByteArrayTag(Name string, Val []byte) ByteArray {
	return ByteArray{Name, Val}
}

func (NBTE *NBTEncoder) EncodeByteArray(Name string, Value []byte) {
	NBTE.EncodeTag(TagByteArray, Name)
	NBTE.data = append(NBTE.data, utils.Int32ToByteArray(int32(len(Value)))...)
	NBTE.data = append(NBTE.data, Value...)
}
