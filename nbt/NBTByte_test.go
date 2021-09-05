package nbt

import (
	"testing"
)

func BenchmarkWriteByte(B *testing.B) {
	NBTW := CreateNBTWriter("bruh")
	B.ReportAllocs()
	B.ResetTimer()
	for i := 0; i < B.N; i++ {
		NBTW.AddTag(CreateByteTag("testing", 128))
	}
}

func TestReadByte(T *testing.T) {
	NBTW := CreateNBTWriter("test")
	NBTW.AddTag(CreateByteTag("testing", 128))
	NBTW.EndCompoundTag()
	NBTW.Encode()
	//
	NBTR, err := CreateNBTReader(NBTW.Data)
	if err != nil {
		panic(err)
	}
	Type, Name, err := NBTR.readTagType()
	if err != nil {
		panic(err)
	}
	if Type != TagCompound {
		T.Error(NBTType)
	}
	if Name != "test" {
		T.Error(NBTName)
	}
	//
	Type, Name, err = NBTR.readTagType()
	if err != nil {
		panic(err)
	}
	if Type != TagByte {
		T.Error(NBTType)
	}
	if Name != "testing" {
		T.Error(NBTName)
	}
	Byte, err := NBTR.readByte()
	if err != nil {
		panic(err)
	}
	if Byte != 128 {
		T.Error(NBTValue)
	}
}

func TestReadByteArray(T *testing.T) {
	NBTW := CreateNBTWriter("test")
	NBTW.AddTag(CreateByteArrayTag("testing", []byte{128, 120, 168, 120, 100, 69, 69, 69, 69, 69, 128}))
	NBTW.EndCompoundTag()
	NBTW.Encode()
	//TagCompound
	NBTR, err := CreateNBTReader(NBTW.Data)
	if err != nil {
		panic(err)
	}
	Type, Name, err := NBTR.readTagType()
	if err != nil {
		panic(err)
	}
	if Type != TagCompound {
		T.Error(NBTType)
	}
	if Name != "test" {
		T.Error(NBTName)
	}
	//TagByteArray
	Type, Name, err = NBTR.readTagType()
	if err != nil {
		panic(err)
	}
	if Type != TagByteArray {
		T.Error(NBTType)
	}
	if Name != "testing" {
		T.Error(NBTName)
	}
	Type, Name, err = NBTR.readTagType()
	if err != nil {
		panic(err)
	}
	if Type != TagEnd {
		T.Error(NBTType)
	}
	if Name != "" {
		T.Error(NBTName)
	}
}
