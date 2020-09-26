package main

import (
	"Packet"
	"blocks"
	"bufio"
	"config"
	"fmt"
	"net"
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"
	"server"
	"syscall"
	"time"
	"world"
	//	"worldtime"
	//_ "net/http/pprof"
	//	"net/http"
	logging "github.com/op/go-logging"
)

//R.I.P Alex, I'll miss you
var (
	format         = logging.MustStringFormatter("%{color}[%{time:01-02-2006 15:04:05.000}] [%{level}] [%{shortfunc}]%{color:reset} %{message}")
	HoneyGOVersion = "1.0.0 (Build 28)"
	BVersion       = 28
	Log            = logging.MustGetLogger("HoneyGO")
	ServerPort     string
	conf           *config.Config
	Connection     net.Conn
	netlisten      net.Listener
	err            error
	Run            bool
)

func main() {
	//Hello from HoneyGO
	debug.SetGCPercent(30)
	//Logger Creation Start
	B1 := logging.NewLogBackend(os.Stderr, "", 0)       //New Backend
	B1Format := logging.NewBackendFormatter(B1, format) //Set Format
	B1LF := logging.AddModuleLevel(B1Format)            //Add formatting Levels
	B1LF.SetLevel(logging.DEBUG, "")
	logging.SetBackend(B1LF)
	server.Log = Log
	//Logger Creation END
	Log.Info("HoneyGO ", HoneyGOVersion, " starting...")
	Run = true
	conf := config.ConfigStart()
	ServerPort = conf.Server.Port
	netlisten, err = net.Listen("tcp", ServerPort)
	if err != nil {
		Log.Fatal(err.Error())
		return
	}
	//--//
	Log.Info("Server Network Listener Started on port", ServerPort)
	Log.Info("Number of logical CPU's: ", runtime.NumCPU())
	if conf.Performance.CPU == 0 {
		Log.Info("Setting GOMAXPROCS to all available logical CPU's")
		runtime.GOMAXPROCS(runtime.NumCPU()) //Set it to the value of how many cores
	} else {
		Log.Info("Setting GOMAXPROCS to config")
		runtime.GOMAXPROCS(conf.Performance.CPU)
	}
	if runtime.NumCPU() < 2 || conf.Performance.CPU <= 2 && conf.Performance.CPU != 0 {
		Log.Critical("Number of CPU's is less than 3 this could impact performance as this is a heavily threaded application")
	}
	Log.Info("Generating Key Chain")
	//NOTE: goroutines are light weight threads that can be reused with the same stack created before,
	//this will be useful when multiple clients connect but with some slight added memory usage
	Packet.KeyGen() //Generate Keys used for client Authenication, offline mode will not be supported (no piracy here bois)
	server.Init()   //Initalise server
	go Console()
	go Shutdown()
	//DebugOps()
	//Accepts connection and creates new goroutine for the connection to be handled
	//other goroutines are stemmed from HandleConnection
	Log.Info("Accepting Connections")
	// go func() {
	// 	fmt.Println(http.ListenAndServe("localhost:6060", nil))
	// }()
	//chunk.CreateNewChunkSection()
	for Run {
		Connection, err = netlisten.Accept()
		if err != nil && Run == true {
			Log.Error(err.Error())
			continue
		}
		Connection.SetDeadline(time.Now().Add(time.Duration(1000000000 * conf.Server.Timeout)))
		Log.Debug("Handshake Process Initiated")
		go server.HandleConnection(server.CreateClientConnection(Connection, server.HANDSHAKE))
	}
}

var shutdown = make(chan os.Signal, 1)

func Shutdown() {
	//shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	select {
	case <-shutdown:
		{
			Log.Warning("Starting shutdown")
			Run = false
			time.Sleep(2000000000) //Let the loop finish before we do stuff
			if netlisten != nil && Connection != nil {
				//worldtime.Shutdown()
				Connection.Close()
				Log.Info("Connection Closed")
				netlisten.Close()
				Log.Info("Net Listen Closed")
				//worldtime.Shutdown()
				os.Exit(0)
			}
			//worldtime.Shutdown()
			os.Exit(0)
		}
	}
}

//DebugOps - Do stuff for debugging that runs on startup
func DebugOps() {
	Block := blocks.GetBlockID(1)
	fmt.Print(Block)
	// chunk.BuildChunk(0, 0, 8)
	// T, S := chunk.COORDSToInts("-69,420")
	// fmt.Print("\n", T, S)
	// worker.CreateFlatStoneWorld()
	CreateRegions()
	Region, err := world.GetRegionByID("1,1")
	if err != nil {
		panic("Region ikke fundet")
	}
	// for i := 0; i < /*len(Fuck.Data)*/ 0; i++ {
	// 	fmt.Print("\n", Fuck.Data[i].ChunkPosX)
	// 	fmt.Print("\n", Fuck.Data[i].ChunkPosZ)
	// 	runtime.GC()
	// 	//time.Sleep(100000)
	// }
	//C := world.GetChunkFromRegion(Region, 290, 511)
	//fmt.Print("\n", C.ChunkPosX, "\n", C.ChunkPosZ)
	READ := false
	if READ {
		for j := 256; j <= 511; j++ {
			for i := 256; i <= 511; i++ {
				C, err := world.GetChunkFromRegion(Region, i, j)
				if err != nil {
					panic(err)
				}
				fmt.Print(C.ChunkPosX, " ", C.ChunkPosZ, " ")
			}
		}
	}
}

func CreateRegions() {
	world.CreateRegion(0, 0) //Each around 5MB each
	world.CreateRegion(0, 1)
	world.CreateRegion(1, 0)
	world.CreateRegion(1, 1)
	// for j := 0; j <= 8; j++ {
	// 	for i := 0; i <= 8; i++ {
	// 		go world.CreateRegion(int64(i), int64(j))
	// 		runtime.GC()
	// 	}
	// }
	//runtime.GC()
	// for j := 0; j <= 20; j++ {
	// 	for i := 0; i <= 20; i++ {
	// 		R := world.GetRegionByInt(i, j)
	// 		fmt.Print(R.ID)
	// 	}
	// }

	//go world.CreateRegion(0, 0)
	// if val, tmp := world.RegionChunkMap.Get("0,0"); tmp {
	// 	fmt.Print(val)
	// }
}

func Console() {
	Log := logging.MustGetLogger("HoneyGO")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		switch scanner.Text() {
		case "help":
			Log.Warning("There is no help atm :(")
			Log.Warning("This is a simple, quick and dirty way of doing commands, a proper thing is being made bts")
		case "shutdown":
			shutdown <- os.Interrupt
		default:
			Log.Warning("Unknown command")
		}
	}
}
