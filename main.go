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
	"sync"
	"syscall"
	"time"
	"utils"
	"world"
	//	"worldtime"
	//_ "net/http/pprof"
	//	"net/http"
	logging "github.com/op/go-logging"
)

//R.I.P Alex, I'll miss you
var (
	format         = logging.MustStringFormatter("%{color}[%{time:01-02-2006 15:04:05.000}] [%{level}] [%{shortfunc}]%{color:reset} %{message}")
	HoneyGOVersion = "1.0.1 (Build 35)"
	BVersion       = 35
	Log            = logging.MustGetLogger("HoneyGO")
	ServerPort     string
	conf           *config.Config
	Connection     net.Conn
	netlisten      net.Listener
	err            error
	Run            bool
	RunMutex       = sync.Mutex{}
	mem            runtime.MemStats
)

func main() {
	//Hello from HoneyGO
	//Logger Creation Start
	B1 := logging.NewLogBackend(os.Stderr, "", 0)       //New Backend
	B1Format := logging.NewBackendFormatter(B1, format) //Set Format
	B1LF := logging.AddModuleLevel(B1Format)            //Add formatting Levels
	B1LF.SetLevel(logging.DEBUG, "")
	logging.SetBackend(B1LF)
	server.Log = Log
	//Logger Creation END
	Log.Info("HoneyGO ", HoneyGOVersion, " starting...")
	fmt.Print(utils.Ascii, utils.Ascii2, "\n")
	Run = true
	conf := config.ConfigStart()
	if conf.Performance.GCPercent == 0 {
		Log.Warning("GCPercent is 0!, GC will only activate via playerGC")
	}
	debug.SetGCPercent(conf.Performance.GCPercent)
	ServerPort = conf.Server.Port
	if ServerPort == "" {
		Log.Warning("Server port not defined!")
	}
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
		Log.Info("Setting GOMAXPROCS to config: ", conf.Performance.CPU)
		runtime.GOMAXPROCS(conf.Performance.CPU)
	}
	if runtime.NumCPU() <= 3 || conf.Performance.CPU <= 2 {
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
	//ConTO, err := net.Dial("tcp", "192.168.0.42:25565")
	// go func() {
	// 	fmt.Println(http.ListenAndServe("localhost:6060", nil))
	// }()
	//chunk.CreateNewChunkSection()
	//var ConnTO net.Conn
	if conf.DEBUGOPTS.PacketAnal {
		Log.Warning("MITM Proxy mode enable")
		//ConnTO = ConnTO
		// go server.AnalProbe(Client net.Conn, Server net.Conn)(ConnTO)
	}
	for /*Run*/ GetRun() {
		Connection, err = netlisten.Accept()
		if err != nil && Run == true {
			Log.Error(err.Error())
			continue
		}
		Connection.SetDeadline(time.Now().Add(time.Duration(1000000000 * conf.Server.Timeout)))
		Log.Debug("Handshake Process Initiated")
		// if conf.DEBUGOPTS.PacketAnal {
		// 	ConnTO, err := net.Dial("tcp", conf.DEBUGOPTS.PacketAnalAddress)
		// 	if err != nil {
		// 		Log.Error("Couldn't connect to server defined in config, packet-anal-address: ", conf.DEBUGOPTS.PacketAnalAddress)
		// 	}
		// 	server.AnalProbe(Connection, ConnTO)
		// } else {
		go server.HandleConnection(server.CreateClientConnection(Connection, server.HANDSHAKE))
	}
}

//
var shutdown = make(chan os.Signal, 1)

func Shutdown() {
	//shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	select {
	case <-shutdown:
		{
			Log.Warning("Starting shutdown")
			SetRun(false)
			server.SetRun(false)
			conf := config.GetConfig()
			if conf.Performance.EnableGCPlayer {
				server.GCPShutdown <- true
			}
			server.ClientConnectionMutex.Lock()
			Log.Debug(server.ClientConnectionMap)
			DEBUG := true
			if DEBUG {
				printDebugStats(mem)
			}
			// netlisten.Close()
			// Log.Info("Net Listen Closed")
			// Connection.Close()
			// Log.Info("Connection Closed")
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
	//TESTING -1,-1
	Region, bool, err := world.GetRegionByInt(1, -1)
	if err != nil || bool != true {
		panic("Region ikke fundet")
	}
	fmt.Print("\n", Region.ID, "\n")
	C, err := world.GetChunkFromRegion(Region, 511, -511)
	if err != nil {
		Log.Error(err)
		fmt.Print("ERROR")
	}
	fmt.Print("\n", len(Region.Data))
	fmt.Print("\n", C.ChunkPosX, ",", C.ChunkPosZ)
	//--//
	//Testing -1,1
	testregion(-1, 1, -500, 500)
	//Testing 1,-1
	testregion(1, -1, 500, -500)
	//Error Check
	Log.Warning("this region is meant to fail")
	testregion(1, 1, -1000, -1000)
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
	world.CreateRegion(-1, -1)
	world.CreateRegion(-1, 1)
	world.CreateRegion(1, -1)
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
		case "stop":
			shutdown <- os.Interrupt
		case "reload":
			SetRun(false)
			server.SetRun(false)
			config.ConfigReload()
			server.GCPShutdown <- true
			server.ServerReload()
			server.SetRun(true)
			SetRun(true)
		case "GC":
			runtime.GC()
		case "mem":
			printDebugStats(mem)
		default:
			Log.Warning("Unknown command")
		}
	}
}

func printDebugStats(mem runtime.MemStats) {

	runtime.ReadMemStats(&mem)

	fmt.Println("mem.Alloc:", mem.Alloc)

	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)

	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)

	fmt.Println("mem.NumGC:", mem.NumGC)

	fmt.Println("-----")

}

func GetRun() bool {
	RunMutex.Lock()
	R := Run
	RunMutex.Unlock()
	return R
}

func SetRun(v bool) {
	RunMutex.Lock()
	Run = v
	RunMutex.Unlock()
}

func testregion(X int64, Z int64, CX int, CZ int) {
	//TESTING -1,-1
	Region, bool, err := world.GetRegionByInt(X, Z)
	if err != nil || bool != true {
		panic("Region ikke fundet")
	}
	fmt.Print("\n", Region.ID, "\n")
	C, err := world.GetChunkFromRegion(Region, CX, CZ)
	if err != nil {
		Log.Error(err)
		fmt.Print("ERROR")
	}
	fmt.Print("\n", len(Region.Data))
	fmt.Print("\n", C.ChunkPosX, ",", C.ChunkPosZ, "\n")
	//--//
}
