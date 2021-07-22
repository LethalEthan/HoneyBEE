package utils

import (
	"fmt"
	"runtime"
)

func PrintDebugStats() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("mem.NumForcedGC:", mem.NumForcedGC)
	fmt.Println("-----")
}
