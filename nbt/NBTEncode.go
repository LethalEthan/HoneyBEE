package nbt

func (NBTW *NBTWriter) Encode() {
	for _, v := range NBTW.Result {
		NBTW.writeCompoundTag(v.Name)
		NBTW.traverseCompound(*v)
	}
}

func (NBTW *NBTWriter) traverseCompound(TC TCompound) {
	for _, v := range TC.Value {
		switch v.(type) {
		case TEnd:
			NBTW.writeTag(TagEnd, "")
			break
		case TByte:
			NBTW.writeByte(v.(TByte).Name, v.(TByte).Value)
		case TShort:
			NBTW.writeShort(v.(TShort).Name, v.(TShort).Value)
		case TInt:
			NBTW.writeInt(v.(TInt).Name, v.(TInt).Value)
		case TLong:
			NBTW.writeLong(v.(TLong).Name, v.(TLong).Value)
		case TFloat:
			NBTW.writeFloat(v.(TFloat).Name, v.(TFloat).Value)
		case TDouble:
			NBTW.writeDouble(v.(TDouble).Name, v.(TDouble).Value)
		case TByteArray:
			NBTW.writeByteArray(v.(TByteArray).Name, v.(TByteArray).Value)
		case TString:
			NBTW.writeString(v.(TString).Name, v.(TString).Value)
		case TList:
			NBTW.writeList(v.(TList).Name, v.(TList).Type, v.(TList).Value)
		case TCompound:
			//NBTW.writeTag(TagCompound, v.(TCompound).Name)
			NBTW.writeCompoundTag(v.(TCompound).Name)
			NBTW.traverseCompound(v.(TCompound))
		}
	}
}
