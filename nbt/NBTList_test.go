package nbt

import (
	"fmt"
	"os"
	"testing"
)

func TestList(T *testing.T) {
	NBTE := CreateNBTEncoder()
	List := CreateListTag("TESTIES", TagCompound)
	C := CreateCompoundTag("")
	C.AddTag(CreateByteTag("PENIS", 127))
	C.EndTag()
	List.AddTag(C)
	NBTE.AddTag(List)
	NBTE.EndCompoundTag()
	// NBTE.EncodeList(List)
	NBTE.Encode()
	fmt.Printf("D: %x ", NBTE.data)
	fmt.Println(NBTE.rootCompound.value...)
}

func TestList2(T *testing.T) {
	NBTE := CreateNBTEncoder()
	List := CreateListTag("TESTIES", TagCompound)
	C := CreateCompoundTag("")
	C.AddTag(CreateByteTag("PENIS", 127))
	C2 := CreateCompoundTag("Hello")
	C2.AddTag(CreateByteTag("Bruh", 127))
	C2.EndTag()
	C.AddTag(C2)
	C.EndTag()
	List.AddTag(C)
	NBTE.AddTag(List)
	NBTE.EndCompoundTag()
	// NBTE.EncodeList(List)
	NBTE.Encode()
	fmt.Printf("D: %x ", NBTE.data)
	fmt.Println(NBTE.rootCompound.value...)
	f, err := os.Create("TEST.nbt")
	if err != nil {
		panic(err)
	}
	f.Write(NBTE.data)
	f.Close()
}

var TL3KnownGood = []byte{0x0A, 0x00, 0x00, 0x09, 0x00, 0x08, 0x74, 0x65, 0x73, 0x74, 0x6C, 0x69, 0x73, 0x74, 0x0A, 0x00, 0x00, 0x00, 0x02, 0x01, 0x00, 0x08, 0x62, 0x79, 0x74, 0x65, 0x68, 0x65, 0x72, 0x65, 0x78, 0x0A, 0x00, 0x04, 0x6C, 0x6D, 0x61, 0x6F, 0x00, 0x00, 0x01, 0x00, 0x0f, 0x61, 0x20, 0x62, 0x79, 0x74, 0x65, 0x20, 0x68, 0x65, 0x72, 0x65, 0x20, 0x6f, 0x6d, 0x67, 0x08, 0, 0}

func TestList3(T *testing.T) {
	NBTE := CreateNBTEncoder()
	List := CreateListTag("testlist", TagCompound)
	TC := CreateCompoundTag("")
	TC.AddTag(CreateByteTag("bytehere", 120))
	TC2 := CreateCompoundTag("lmao")
	TC2.EndTag()
	TC.AddTag(TC2)
	TC.EndTag()
	//
	TC3 := CreateCompoundTag("")
	TC3.AddTag(CreateByteTag("a byte here omg", 8))
	TC3.EndTag()
	List.AddTag(TC)
	List.AddTag(TC3)
	NBTE.AddTag(List)
	NBTE.EndCompoundTag()
	Log.Debug("Encoded: ", NBTE.Encode())
	fmt.Printf("D: %x ", NBTE.data)
	fmt.Println(NBTE.rootCompound.value...)
	f, err := os.Create("TEST2.nbt")
	if err != nil {
		panic(err)
	}
	f.Write(NBTE.Encode())
	f.Close()
	for i := range NBTE.data {
		if NBTE.data[i] != TL3KnownGood[i] {
			T.Error(NBTValue)
		}
	}
}
