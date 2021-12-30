package nbt

import "HoneyBEE/utils"

// I am implementing list version of tags, since no names are needed they can be excluded and different logic can apply to it

type Byte struct {
	Name  string
	Value byte
}

type ListByte struct {
	Value byte
}

type ByteArray struct {
	Name  string
	Value []byte
}

type ListByteArray struct {
	Value []byte
}

func CreateByteTag(Name string, Val byte) Byte {
	return Byte{Name, Val}
}

func CreateListByteTag(Val byte) ListByte {
	return ListByte{Val}
}

func (NBTE *NBTEncoder) EncodeByte(Name string, Value byte) {
	NBTE.EncodeTag(TagByte, Name)
	NBTE.data = append(NBTE.data, Value)
}

func CreateByteArrayTag(Name string, Val []byte) ByteArray {
	return ByteArray{Name, Val}
}

func CreateListByteArrayTag(Val []byte) ListByteArray {
	return ListByteArray{Val}
}

func (NBTE *NBTEncoder) EncodeByteArray(Name string, Value []byte) {
	NBTE.EncodeTag(TagByteArray, Name)
	NBTE.data = append(NBTE.data, utils.Int32ToByteArray(int32(len(Value)))...)
	NBTE.data = append(NBTE.data, Value...)
}
