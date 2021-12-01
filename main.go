package main

import (
	"HoneyBEE/config"
	"HoneyBEE/console"
	"HoneyBEE/server"
	"HoneyBEE/utils"
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"time"

	logging "github.com/op/go-logging"
)

//R.I.P Alex, I'll miss you
//R.I.P Winnie
//R.I.P Grandad
//R.I.P Julia
//R.I.P Grandpa

//You will all be missed and never forgotten

//Most things defined in main have moved
var (
	format        = logging.MustStringFormatter("%{color}[%{time:01-02-2006 15:04:05.000}] [%{level}] [%{shortfunc}]%{color:reset} %{message}")
	Log           = logging.MustGetLogger("HoneyBEE")
	Panicked bool = false
)

func main() {
	// Hello from HoneyBEE
	// defer profile.Start(profile.MemProfile).Stop()
	// Logger Creation Start
	defer console.DRECOVER()
	B1 := logging.NewLogBackend(os.Stderr, "", 0)       // New Backend
	B1Format := logging.NewBackendFormatter(B1, format) // Set Format
	B1LF := logging.AddModuleLevel(B1Format)            // Add formatting Levels
	err := config.ConfigStart()
	if err != nil {
		panic(err)
	}
	if config.GConfig.Server.DEBUG {
		B1LF.SetLevel(logging.DEBUG, "")
		Log.Debug("Debug mode enabled")
	} else {
		B1LF.SetLevel(logging.INFO, "")
	}
	logging.SetBackend(B1LF)
	// Logger Creation END
	Log.Info("HoneyBEE", utils.GetVersionString(), "starting...")
	fmt.Print(utils.Ascii, utils.Ascii2, "\n")
	// Remove unused Ascii strings for less memory cosumption
	utils.Ascii = ""
	utils.Ascii2 = ""
	// SetGCPercent
	if config.GConfig.Performance.GCPercent > 0 {
		debug.SetGCPercent(config.GConfig.Performance.GCPercent)
	}
	if config.GConfig.Server.Port == "" {
		panic("Server port not defined!")
	}
	// Server Config Check
	Log.Info("Server Network Listener Started on port ", config.GConfig.Server.Port)
	Log.Info("Number of logical CPU's: ", runtime.NumCPU())
	if config.GConfig.Performance.CPU == 0 {
		Log.Info("Setting GOMAXPROCS to all available logical CPU's")
		runtime.GOMAXPROCS(runtime.NumCPU())
	} else {
		Log.Info("Setting GOMAXPROCS to config: ", config.GConfig.Performance.CPU)
		runtime.GOMAXPROCS(config.GConfig.Performance.CPU)
	}
	if runtime.NumCPU() <= 3 || config.GConfig.Performance.CPU <= 3 {
		Log.Critical("Number of CPU's is less than 3 this could impact performance as this is a heavily threaded application")
	}
	go console.Console()
	go console.Shutdown()
	runtime.GC()
	if console.Panicked {
		Log.Warning("Main: Panic is true, blocked main thread")
		for {
			time.Sleep(20000000)
		}
	}
	_, err = server.NewServer(config.GConfig.Server.Host, config.GConfig.Server.Port, config.GConfig.Server.MultiCore, false, config.GConfig.Server.LockOSThread, config.GConfig.Server.Reuse, config.GConfig.Server.SendBuf, config.GConfig.Server.RecieveBuf, config.GConfig.Server.ReadBufferCap)
	if err != nil {
		Log.Critical(err)
	}
}
