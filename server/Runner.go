package server

import (
	"HoneyGO/config"
	"net"
	"os"
	"sync"
	"time"
)

var (
	Connection        net.Conn
	netlisten         net.Listener
	err               error
	NetServerRun      bool
	NetServerRunMutex sync.Mutex
)

func Runner() {
	if !config.GConfig.DEBUGOPTS.NewServer {
		netlisten, err = net.Listen("tcp", config.GConfig.Server.Port)
		if err != nil {
			Log.Fatal(err.Error())
			os.Exit(1)
		} else {
			Log.Info("NetListen started")
		}
	}
	for GetNetServerRun() {
		Connection, err = netlisten.Accept()
		if err != nil && Run == true {
			Log.Error(err.Error())
			continue
		}
		Log.Debug("C: ", Connection)
		Connection.SetDeadline(time.Now().Add(time.Duration(time.Second * 5)))
		//Connection.SetDeadline(time.Now().Add(time.Duration(1000000000 * conf.Server.Timeout)))
		Log.Debug("Handshake Process Initiated")
		if config.GConfig.DEBUGOPTS.PacketAnal {
			AnalProbe(Connection, config.GConfig.DEBUGOPTS.PacketAnalAddress)
		} else {
			go HandleConnection(CreateClientConnection(Connection, HANDSHAKE))
		}
	}
}

func GetNetServerRun() bool {
	RunMutex.Lock()
	R := Run
	RunMutex.Unlock()
	return R
}
func SetNetServerRun(v bool) {
	RunMutex.Lock()
	Run = v
	RunMutex.Unlock()
}
