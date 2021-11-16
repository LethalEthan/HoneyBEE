package nbt

import "HoneyBEE/utils"

type List struct {
	Name    string
	tagtype byte
	value   []interface{}
}

func CreateListTag(Name string, Type byte) *List {
	List := new(List)
	List.Name = Name
	List.tagtype = Type
	List.value = make([]interface{}, 0, 16)
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
	//NBTE.EncodeString("", L.Name)
	NBTE.data = append(NBTE.data, L.tagtype)
	NBTE.data = append(NBTE.data, utils.Int32ToByteArray(int32(len(L.value)))...)
	for i := range L.value {
		switch v := L.value[i].(type) {
		case Byte:
			NBTE.data = append(NBTE.data, v.Value)
		case Compound:
			v.name = ""
			NBTE.encodeListCompound(&v)
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
