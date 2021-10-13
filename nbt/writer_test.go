package nbt

import "testing"

func BenchmarkNBTWriter(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NBTW := CreateNBTWriter("hello world")
		NBTW.AddTag(CreateStringTag("hello", "bruh"))
		NBTW.AddTag(CreateStringTag("bruh", "hello"))
		NBTW.EndCompoundTag()
		NBTW.Encode()
	}
}

func BenchmarkNBTWriterEncode(b *testing.B) {
	NBTW := CreateNBTWriter("hello world")
	NBTW.AddTag(CreateStringTag("hello", "bruh"))
	NBTW.AddTag(CreateStringTag("bruh", "hello"))
	NBTW.EndCompoundTag()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NBTW.Encode()
	}
}
