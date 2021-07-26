package nbt

import "HoneyGO/utils"

///
///Int16/Short
///

func CreateShortTag(Name string, Val int16) TShort {
	return TShort{Name, Val}
}

func (NBTW *NBTWriter) writeShort(Name string, val int16) {
	NBTW.writeTag(TagShort, Name)
	NBTW.AppendByteSlice(utils.Int16ToByteArray(val))
}

///
///Int32/Integer
///

func CreateIntTag(Name string, Val int32) TInt {
	return TInt{Name, Val}
}

func CreateIntArrayTag(Name string, Val []int32) TIntArray {
	return TIntArray{Name, Val, len(Val), 0}
}

func (NBTW *NBTWriter) writeInt(Name string, val int32) {
	if Name != "" {
		NBTW.writeTag(TagInt, Name)
	}
	NBTW.AppendByteSlice(utils.Int32ToByteArray(val))
}

func (NBTW *NBTWriter) writeIntArray(Name string, val []int32) {
	if Name != "" {
		NBTW.writeTag(TagIntArray, Name)
	}
	NBTW.writeInt("", int32(len(val)))
	for i := 0; i < len(val); i++ {
		NBTW.AppendByteSlice(utils.Int32ToByteArray(val[i]))
	}
}

///
///Int64/Long
///

func CreateLongTag(Name string, Val int64) TLong {
	return TLong{Name, Val}
}
func CreateLongArrayTag(Name string, Val []int64) TLongArray {
	return TLongArray{Name, Val, len(Val), 0}
}

func (NBTW *NBTWriter) writeLong(Name string, val int64) {
	if Name != "" {
		NBTW.writeTag(TagLong, Name)
	}
	NBTW.AppendByteSlice(utils.Int64ToByteArray(val))
}

func (NBTW *NBTWriter) writeLongArray(Name string, val []int64) {
	if Name != "" {
		NBTW.writeTag(TagLongArray, Name)
	}
	NBTW.writeInt("", int32(len(val)))
	for i := 0; i < len(val); i++ {
		NBTW.AppendByteSlice(utils.Int64ToByteArray(val[i]))
	}
}
