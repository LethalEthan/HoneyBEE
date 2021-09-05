package nbt

import "HoneyGO/utils"

func CreateStringTag(Name string, Val string) TString {
	return TString{Name, Val}
}

func (NBTW *NBTWriter) writeString(Name string, val string) {
	if Name != "" {
		NBTW.AppendByteSlice([]byte{TagString})
		NBTW.AppendByteSlice(utils.Int16ToByteArray(int16(len(Name))))
		NBTW.AppendByteSlice([]byte(Name))
	}
	if val != "" {
		NBTW.AppendByteSlice(utils.Int16ToByteArray(int16(len(val))))
		NBTW.AppendByteSlice([]byte(val))
	} else {
		if Name != "" {
			NBTW.AppendByteSlice([]byte{0, 0})
		}
	}
	return
}

func (TS TString) Encode() []byte {
	Data := make([]byte, 0, len(TS.Name)+len(TS.Value))
	if TS.Name != "" {
		Data = append(Data, []byte{TagString}...)
		Data = append(Data, utils.Int16ToByteArray(int16(len(TS.Name)))...)
		Data = append(Data, []byte(TS.Name)...)
	}
	if TS.Value != "" {
		Data = append(Data, utils.Int16ToByteArray(int16(len(TS.Value)))...)
		Data = append(Data, []byte(TS.Value)...)
	} else {
		if TS.Name == "" {
			Data = append(Data, []byte{0, 0}...)
		}
	}
	return Data
}
