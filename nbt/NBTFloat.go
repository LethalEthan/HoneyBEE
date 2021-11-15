package nbt

import (
	"encoding/binary"
	"math"
)

type Float struct {
	Name  string
	Value float32
}

func CreateFloatTag(Name string, Val float32) Float {
	return Float{Name, Val}
}

func (NBTE *NBTEncoder) EncodeFloat(Name string, val float32) {
	NBTE.EncodeTag(TagFloat, Name)
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, math.Float32bits(val))
	NBTE.data = append(NBTE.data, buf...)
}
