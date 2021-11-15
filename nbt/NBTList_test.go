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
	NBTE.AddTag(*List)
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
	NBTE.AddTag(*List)
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

}
