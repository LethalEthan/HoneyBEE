package console

import (
	"HoneyGO/config"
	"HoneyGO/npacket"
	"HoneyGO/player"
	"HoneyGO/utils"
	"fmt"
	"runtime"
)

///
///This recovers panics and spits out additional info before exiting
///

//DRECOVER - Recovery -- TBD close server go routines
func DRECOVER() {
	if r := recover(); r != nil {
		Panicked = true //TBD: link to event system to stop everything
		//go SetRun(false)
		//go server.SetRun(false)
		go func() { //Lock mutexes in case something carries on upon recovery
			// server.RunMutex.Lock()
			// server.PlayerMapMutex.Lock()
			// server.ConnPlayerMutex.Lock()
			// server.ClientConnectionMutex.Lock()
			// server.StatusMutex.Lock()
			// server.PlayerConnMutex.Lock()
			// server.RunMutex.Lock()
			player.OnlinePlayerMutex.Lock()
			player.PlayerEntityMutex.Lock()
			player.PlayerObjectMutex.Lock()
			//server.Run = false //reset as false in case the go routine did not
			//server.GCPShutdown <- true
		}()
		//go server.StatusSemaphore.StopSemaphore() //TODO: Make all semaphores be able to be stopped
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		fmt.Println("Server encountered panic! reason: ", r)
		fmt.Println("Printing Debug, please create an issue and send this!")
		fmt.Println("----------------------------------------")
		fmt.Println("Server Maps and states (package server)")
		fmt.Println("----------------------------------------")
		// fmt.Println("Server run state:", server.Run, "Mutex: ", server.RunMutex)
		fmt.Println("HoneyGOVersion: ", utils.HoneyGOVersion, "BVersion: ", utils.BVersion, "FH: ", Hash())
		// fmt.Println("Server Init: ", server.ServerInitialised, "REINIT: ", server.ServerREINIT)
		fmt.Println("Config:", config.GConfig)
		// fmt.Println("PlayerMap: ", server.PlayerMap)
		// fmt.Println("PlayerConnMap", server.PlayerConnMap)
		// fmt.Println("ConnPlayerMap", server.ConnPlayerMap)
		// fmt.Println("Mutexes: ", "PlayerMap: ", server.PlayerMapMutex, "PCM: ", server.PlayerConnMutex, "CPM: ", server.ConnPlayerMutex)
		fmt.Println("----------------------------------------")
		fmt.Println("Network")
		fmt.Println("----------------------------------------")
		fmt.Println("Auth: ", npacket.Hash())
		// fmt.Println("ClientConnectionMap: ", server.ClientConnectionMap)
		// fmt.Println("CCM Mutex: ", server.ClientConnectionMutex)
		// fmt.Println("Status Cache Map: ", server.StatusCache)
		// fmt.Println("Status Mutex: ", server.StatusMutex)
		// fmt.Println("StatusSemaphore", server.StatusSemaphore)
		fmt.Println("----------------------------------------")
		fmt.Println("Player Maps and states (package player)")
		fmt.Println("----------------------------------------")
		fmt.Println("POM: ", player.PlayerObjectMap)
		fmt.Println("PEM: ", player.PlayerEntityMap)
		fmt.Println("OPM: ", player.OnlinePlayerMap)
		fmt.Println("Mutexes: ", "POM: ", player.PlayerObjectMutex, "PEM: ", player.PlayerEntityMutex, "OPM", player.OnlinePlayerMutex)
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
