package utils

import (
	"fmt"
	"runtime"
)

func PrintDebugStats() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Sys: ", mem.Sys)
	fmt.Println("mem.OtherSys:", mem.OtherSys)
	fmt.Println("mem.GCSys:", mem.GCSys)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.HeapIdle: ", mem.HeapIdle)
	fmt.Println("mem.HeapInuse: ", mem.HeapInuse)
	fmt.Println("mem.HeapObjects: ", mem.HeapObjects)
	fmt.Println("mem.HeapReleased: ", mem.HeapReleased)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("mem.NumForcedGC:", mem.NumForcedGC)
	fmt.Println("mem.StackInuse:", mem.StackInuse)
	fmt.Println("mem.StackSys:", mem.StackSys)
	fmt.Println("mem.D/E_GC:", mem.DebugGC, mem.EnableGC)
	fmt.Println("mem.CPUFraction:", mem.GCCPUFraction)
	fmt.Println("mem.Frees:", mem.Frees)
	fmt.Println("mem.Mallocs:", mem.Mallocs)
	fmt.Println("-----")
}
