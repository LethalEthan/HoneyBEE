package main

import (
	"Packet"
	"config"
	//	"fmt"

	"net"
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"
	"server"
	"syscall"
	"time"
	//	"worldtime"
	//	_ "net/http/pprof"
	//	"net/http"
	logging "github.com/op/go-logging"
)

//R.I.P Alex, I'll miss you
var (
	format         = logging.MustStringFormatter("%{color}[%{time:01-02-2006 15:04:05.000}] [%{level}] [%{shortfunc}]%{color:reset} %{message}")
	HoneyGOVersion = "1.0.0 (Build 23)"
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
	//chunk.BuildChunk(0, 0) //DEBUG: Remove me later
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
	if runtime.NumCPU() < 2 || conf.Performance.CPU < 2 && conf.Performance.CPU != 0 {
		Log.Critical("Number of CPU's is less than 2 this could impact performance as this is a heavily threaded application")
	}
	Log.Info("Generating Key Chain")
	//NOTE: goroutines are light weight threads that can be reused with the same stack created before,
	//this will be useful when multiple clients connect but with some slight added memory usage
	Packet.KeyGen() //Generate Keys used for client Authenication, offline mode will not be supported (no piracy here bois)
	//server.OnStart()
	//S := make(chan bool)
	//go worldtime.WorldTime(S)
	server.Init() //Initalise server
	go Shutdown()
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

func Shutdown() {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	select {
	case <-shutdown:
		{
			Log.Warning("Starting shutdown")
			Run = false
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
