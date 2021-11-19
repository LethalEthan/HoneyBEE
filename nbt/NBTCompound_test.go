package nbt

import (
	"fmt"
	"testing"
)

func TestCompoundTag(T *testing.T) {
	NBTE := CreateNBTEncoder()
	NBTE.AddCompoundTag("TEST")
	NBTE.AddTag(Byte{"", 128})
	NBTE.EndCompoundTag()
	NBTE.EndCompoundTag()
	NBTE.Encode()
	fmt.Println(NBTE.data)
	fmt.Println("NBTEVAL: ")
	fmt.Print(NBTE.rootCompound.value)
}

func TestCompoundTagNoEnding(T *testing.T) {
	NBTE := CreateNBTEncoder()
	NBTE.AddCompoundTag("TEST")
	NBTE.AddTag(Byte{"", 128})
	NBTE.Encode()
	fmt.Println(NBTE.data)
	fmt.Println("NBTEVAL: ")
	fmt.Print(NBTE.rootCompound.value)
}

func BenchmarkCompoundTag(B *testing.B) {
	B.ReportAllocs()
	B.ResetTimer()
	NBTE := CreateNBTEncoder()
	for i := 0; i < B.N; i++ {
		NBTE.AddCompoundTag("TEST")
		NBTE.AddTag(Byte{"", 127})
		NBTE.EndCompoundTag()
		NBTE.EndCompoundTag()
		NBTE.Encode()
		NBTE.Reset()
	}
}
