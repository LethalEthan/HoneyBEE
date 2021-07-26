package nbt

import "HoneyGO/utils"

func CreateStringTag(Name string, Val string) TString {
	return TString{Name, Val}
}

func (NBTW *NBTWriter) writeString(Name string, val string) {
	if Name != "" {
		NBTW.AppendByteSlice([]byte{TagString})
		NBTW.AppendByteSlice([]byte(utils.Int16ToByteArray(int16(len(Name)))))
		NBTW.AppendByteSlice([]byte(Name))
	}
	if val != "" {
		NBTW.AppendByteSlice([]byte(utils.Int16ToByteArray(int16(len(val)))))
		NBTW.AppendByteSlice([]byte(val))
	}
	return
}
