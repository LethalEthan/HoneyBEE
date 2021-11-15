package nbt

import (
	"fmt"
	"testing"
)

func BenchmarkAddTagByte(B *testing.B) {
	NBTE := CreateNBTEncoder()
	B.ReportAllocs()
	B.ResetTimer()
	for i := 0; i < B.N; i++ {
		NBTE.AddTag(CreateByteTag("testing", 128))
	}
}

func BenchmarkWriteByte(B *testing.B) {
	NBTE := CreateNBTEncoder()
	B.ResetTimer()
	NBTE.AddTag(Byte{"testing", 128})
	for i := 0; i < B.N; i++ {
		NBTE.Encode()
	}
}

func TestWriteByte(T *testing.T) {
	NBTE := CreateNBTEncoder()
	NBTE.AddTag(Byte{"testing", 128}) //CreateByteTag("testing", 128))
	NBTE.EndCompoundTag()
	NBTE.Encode()
	fmt.Println(NBTE.data)
	fmt.Println("NBTEVAL: ")
	fmt.Print(NBTE.rootCompound.value...)
}
