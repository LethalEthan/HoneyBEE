package server

import (
	"HoneyBEE/config"
	"HoneyBEE/packet"
	"sync"
	"time"

	"github.com/op/go-logging"
	"github.com/panjf2000/gnet"
)

var Log = logging.MustGetLogger("HoneyBEE")
var GlobalServer Server

type Server struct {
	*gnet.EventServer
	//pool *goroutine.Pool
	ConnectedSockets sync.Map
}

func NewServer(ip string, port string, multicore bool, tick bool, lockosthread bool, reuse bool, sendBuf int, recvBuf int, readBufferCap int) (*Server, error) {
	Log.Info("Generating Key chain")
	packet.GenerateKeys()
	GlobalServer = *new(Server)
	packet.CreateEntries()
	err := gnet.Serve(&GlobalServer, "tcp://"+ip+port, gnet.WithMulticore(multicore), gnet.WithNumEventLoop(config.GConfig.Server.NumEventLoop), gnet.WithTicker(tick), gnet.WithLockOSThread(lockosthread), gnet.WithReusePort(reuse), gnet.WithSocketSendBuffer(sendBuf), gnet.WithSocketRecvBuffer(recvBuf), gnet.WithReadBufferCap(readBufferCap), gnet.WithTCPKeepAlive(5*time.Second))
	return &GlobalServer, err
}

func (S *Server) OnInitComplete(Srv gnet.Server) (Action gnet.Action) {
	Log.Infof("HoneyBEE is listening on %s (multi-cores: %t, SO_REUSE: %t, Timeout: %d, loops: %d) ", Srv.Addr.String(), Srv.Multicore, Srv.ReusePort, Srv.TCPKeepAlive, Srv.NumEventLoop)
	return gnet.None
}

func (S *Server) OnOpened(Conn gnet.Conn) (Out []byte, Action gnet.Action) {
	Log.Infof("Socket with addr: %s has been opened...\n", Conn.RemoteAddr().String())
	C := *new(Client)
	C.RemoteAddr = Conn.RemoteAddr().String()
	// C.Conn = Conn
	C.State = HANDSHAKE
	C.PR = packet.CreatePacketReader([]byte{0x00})
	C.PW = packet.CreatePacketWriterWithCapacity(0x00, 8192)
	C.Read = make(chan []byte, 10)
	// Conn.SetContext(&C)
	S.ConnectedSockets.Store(Conn.RemoteAddr().String(), C)
	go C.ClientReact(Conn)
	Log.Debug("CR")
	Log.Debug(Conn.RemoteAddr().String())
	return nil, gnet.None
}

func (S *Server) OnClosed(Conn gnet.Conn, err error) (Action gnet.Action) {
	Log.Debugf("Socket with addr: %s is closing...\n", Conn.RemoteAddr().String())
	CO, load := S.ConnectedSockets.LoadAndDelete(Conn.RemoteAddr().String())
	if load {
		CC := CO.(Client)
		CC.Read <- nil
		Conn.SetContext(nil)
		close(CC.Read)
		Log.Infof("Socket with addr: %s is closed\n", Conn.RemoteAddr().String())
	} else {
		Log.Infof("Client object not found")
	}
	return
}
