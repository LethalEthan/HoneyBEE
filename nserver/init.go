package nserver

import (
	"HoneyBEE/npacket"
	"HoneyBEE/world"
	"sync"
	"time"

	logging "github.com/op/go-logging"
	"github.com/panjf2000/gnet"
)

var Log = logging.MustGetLogger("HoneyBEE")
var DEBUG = true
var GlobalServer *Server

type Server struct {
	*gnet.EventServer
	//pool *goroutine.Pool
	ConnectedSockets sync.Map
}

//Bri-ish init bruv
func Init() {
	// go server.StatusSemaphore.Start()
	// server.StatusSemaphore.FlushAndSetSemaphore(server.StatusCache)
	// server.CurrentStatus = server.CreateStatusObject(utils.PrimaryMinecraftProtocolVersion, utils.PrimaryMinecraftVersion)
	go world.Init()
	if DEBUG {
		Log.Debug("Server initialised")
	}
}

func NewServer(ip string, port string, multicore bool, tick bool, lockosthread bool, reuse bool, sendBuf int, recvBuf int, readBufferCap int) (Server, error) {
	Log.Info("Generating Key chain")
	npacket.Keys()
	S := new(Server)
	GlobalServer = S
	err := gnet.Serve(S, "tcp://"+ip+port, gnet.WithMulticore(multicore), gnet.WithTicker(tick), gnet.WithLockOSThread(lockosthread), gnet.WithReusePort(reuse), gnet.WithSocketSendBuffer(sendBuf), gnet.WithSocketRecvBuffer(recvBuf), gnet.WithReadBufferCap(readBufferCap), gnet.WithTCPKeepAlive(5*time.Second))
	return *S, err
}

func (S *Server) OnInitComplete(Srv gnet.Server) (Action gnet.Action) {
	Log.Infof("HoneyBEE is listening on %s (multi-cores: %t, SO_REUSE: %t, Timeout: %d, loops: %d) ", Srv.Addr.String(), Srv.Multicore, Srv.ReusePort, Srv.TCPKeepAlive, Srv.NumEventLoop)
	return gnet.None
}

func (S *Server) OnOpened(Conn gnet.Conn) (Out []byte, Action gnet.Action) {
	Log.Infof("Socket with addr: %s has been opened...\n", Conn.RemoteAddr().String())
	C := new(Client)
	C.Name = Conn.RemoteAddr().String()
	C.Conn = Conn
	C.State = HANDSHAKE
	C.FrameChannel = make(chan []byte, 10)
	C.Close = make(chan bool)
	S.ConnectedSockets.Store(Conn.RemoteAddr().String(), C)
	Conn.SetContext(C)
	Log.Debug(Conn.RemoteAddr().String())
	go C.React(C.FrameChannel, C.Close) //the goroutine that does packet logic
	return
}

func (S *Server) OnClosed(Conn gnet.Conn, err error) (Action gnet.Action) {
	Log.Infof("Socket with addr: %s is closing...\n", Conn.RemoteAddr().String())
	S.ConnectedSockets.Delete(Conn.RemoteAddr().String())
	C, tmp := Conn.Context().(*Client)
	if tmp == false {
		Log.Critical("Conn Context is nil!")
	} else {
		C.Close <- true
		close(C.FrameChannel)
		close(C.Close)
	}
	Conn.SetContext(nil)
	Log.Infof("Socket with addr: %s is closed\n", Conn.RemoteAddr().String())
	return
}
