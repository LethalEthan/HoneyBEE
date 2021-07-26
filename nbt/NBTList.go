package nbt

func CreateListTag(Name string, Type byte) TList {
	List := new(TList)
	List.Name = Name
	List.Type = Type
	List.Value = make([]interface{}, 0, 16)
	return *List
}

func (NBTW *NBTWriter) writeListTag(val string, Type byte, Length int) {
	NBTW.AppendByteSlice([]byte{TagList})
	NBTW.writeString("", val)
	NBTW.writeByte("", Type)
	NBTW.writeInt("", int32(Length))
}

func (L *TList) AddToList(val interface{}) {
	switch val.(type) {
	case TEnd:
		panic("Cannot add end tag to list")
	case TByte:
		L.Value = append(L.Value, val.(TByte))
	case TShort:
		L.Value = append(L.Value, val.(TShort))
	case TInt:
		L.Value = append(L.Value, val.(TInt))
	case TLong:
		L.Value = append(L.Value, val.(TLong))
	case TFloat:
		L.Value = append(L.Value, val.(TFloat))
	case TDouble:
		L.Value = append(L.Value, val.(TDouble))
	case TByteArray:
		L.Value = append(L.Value, val.(TByteArray))
	case TString:
		L.Value = append(L.Value, val.(TString))
	case TList:
		L.Value = append(L.Value, val.(TList))
	case TCompound:
		L.Value = append(L.Value, val.(TCompound))
	case TIntArray:
		L.Value = append(L.Value, val.(TIntArray))
	case TLongArray:
		L.Value = append(L.Value, val.(TLongArray))
	default:
		return
	}
	L.length++
}

func (NBTW *NBTWriter) writeList(Name string, Type byte, val []interface{}) {
	var ok bool
	NBTW.writeListTag(Name, Type, len(val))
	switch Type {
	case TagEnd:
		panic("Cannot list end tags")
	case TagByte:
		var B TByte
		for i := 0; i < len(val); i++ {
			B, ok = val[i].(TByte)
			if ok != true {
				panic("Type assertion failed for TList")
			}
			NBTW.writeByte("", B.Value)
		}
	case TagShort:
		var S TShort
		for i := 0; i < len(val); i++ {
			S, ok = val[i].(TShort)
			if ok != true {
				panic("TA failed for TShort")
			}
			NBTW.writeShort("", S.Value)
		}
	case TagInt:
		var I TInt
		for i := 0; i < len(val); i++ {
			I, ok = val[i].(TInt)
			if ok != true {
				panic("TA failed for TInt")
			}
			NBTW.writeInt("", I.Value)
		}
	case TagLong:
		var L TLong
		for i := 0; i < len(val); i++ {
			L, ok = val[i].(TLong)
			if ok != true {
				panic("TA failed for TLong")
			}
			NBTW.writeLong("", L.Value)
		}
	case TagFloat:
		var F TFloat
		for i := 0; i < len(val); i++ {
			F, ok = val[i].(TFloat)
			if ok != true {
				panic("TA failed for TFloat")
			}
			NBTW.writeFloat("", F.Value)
		}
	case TagDouble:
		var D TDouble
		for i := 0; i < len(val); i++ {
			D, ok = val[i].(TDouble)
			if ok != true {
				panic("TA failed for TDouble")
			}
			NBTW.writeDouble("", D.Value)
		}
	case TagByteArray:
		var BA TByteArray
		for i := 0; i < len(val); i++ {
			BA, ok = val[i].(TByteArray)
			if ok != true {
				panic("TA failed for TByteArray")
			}
			NBTW.writeByteArray("", BA.Value)
		}
	case TagString:
		var S TString
		for i := 0; i < len(val); i++ {
			S, ok = val[i].(TString)
			if ok != true {
				panic("TA failed for TString")
			}
			NBTW.writeString("", S.Value)
		}
	case TagList:
		var List TList
		for i := 0; i < len(val); i++ {
			List, ok = val[i].(TList)
			if ok != true {
				panic("TA failed for TList")
			}
			//NBTW.writeListTag(List.Name, List.Type, len(List.Value))
			NBTW.writeList(List.Name, List.Type, List.Value)
		}
	case TagCompound:
		var C TCompound
		for i := 0; i < len(val); i++ {
			C, ok = val[i].(TCompound)
			if ok != true {
				panic("TA failed for TCompound")
			}
			NBTW.traverseCompound(C)
		}
	case TagIntArray:
		var IA TIntArray
		for i := 0; i < len(val); i++ {
			IA, ok = val[i].(TIntArray)
			if ok != true {
				panic("TA failed for TIntArray")
			}
			NBTW.writeIntArray("", IA.Value)
		}
	case TagLongArray:
		var LA TLongArray
		for i := 0; i < len(val); i++ {
			LA, ok = val[i].(TLongArray)
			if ok != true {
				panic("TA failed for TLongArray")
			}
			NBTW.writeLongArray("", LA.Value)
		}
	}
}
