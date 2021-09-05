package nbt

func CreateByteTag(Name string, Val byte) TByte {
	return TByte{Name, Val}
}

func (NBTW *NBTWriter) writeByte(Name string, val byte) {
	if Name != "" {
		NBTW.writeTag(TagByte, Name)
	}
	NBTW.AppendByteSlice([]byte{val})
}

func CreateByteArrayTag(Name string, Val []byte) TByteArray {
	return TByteArray{Name, Val, len(Val), 0}
}

func (NBTW *NBTWriter) writeByteArray(Name string, val []byte) {
	if Name != "" {
		NBTW.writeTag(TagByteArray, Name)
	}
	NBTW.writeInt("", int32(len(val)))
	NBTW.AppendByteSlice(val)
}
