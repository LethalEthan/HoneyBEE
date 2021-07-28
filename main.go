package main

import (
	"HoneyGO/config"
	"HoneyGO/console"
	"HoneyGO/nserver"
	"HoneyGO/server"
	"HoneyGO/utils"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/debug"
	"runtime/pprof"

	logging "github.com/op/go-logging"
)

//R.I.P Alex, I'll miss you
//Most things defined in main have moved
var (
	format   = logging.MustStringFormatter("%{color}[%{time:01-02-2006 15:04:05.000}] [%{level}] [%{shortfunc}]%{color:reset} %{message}")
	Log      = logging.MustGetLogger("HoneyGO")
	conf     *config.Config
	err      error
	Panicked bool = false
	hprof    *os.File
	cprof    *os.File
)

func init() {
	//Hello from HoneyGO
	//Logger Creation Start
	defer console.DRECOVER()
	B1 := logging.NewLogBackend(os.Stderr, "", 0)       //New Backend
	B1Format := logging.NewBackendFormatter(B1, format) //Set Format
	B1LF := logging.AddModuleLevel(B1Format)            //Add formatting Levels
	B1LF.SetLevel(logging.DEBUG, "")
	logging.SetBackend(B1LF)
	server.Log = Log
	//Logger Creation END
	Log.Info("HoneyGO ", utils.HoneyGOVersion, " starting...")
	fmt.Print(utils.Ascii, utils.Ascii2, "\n")
	//Remove unused Ascii strings for less memory cosumption
	utils.Ascii = ""
	utils.Ascii2 = ""
	runtime.GC()
	//
	conf = config.ConfigStart()
	//MemProf
	if config.Memprofile != "" {
		hprof, err = os.Create(config.Memprofile)
		if err != nil {
			Log.Fatal(err)
		}
	}
	//CPUProf
	if config.Cpuprofile != "" {
		cprof, err = os.Create(config.Cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		pprof.StartCPUProfile(cprof)
	}
	//
	if conf.Performance.GCPercent == 0 {
		Log.Warning("GCPercent is 0!, GC will only activate via playerGC")
	}
	//SetGCPercent
	debug.SetGCPercent(conf.Performance.GCPercent)
	if config.GConfig.Server.Port == "" {
		panic("Server port not defined!")
	}
	//Server Config Check
	if conf.Server.ClientFrameBuffer == 0 || conf.Server.ReadBufferCap == 0 || conf.Server.RecieveBuf == 0 || conf.Server.SendBuf == 0 || conf.Server.Timeout <= 3 {
		panic("Please don't be stupid and set the buffers to 0 or timeout as less than 3 :/")
	}
	//--//
	Log.Info("Server Network Listener Started on port ", config.GConfig.Server.Port)
	Log.Info("Number of logical CPU's: ", runtime.NumCPU())
	if conf.Performance.CPU == 0 {
		Log.Info("Setting GOMAXPROCS to all available logical CPU's")
		runtime.GOMAXPROCS(runtime.NumCPU()) //Set it to the value of how many cores
	} else {
		Log.Info("Setting GOMAXPROCS to config: ", conf.Performance.CPU)
		runtime.GOMAXPROCS(conf.Performance.CPU)
	}
	if runtime.NumCPU() <= 3 || conf.Performance.CPU <= 3 {
		Log.Critical("Number of CPU's is less than 3 this could impact performance as this is a heavily threaded application")
	}
	//Log.Info("Generating Key Chain")
	//Packet.KeyGen() //Generate Keys used for client Authenication, offline mode will not be supported (no piracy here bois)
	if conf.DEBUGOPTS.PacketAnal {
		Log.Warning("Packet Analysis enabled, server will not be initialised")
	} else {
		if !config.GConfig.DEBUGOPTS.NewServer {
			Log.Critical("You are using the old server, this is deprecated and unsupported and will removed!")
			server.Init() //Initalise server
		} else {
			nserver.Init()
		}
	}
	go console.Console()
	go console.Shutdown()
	if conf.DEBUGOPTS.PacketAnal {
		Log.Warning("MITM Proxy mode enable")
	}
}

func main() {
	defer console.DRECOVER()
	if console.Panicked {
		Log.Warning("Main: Panic is true, blocked main thread")
		for {
		}
	}
	if config.GConfig.DEBUGOPTS.NewServer {
		_, err := nserver.NewServer(config.GConfig.Server.Host, config.GConfig.Server.Port, config.GConfig.Server.MultiCore, false, config.GConfig.Server.LockOSThread, config.GConfig.Server.Reuse, config.GConfig.Server.SendBuf, config.GConfig.Server.RecieveBuf, config.GConfig.Server.ReadBufferCap)
		Log.Critical(err)
	} else {
		Log.Info("Accepting Connections")
		Log.Critical("OLD SERVER DEPRECATED AND REMOVED WILL DO NOTHING!")
		server.Runner()
	}
}
