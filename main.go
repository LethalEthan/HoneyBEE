package main

import (
	"Packet"
	"config"
	"net"
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"
	"server"
	"syscall"
	"time"

	logging "github.com/op/go-logging"
)

var (
	format         = logging.MustStringFormatter("%{color}[%{time:01-02-2006 15:04:05.000}] [%{level}] [%{shortfunc}]%{color:reset} %{message}")
	HoneyGOVersion = "1.0.0 (Build 19)"
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
	server.CurrentStatus = server.CreateStatusObject(578, "1.15.2")
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
	Log.Info("Server Network Listener Started on port", ServerPort)
	Log.Info("Number of logical CPU's: ", runtime.NumCPU())
	if conf.Performance.CPU == 0 {
		Log.Info("Setting GOMAXPROCS to all available logical CPU's")
		runtime.GOMAXPROCS(runtime.NumCPU()) //Set it to the value of how many cores
	} else {
		Log.Info("Setting GOMAXPROCS to config")
		runtime.GOMAXPROCS(conf.Performance.CPU)
	}
	if runtime.NumCPU() < 2 {
		Log.Critical("Number of CPU's is less than 2 this could impact performance as this is a heavily threaded application")
	}
	Log.Info("Generating Key Chain")
	//NOTE: goroutines are light weight threads that can be reused with the same stack created before,
	//this will be useful when multiple clients connect but with some slight added memory usage
	go Packet.KeyGen() //Generate Keys used for client Authenication, offline mode will not be supported (no piracy here bois)
	//Accepts connection and creates new goroutine for the connection to be handled
	//other goroutines are stemmed from HandleConnection
	//server.OnStart()
	go Shutdown()
	Log.Info("Accepting Connections")
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
				Connection.Close()
				Log.Info("Connection Closed")
				netlisten.Close()
				Log.Info("Net Listen Closed")
				os.Exit(0)
			}
			os.Exit(0)
		}
	}
}
