package nbt

import (
	"testing"
)

func TestAddCompoundTag(T *testing.T) {
	NBTW := CreateNBTWriter("test")
	NBTW.AddCompoundTag("testing")
	NBTW.EndCompoundTag()
	NBTW.Encode()
}

func TestCreateCompoundTag(T *testing.T) {
	NBTW := CreateNBTWriter("kacper")
	TC := CreateCompoundTag("naomi", 8)
	TC.AddTag(CreateStringTag("felenov", "back-door plugins"))
	TC.AddTag(CreateStringTag("fahlur", "waypoint plugin"))
	TC.AddTag(CreateStringTag("lynxplay", "lord and saviour, the coding god"))
	err := NBTW.AddTag(TC)
	if err != nil {
		T.Error(err)
	}
	TC2 := CreateCompoundTag("Sloker", 0)
	TC2.AddTag(TC)
}
