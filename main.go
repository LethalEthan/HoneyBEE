package main

import (
	"Packet"
	"bufio"
	"config"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net"
	"nserver"
	"os"
	"os/signal"
	"player"
	"runtime"
	"runtime/debug"
	"runtime/pprof"
	"server"
	"sync"
	"syscall"
	"time"
	"utils"

	//	"worldtime"
	//_ "net/http/pprof"
	//	"net/http"

	logging "github.com/op/go-logging"
)

//R.I.P Alex, I'll miss you
var (
	format         = logging.MustStringFormatter("%{color}[%{time:01-02-2006 15:04:05.000}] [%{level}] [%{shortfunc}]%{color:reset} %{message}")
	HoneyGOVersion = "1.1.1 (Build 71)"
	BVersion       = 71
	Log            = logging.MustGetLogger("HoneyGO")
	ServerPort     string
	conf           *config.Config
	Connection     net.Conn
	netlisten      net.Listener
	err            error
	Run            bool
	RunMutex       = sync.Mutex{}
	mem            runtime.MemStats
	Panicked       bool = false
)

func init() {
	//Hello from HoneyGO
	//Logger Creation Start
	defer DRECOVER()
	B1 := logging.NewLogBackend(os.Stderr, "", 0)       //New Backend
	B1Format := logging.NewBackendFormatter(B1, format) //Set Format
	B1LF := logging.AddModuleLevel(B1Format)            //Add formatting Levels
	B1LF.SetLevel(logging.DEBUG, "")
	logging.SetBackend(B1LF)
	server.Log = Log
	//Logger Creation END
	Log.Info("HoneyGO ", HoneyGOVersion, " starting...")
	fmt.Print(utils.Ascii, utils.Ascii2, "\n")
	//Remove unused Ascii strings for less memory cosumption
	utils.Ascii = ""
	utils.Ascii2 = ""
	runtime.GC()
	//
	Run = true
	conf = config.ConfigStart()
	if conf.Performance.GCPercent == 0 {
		Log.Warning("GCPercent is 0!, GC will only activate via playerGC")
	}
	debug.SetGCPercent(conf.Performance.GCPercent)
	ServerPort = conf.Server.Port
	if ServerPort == "" {
		Log.Warning("Server port not defined!")
	}
	if !conf.DEBUGOPTS.NewServer {
		netlisten, err = net.Listen("tcp", ServerPort)
		if err != nil {
			Log.Fatal(err.Error())
			os.Exit(1)
		} else {
			Log.Info("NetListen started")
		}
	}
	if conf.Server.ClientFrameBuffer == 0 || conf.Server.ReadBufferCap == 0 || conf.Server.RecieveBuf == 0 || conf.Server.SendBuf == 0 || conf.Server.Timeout == 0 {
		panic("Please don't be stupid and set the buffers or timeout as 0 :/")
	}
	if Panicked {
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
	if conf.DEBUGOPTS.PacketAnal {
		Log.Warning("Packet Analysis enabled, server will not be initialised")
	} else {
		server.Init() //Initalise server
	}
	go Console()
	go Shutdown()
	//DebugOps()
	if conf.DEBUGOPTS.PacketAnal {
		Log.Warning("MITM Proxy mode enable")
	}
}

func main() {
	defer DRECOVER()
	if Panicked {
		Log.Warning("Main: Panic is true, blocked main thread")
		for {
		}
	}
	conf = config.GetConfig()
	if conf.DEBUGOPTS.NewServer {
		nserver.NewServer(conf.Server.Host, conf.Server.Port, conf.Server.MultiCore, false, conf.Server.Reuse, conf.Server.SendBuf, conf.Server.RecieveBuf, conf.Server.ReadBufferCap)
	} else {
		Log.Info("Accepting Connections")
		for GetRun() {
			Connection, err = netlisten.Accept()
			if err != nil && Run == true {
				Log.Error(err.Error())
				continue
			}
			Log.Debug("C: ", Connection)
			Connection.SetDeadline(time.Now().Add(time.Duration(time.Second * 5)))
			//Connection.SetDeadline(time.Now().Add(time.Duration(1000000000 * conf.Server.Timeout)))
			Log.Debug("Handshake Process Initiated")
			if conf.DEBUGOPTS.PacketAnal {
				server.AnalProbe(Connection, conf.DEBUGOPTS.PacketAnalAddress)
			} else {
				go server.HandleConnection(server.CreateClientConnection(Connection, server.HANDSHAKE))
			}
		}
	}
}

//DRECOVER - Recovery -- TBD close server go routines
func DRECOVER() {
	if r := recover(); r != nil {
		Panicked = true //TBD: link to event system to stop everything
		go SetRun(false)
		go server.SetRun(false)
		go func() { //Lock mutexes in case something carries on upon recovery
			server.RunMutex.Lock()
			server.PlayerMapMutex.Lock()
			server.ConnPlayerMutex.Lock()
			server.ClientConnectionMutex.Lock()
			server.StatusMutex.Lock()
			server.PlayerConnMutex.Lock()
			server.RunMutex.Lock()
			player.OnlinePlayerMutex.Lock()
			player.PlayerEntityMutex.Lock()
			player.PlayerObjectMutex.Lock()
			server.Run = false //reset as false in case the go routine did not
			server.GCPShutdown <- true
		}()
		go server.StatusSemaphore.StopSemaphore() //TODO: Make all semaphores be able to be stopped
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		fmt.Println("Server encountered panic! reason: ", r)
		fmt.Println("Printing Debug, please create an issue and send this!")
		fmt.Println("----------------------------------------")
		fmt.Println("Server Maps and states (package server)")
		fmt.Println("----------------------------------------")
		fmt.Println("Server run state:", server.Run, "Mutex: ", server.RunMutex)
		fmt.Println("HoneyGOVersion: ", HoneyGOVersion, "BVersion: ", BVersion, "FH: ", Hash())
		fmt.Println("Server Init: ", server.ServerInitialised, "REINIT: ", server.ServerREINIT)
		fmt.Println("Config:", config.GConfig)
		fmt.Println("PlayerMap: ", server.PlayerMap)
		fmt.Println("PlayerConnMap", server.PlayerConnMap)
		fmt.Println("ConnPlayerMap", server.ConnPlayerMap)
		fmt.Println("Mutexes: ", "PlayerMap: ", server.PlayerMapMutex, "PCM: ", server.PlayerConnMutex, "CPM: ", server.ConnPlayerMutex)
		fmt.Println("----------------------------------------")
		fmt.Println("Network")
		fmt.Println("----------------------------------------")
		fmt.Println("Auth: ", server.Hash())
		fmt.Println("ClientConnectionMap: ", server.ClientConnectionMap)
		fmt.Println("CCM Mutex: ", server.ClientConnectionMutex)
		fmt.Println("Status Cache Map: ", server.StatusCache)
		fmt.Println("Status Mutex: ", server.StatusMutex)
		fmt.Println("StatusSemaphore", server.StatusSemaphore)
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
		printDebugStats(mem)
		//os.Exit(1)
		panic(r)
	}
}

var shutdown = make(chan os.Signal, 1)

//Shutdown - listens for sigterm and exits
func Shutdown() {
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	select {
	case <-shutdown:
		{
			Log.Warning("Starting shutdown")
			conf := config.GetConfig()
			if !conf.DEBUGOPTS.NewServer {
				SetRun(false)
				server.SetRun(false)
			} else {
				nserver.GlobalServer.Shutdown()
			}
			if conf.Performance.EnableGCPlayer { //Don't send on an unintialised channel otherwise it will deadlock
				server.GCPShutdown <- true
			}
			server.ClientConnectionMutex.Lock()
			Log.Debug(server.ClientConnectionMap) //Check if any connections are still active in the map, there shouldn't be any left over
			DEBUG := true
			if DEBUG {
				printDebugStats(mem)
			}
			os.Exit(0)
		}
	}
}

//DebugOps - Do stuff for debugging that runs on startup
// func DebugOps() {
// 	Block := blocks.GetBlockID(1)
// 	fmt.Print(Block)
// 	// chunk.BuildChunk(0, 0, 8)
// 	// T, S := chunk.COORDSToInts("-69,420")
// 	// fmt.Print("\n", T, S)
// 	// worker.CreateFlatStoneWorld()
// 	CreateRegions()
// 	//TESTING -1,-1
// 	Region, bool, err := world.GetRegionByInt(1, -1)
// 	if err != nil || bool != true {
// 		panic("Region ikke fundet")
// 	}
// 	fmt.Print("\n", Region.ID, "\n")
// 	C, err := world.GetChunkFromRegion(Region, 511, -511)
// 	if err != nil {
// 		Log.Error(err)
// 		fmt.Print("ERROR")
// 	}
// 	fmt.Print("\n", len(Region.Data))
// 	fmt.Print("\n", C.ChunkPosX, ",", C.ChunkPosZ)
// 	//--//
// 	//Testing -1,1
// 	testregion(-1, 1, -500, 500)
// 	//Testing 1,-1
// 	testregion(1, -1, 500, -500)
// 	//Error Check
// 	Log.Warning("this region is meant to fail")
// 	testregion(1, 1, -1000, -1000)
// 	// for i := 0; i < /*len(Fuck.Data)*/ 0; i++ {
// 	// 	fmt.Print("\n", Fuck.Data[i].ChunkPosX)
// 	// 	fmt.Print("\n", Fuck.Data[i].ChunkPosZ)
// 	// 	runtime.GC()
// 	// 	//time.Sleep(100000)
// 	// }
// 	//C := world.GetChunkFromRegion(Region, 290, 511)
// 	//fmt.Print("\n", C.ChunkPosX, "\n", C.ChunkPosZ)
// 	READ := false
// 	if READ {
// 		for j := 256; j <= 511; j++ {
// 			for i := 256; i <= 511; i++ {
// 				C, err := world.GetChunkFromRegion(Region, i, j)
// 				if err != nil {
// 					panic(err)
// 				}
// 				fmt.Print(C.ChunkPosX, " ", C.ChunkPosZ, " ")
// 			}
// 		}
// 	}
// }

// func CreateRegions() {
// world.CreateRegion(-1, -1)
// world.CreateRegion(-1, 1)
// world.CreateRegion(1, -1)
// world.CreateRegion(0, 0) //Each around 5MB each
// world.CreateRegion(0, 1)
// world.CreateRegion(1, 0)
// world.CreateRegion(1, 1)
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
// }

func Console() {
	defer DRECOVER()
	Log := logging.MustGetLogger("HoneyGO")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		switch scanner.Text() {
		case "help":
			Log.Warning("There is no help atm :(")
			Log.Warning("This is a simple, quick and dirty way of doing commands, a proper thing is being made later")
		case "shutdown":
			shutdown <- os.Interrupt
		case "stop":
			shutdown <- os.Interrupt
		case "exit":
			shutdown <- os.Interrupt
		case "reload":
			if !conf.DEBUGOPTS.NewServer {
				SetRun(false)
				server.SetRun(false)
				config.ConfigReload()
				server.GCPShutdown <- true
				server.ServerReload()
				server.SetRun(true)
				SetRun(true)
			} else {
				config.ConfigReload()
				//nserver.Conf = config.GConfig
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
				if conf.Server.ClientFrameBuffer == 0 || conf.Server.ReadBufferCap == 0 || conf.Server.RecieveBuf == 0 || conf.Server.SendBuf == 0 || conf.Server.Timeout == 0 {
					panic("Please don't be stupid and set the buffers or timeout as 0 :/")
				}
				Log.Critical("If you changed new server to old server this will not be reloaded or changed!")
			}
		case "GC":
			runtime.GC()
			Log.Info("GC invoked")
		case "mem":
			printDebugStats(mem)
		case "SSM":
			Log.Debug(server.StatusCache)
		case "CCM":
			Log.Debug(server.ClientConnectionMap)
		case "panic":
			panic("panicked, you told me to :)")
		case "profile":
			pprof.WriteHeapProfile(honeyprof)
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

	fmt.Println("mem.NumForcedGC:", mem.NumForcedGC)

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

// func testregion(X int64, Z int64, CX int, CZ int) {
// 	//TESTING -1,-1
// 	Region, bool, err := world.GetRegionByInt(X, Z)
// 	if err != nil || bool != true {
// 		panic("Region ikke fundet")
// 	}
// 	fmt.Print("\n", Region.ID, "\n")
// 	C, err := world.GetChunkFromRegion(Region, CX, CZ)
// 	if err != nil {
// 		Log.Error(err)
// 		fmt.Print("ERROR")
// 	}
// 	fmt.Print("\n", len(Region.Data))
// 	fmt.Print("\n", C.ChunkPosX, ",", C.ChunkPosZ, "\n")
// }

var MD5 string

func Hash() string {
	file, err := os.Open(os.Args[0])
	if err != nil {
		MD5 = "00000000000000000000000000000000"
	}
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		MD5 = "00000000000000000000000000000000"
	}
	//Get the 16 bytes hash
	hBytes := hash.Sum(nil)[:16]
	file.Close()
	MD5 = hex.EncodeToString(hBytes) //Convert bytes to string
	return MD5
}
