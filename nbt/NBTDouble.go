package nbt

import (
	"encoding/binary"
	"math"
)

type Double struct {
	Name  string
	Value float64
}

func CreateDoubleTag(Name string, Val float64) Double {
	return Double{Name, Val}
}

func (NBTE *NBTEncoder) EncodeDouble(Name string, val float64) {
	NBTE.EncodeTag(TagDouble, Name)
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, math.Float64bits(val))
	NBTE.data = append(NBTE.data, buf...)
}
