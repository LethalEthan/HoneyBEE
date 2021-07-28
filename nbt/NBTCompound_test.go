package nbt

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestCreateCompoundTag(T *testing.T) {
	NBTW := CreateNBTWriter("test")
	NBTW.AddCompoundTag("testing")
	NBTW.EndCompoundTag()
	NBTW.Encode()
}

func TestCreateCompoundTagObject(T *testing.T) {
	NBTW := CreateNBTWriter("kacper")
	TC := CreateCompoundTagObject("naomi", 8)
	TC.AddTag(CreateStringTag("felenov", "back-door plugins"))
	TC.AddTag(CreateStringTag("fahlur", "waypoint plugin"))
	TC.AddTag(CreateStringTag("lynxplay", "lord and saviour, the coding god"))
	err := NBTW.AddTag(TC)
	if err != nil {
		T.Error(err)
	}
	TC2 := CreateCompoundTagObject("Sloker", 0)
	TC2.AddTag(TC)
}

func BenchmarkAdd(B *testing.B) {
	var j uint64
	for i := 0; i < B.N; i++ {
		j++
	}
	fmt.Print(j)
}

var H uint64

func BenchmarkAddAtomicContending(B *testing.B) {
	var j uint64
	go burh2(B.N)
	for i := 0; i < B.N; i++ {
		atomic.AddUint64(&H, 1)
	}
	fmt.Print(j)
}

var J uint64
var M sync.Mutex

func BenchmarkAddMutexContending(B *testing.B) {
	var j uint64
	//var m sync.Mutex
	go bruh(B.N)
	for i := 0; i < B.N; i++ {
		M.Lock()
		J++
		M.Unlock()
	}
	fmt.Print(j)
}

func bruh(l int) {
	for i := 0; i < l*2; i++ {
		M.Lock()
		J++
		M.Unlock()
	}
}

func burh2(l int) {
	for i := 0; i < l*2; i++ {
		atomic.AddUint64(&H, 1)
	}
}

func BenchmarkAddAtomic(B *testing.B) {
	for i := 0; i < B.N; i++ {
		atomic.AddUint64(&H, 1)
	}
}

func BenchmarkAddMutex(B *testing.B) {
	for i := 0; i < B.N; i++ {
		M.Lock()
		J++
		M.Unlock()
	}
}
