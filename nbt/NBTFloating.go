package nbt

import (
	"encoding/binary"
	"math"
)

func CreateFloatTag(Name string, Val float32) TFloat {
	return TFloat{Name, Val}
}

func (NBTW *NBTWriter) writeFloat(Name string, val float32) {
	NBTW.writeTag(TagFloat, Name)
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf[:], math.Float32bits(val))
	NBTW.AppendByteSlice(buf)
}

func CreateDoubleTag(Name string, Val float64) TDouble {
	return TDouble{Name, Val}
}

func (NBTW *NBTWriter) writeDouble(Name string, val float64) {
	if Name != "" {
		NBTW.writeTag(TagDouble, Name)
	}
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf[:], math.Float64bits(val))
	NBTW.AppendByteSlice(buf)
}
