package server

import (
	"HoneyBEE/config"
	"HoneyBEE/packet"
	"sync"
	"time"

	logging "github.com/op/go-logging"
	"github.com/panjf2000/gnet"
)

var Log = logging.MustGetLogger("HoneyBEE")
var GlobalServer *Server

type Server struct {
	*gnet.EventServer
	//pool *goroutine.Pool
	ConnectedSockets sync.Map
}

func NewServer(ip string, port string, multicore bool, tick bool, lockosthread bool, reuse bool, sendBuf int, recvBuf int, readBufferCap int) (Server, error) {
	Log.Info("Generating Key chain")
	packet.GenerateKeys()
	S := new(Server)
	packet.CreateEntries()
	GlobalServer = S
	err := gnet.Serve(S, "tcp://"+ip+port, gnet.WithMulticore(multicore), gnet.WithNumEventLoop(config.GConfig.Server.NumEventLoop), gnet.WithTicker(tick), gnet.WithLockOSThread(lockosthread), gnet.WithReusePort(reuse), gnet.WithSocketSendBuffer(sendBuf), gnet.WithSocketRecvBuffer(recvBuf), gnet.WithReadBufferCap(readBufferCap), gnet.WithTCPKeepAlive(5*time.Second))
	return *S, err
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
	Conn.SetContext(&C)
	S.ConnectedSockets.Store(Conn.RemoteAddr().String(), C)
	go C.ClientReact(Conn)
	Log.Debug("CR")
	Log.Debug(Conn.RemoteAddr().String())
	return nil, gnet.None
}

func (S *Server) OnClosed(Conn gnet.Conn, err error) (Action gnet.Action) {
	Log.Debugf("Socket with addr: %s is closing...\n", Conn.RemoteAddr().String())
	S.ConnectedSockets.Delete(Conn.RemoteAddr().String())
	CC, tmp := Conn.Context().(*Client)
	if !tmp {
		Log.Critical("Conn Context is nil!")
	}
	CC.Read <- nil
	Conn.SetContext(nil)
	Log.Infof("Socket with addr: %s is closed\n", Conn.RemoteAddr().String())
	return
}
