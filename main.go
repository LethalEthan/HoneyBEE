package main

import (
	"Packet"
	"net"
	"os"
	"runtime"
	"runtime/debug"
	"server"
	"time"

	logging "github.com/op/go-logging"
)

var (
	format         = logging.MustStringFormatter("%{color}[%{time:01-02-2006 15:04:05.000}] [%{level}] [%{shortfunc}]%{color:reset} %{message}")
	HoneyGOVersion = "1.0.0 (Build 17)"
	ServerPort     = ":25565"
	Log            = logging.MustGetLogger("HoneyGO")
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
	server.CurrentStatus = server.CreateStatusObject()
	//Logger Creation END
	server.Log = Log

	Log.Info("HoneyGO ", HoneyGOVersion, " starting...")

	//Network Listener on defined port 25565
	//TODO: Finish ConfigHandler for custom ports
	netlisten, err := net.Listen("tcp", ServerPort)
	if err != nil {
		Log.Fatal(err.Error())
		return
	}
	Log.Info("Server Network Listener Started on port", ServerPort)
	Log.Info("Number of logical CPU's: ", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU()) //Set it to the value of how many cores
	if runtime.NumCPU() < 2 {
		Log.Critical("Number of CPU's is less than 2 this could impact performance as this is a heavily threaded application")
	}
	Log.Info("Generating Key Chain")
	//NOTE: goroutines are light weight threads that can be reused with the same stack created before,
	//this will be useful when multiple clients connect but with some slight added memory usage
	go Packet.KeyGen() //Generate Keys used for client Authenication, will be controlled by config file (later release)

	//Accepts connection and creates new goroutine for the connection to be handled
	//other goroutines are stemmed from HandleConnection
	for {
		Connection, err := netlisten.Accept()

		if err != nil {
			Log.Error(err.Error())
			continue
		}
		Connection.SetDeadline(time.Now().Add(time.Second * 5))
		Log.Debug("Handshake Process Initiated")
		go server.HandleConnection(server.CreateClientConnection(Connection, server.HANDSHAKE))
	}
}
