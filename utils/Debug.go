package utils

import (
	"fmt"
	"runtime"
)

func PrintDebugStats() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Sys: ", mem.Sys/1000000)
	fmt.Println("mem.OtherSys:", mem.OtherSys/1000000)
	fmt.Println("mem.GCSys:", mem.GCSys/1000000)
	fmt.Println("mem.Alloc:", mem.Alloc/1000000)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc/1000000)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc/100000)
	fmt.Println("mem.HeapIdle: ", mem.HeapIdle/1000000)
	fmt.Println("mem.HeapInuse: ", mem.HeapInuse/1000000)
	fmt.Println("mem.HeapObjects: ", mem.HeapObjects)
	fmt.Println("mem.HeapReleased: ", mem.HeapReleased/1000000)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("mem.NumForcedGC:", mem.NumForcedGC)
	fmt.Println("mem.StackInuse:", mem.StackInuse/1000000)
	fmt.Println("mem.StackSys:", mem.StackSys/1000000)
	fmt.Println("mem.D/E_GC:", mem.DebugGC, mem.EnableGC)
	fmt.Println("mem.CPUFraction:", mem.GCCPUFraction)
	fmt.Println("mem.Frees:", mem.Frees)
	fmt.Println("mem.Mallocs:", mem.Mallocs)
	fmt.Println("mem.NextGC:", mem.NextGC)
	fmt.Println("mem.LastGC:", mem.Mallocs)
	fmt.Println("mem.PauseTotalns:", mem.PauseTotalNs)
	fmt.Println("mem.MCacheSys:", mem.MCacheSys/1000000)
	fmt.Println("mem.MCacheInuse:", mem.MCacheInuse/1000000)
	fmt.Println("-----")
}
