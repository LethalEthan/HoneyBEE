package console

import (
	"HoneyBEE/config"
	"HoneyBEE/packet"
	"HoneyBEE/utils"
	"fmt"
	"runtime"
)

//DRECOVER - recovers panics and spits out additional info before exiting
func DRECOVER() {
	if r := recover(); r != nil {
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		fmt.Println("Server encountered panic! reason: ", r)
		fmt.Println("Printing Debug, please create an issue and send this!")
		fmt.Println("----------------------------------------")
		fmt.Println("Server                                  ")
		fmt.Println("----------------------------------------")
		fmt.Println("HoneyBEEVersion: ", utils.GetVersionString(), "BVersion: ", utils.BuildVersion, "FH: ", Hash())
		fmt.Println("Config:", config.GConfig)
		fmt.Println("----------------------------------------")
		fmt.Println("Network")
		fmt.Println("----------------------------------------")
		fmt.Println("Auth: ", packet.Hash())
		fmt.Println("----------------------------------------")
		fmt.Println("Other")
		fmt.Println("----------------------------------------")
		fmt.Println("NumCPU: ", runtime.NumCPU())
		fmt.Println("NumGORoutines: ", runtime.NumGoroutine())
		fmt.Println("Arch: ", runtime.GOARCH)
		fmt.Println("OS: ", runtime.GOOS)
		fmt.Println("GOVer: ", runtime.Version())
		fmt.Println("Printed Debug , please create an issue and send this!")
		fmt.Println("")
		utils.PrintDebugStats()
		panic(r)
	}
}
