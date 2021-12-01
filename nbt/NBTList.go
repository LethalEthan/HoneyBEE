package nbt

import (
	"HoneyBEE/utils"
	"math"
)

type List struct {
	Name    string
	tagtype byte
	value   []interface{}
}

func CreateListTag(Name string, Type byte) List {
	List := *new(List)
	List.Name = Name
	List.tagtype = Type
	List.value = make([]interface{}, 0, 16)
	return List
}

func CreateListTagWithCapacity(Name string, Type byte, Capacity int) List {
	List := *new(List)
	List.Name = Name
	List.tagtype = Type
	if Capacity > 0 {
		List.value = make([]interface{}, 0, 16)
	} else {
		List.value = make([]interface{}, 0, Capacity)
	}
	return List
}

func (L *List) AddTag(Value interface{}) {
	switch L.tagtype {
	case TagByte:
		if L.tagtype == TagByte {
			L.value = append(L.value, Value.(Byte))
		}
	case TagCompound:
		if L.tagtype == TagCompound {
			L.value = append(L.value, Value.(Compound))
		}
	}
}

func (NBTE *NBTEncoder) EncodeList(L List) {
	NBTE.data = append(NBTE.data, TagList)
	if L.Name != "" {
		NBTE.data = append(NBTE.data, utils.Int16ToByteArray(int16(len(L.Name)))...)
		NBTE.data = append(NBTE.data, L.Name...)
	}
	NBTE.data = append(NBTE.data, L.tagtype)
	NBTE.data = append(NBTE.data, utils.Int32ToByteArray(int32(len(L.value)))...)
	for i := range L.value {
		switch v := L.value[i].(type) {
		case End:
			NBTE.data = append(NBTE.data, 0)
		case Byte:
			NBTE.data = append(NBTE.data, v.Value)
		case Short:
			NBTE.data = append(NBTE.data, utils.UnsafeCastInt16ToBytes(v.Value)...) //append(NBTE.data, utils.Int16ToByteArray(v.Value)...)
		case Int:
			NBTE.data = append(NBTE.data, utils.UnsafeCastInt32ToBytes(v.Value)...) //utils.Int32ToByteArray(v.Value)...)
		case Long:
			NBTE.data = append(NBTE.data, utils.UnsafeCastInt64ToBytes(v.Value)...) //utils.Int64ToByteArray(v.Value)...)
		case Float:
			NBTE.data = append(NBTE.data, utils.UnsafeCastUint32ToBytes(math.Float32bits(v.Value))...) //utils.Uint32ToByteArray(math.Float32bits(v.Value))...)
		case Double:
			NBTE.data = append(NBTE.data, utils.UnsafeCastUint64ToBytes(math.Float64bits(v.Value))...) //utils.Uint64ToByteArray(math.Float64bits(v.Value))...)
		case ByteArray:
			NBTE.data = append(NBTE.data, v.Value...)
		case String:
			NBTE.data = append(NBTE.data, utils.UnsafeCastInt16ToBytes(int16(len(v.Value)))...) //utils.Int16ToByteArray(int16(len(v.Value)))...)
			NBTE.data = append(NBTE.data, v.Value...)
		case Compound:
			NBTE.encodeListCompound(&v)
		case IntArray:
			NBTE.data = append(NBTE.data, utils.UnsafeCastInt32ToBytes(int32(len(v.Value)))...) //utils.Int32ToByteArray(int32(len(v.Value)))...)
			NBTE.data = append(NBTE.data, utils.UnsafeCastInt32ArrayToBytes(v.Value)...)
			// for _, v := range v.Value {
			// 	NBTE.data = append(NBTE.data, utils.Int32ToByteArray(v)...)
			// }
		case LongArray:
			NBTE.data = append(NBTE.data, utils.UnsafeCastInt32ToBytes(int32(len(v.Value)))...) //append(NBTE.data, utils.Int32ToByteArray(int32(len(v.Value)))...)
			NBTE.data = append(NBTE.data, utils.UnsafeCastInt64ArrayToBytes(v.Value)...)
			// for _, v := range v.Value {
			// 	NBTE.data = append(NBTE.data, utils.Int64ToByteArray(v)...)
			// }
		}
	}
}

func (NBTE *NBTEncoder) encodeListCompound(C *Compound) {
	for _, v := range C.value {
		switch val := v.(type) {
		case End:
			NBTE.data = append(NBTE.data, TagEnd)
			return
		case Byte:
			NBTE.EncodeByte(val.Name, val.Value)
		case Short:
			NBTE.EncodeShort(val.Name, val.Value)
		case Int:
			NBTE.EncodeInt(val.Name, val.Value)
		case Long:
			NBTE.EncodeLong(val.Name, val.Value)
		case Float:
			NBTE.EncodeFloat(val.Name, val.Value)
		case Double:
			NBTE.EncodeDouble(val.Name, val.Value)
		case ByteArray:
			NBTE.EncodeByteArray(val.Name, val.Value)
		case String:
			NBTE.EncodeString(val.Name, val.Value)
		case List:
			NBTE.EncodeList(val)
		case Compound:
			NBTE.EncodeCompound(&val)
		case IntArray:
			NBTE.EncodeIntArray(val.Name, val.Value)
		case LongArray:
			NBTE.EncodeLongArray(val.Name, val.Value)
		}
	}
}
