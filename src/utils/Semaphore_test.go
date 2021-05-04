package utils

import "testing"

var S *Semaphore
var Iterations int = 100
var Size uint32 = 10

//Fix me

func BenchmarkFlushAndSetSemaphore(b *testing.B) {
	S = MakeSemaphore()
	for i := 0; i < Iterations; i++ {
		S.FlushAndSetSemaphore(i)
	}
	return
}

func BenchmarkGetData(b *testing.B) {
	var I interface{}
	S = MakeSemaphore()
	for i := 0; i < Iterations; i++ {
		I = S.GetData()
		_ = I
	}
	return
}

func BenchmarkFlushAndSetSemaphoreNEW(b *testing.B) {
	S = MakeSemaphore()
	for i := 0; i < Iterations; i++ {
		S.FlushAndSetSemaphoreNEW(i)
	}
	return
}

func MakeSemaphore() *Semaphore {
	S = CreateSemaphore(Size)
	S.Start()
	S.FlushAndSetSemaphore(19)
	return S
}
