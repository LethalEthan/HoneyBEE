package main

import (
	"HoneyBEE/config"
	"HoneyBEE/console"
	"HoneyBEE/mitm"
	"HoneyBEE/server"
	"HoneyBEE/utils"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/debug"
	"runtime/pprof"
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
	format   = logging.MustStringFormatter("%{color}[%{time:01-02-2006 15:04:05.000}] [%{level}] [%{shortfunc}]%{color:reset} %{message}")
	Log      = logging.MustGetLogger("HoneyBEE")
	conf     *config.Config
	err      error
	Panicked bool = false
	hprof    *os.File
	cprof    *os.File
)

func init() {
	//Hello from HoneyBEE
	//Logger Creation Start
	debug.SetMaxThreads(1024)
	debug.SetMaxStack(4294967296)
	defer console.DRECOVER()
	B1 := logging.NewLogBackend(os.Stderr, "", 0)       //New Backend
	B1Format := logging.NewBackendFormatter(B1, format) //Set Format
	B1LF := logging.AddModuleLevel(B1Format)            //Add formatting Levels
	B1LF.SetLevel(logging.DEBUG, "")
	logging.SetBackend(B1LF)
	//Logger Creation END
	Log.Info("HoneyBEE", utils.GetVersionString(), "starting...")
	fmt.Print(utils.Ascii, utils.Ascii2, "\n")
	//Remove unused Ascii strings for less memory cosumption
	utils.Ascii = ""
	utils.Ascii2 = ""
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
	//SetGCPercent
	debug.SetGCPercent(conf.Performance.GCPercent)
	if config.GConfig.Server.Port == "" {
		panic("Server port not defined!")
	}
	// pr := packet.CreatePacketReader([]byte{0xCC, 0x16, 0xC4, 0xF6, 0x01, 0x78, 0x9C, 0xED, 0x9D, 0x5F, 0x6C, 0x1C, 0x47, 0x19, 0xC0, 0xE7}) //0x03, 0x03, 0x80, 0x02}) //[]byte{0x03, 0xC4, 0x80})
	// T, NR, err := pr.ReadVarInt()
	// Log.Debug("T: ", T, "NR", NR, "err", err)
	// T2, NR2, err := pr.ReadVarInt()
	// Log.Debug("T2: ", T2, "NR", NR2, "err", err)
	// T3, NR3, err := pr.ReadVarInt()
	// Log.Debug("T3: ", T3, "NR", NR3, "err", err)
	//Log.Debug("Test", (0xC4 & 0x7F))
	// err = nil

	//Server Config Check
	Log.Info("Server Network Listener Started on port ", config.GConfig.Server.Port)
	Log.Info("Number of logical CPU's: ", runtime.NumCPU())
	if conf.Performance.CPU == 0 {
		Log.Info("Setting GOMAXPROCS to all available logical CPU's")
		runtime.GOMAXPROCS(runtime.NumCPU())
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
		go mitm.StartClient()
		if err != nil {
			panic(err)
		}
	} else {
		server.Init()
	}
	server.Init()
	go console.Console()
	go console.Shutdown()
	//go server.DebugServer()
	if conf.DEBUGOPTS.PacketAnal {
		Log.Warning("MITM Proxy mode enable")
	}
	runtime.GC()
}

func main() {
	defer console.DRECOVER()
	if console.Panicked {
		Log.Warning("Main: Panic is true, blocked main thread")
		for {
			time.Sleep(20000000)
		}
	}
	_, err := server.NewServer(config.GConfig.Server.Host, config.GConfig.Server.Port, config.GConfig.Server.MultiCore, false, config.GConfig.Server.LockOSThread, config.GConfig.Server.Reuse, config.GConfig.Server.SendBuf, config.GConfig.Server.RecieveBuf, config.GConfig.Server.ReadBufferCap)
	if err != nil {
		Log.Critical(err)
	}
}
