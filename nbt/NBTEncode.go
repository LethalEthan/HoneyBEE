package nbt

func (NBTW *NBTWriter) Encode() {
	for _, v := range NBTW.Result {
		NBTW.writeCompoundTag(v.Name)
		NBTW.traverseCompound(*v)
	}
}

func (NBTW *NBTWriter) traverseCompound(TC TCompound) {
	for _, v := range TC.Value {
		switch v := v.(type) {
		case TEnd:
			NBTW.writeTag(TagEnd, "")
		case TByte:
			NBTW.writeByte(v.Name, v.Value)
		case TShort:
			NBTW.writeShort(v.Name, v.Value)
		case TInt:
			NBTW.writeInt(v.Name, v.Value)
		case TLong:
			NBTW.writeLong(v.Name, v.Value)
		case TFloat:
			NBTW.writeFloat(v.Name, v.Value)
		case TDouble:
			NBTW.writeDouble(v.Name, v.Value)
		case TByteArray:
			NBTW.writeByteArray(v.Name, v.Value)
		case TString:
			NBTW.writeString(v.Name, v.Value)
		case TList:
			NBTW.writeList(v.Name, v.Type, v.Value)
		case TCompound:
			NBTW.writeCompoundTag(v.Name)
			NBTW.traverseCompound(v)
		}
	}
}
