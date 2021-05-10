package nserver

import (
	"npacket"
	"time"

	logging "github.com/op/go-logging"
	"github.com/panjf2000/gnet"
)

var Log = logging.MustGetLogger("HoneyGO")
var DEBUG = true

type Server struct {
	*gnet.EventServer
	//pool *goroutine.Pool
	//ConnectedSockets sync.Map //PROPOSAL: Create map standard that can be used concurrently and uses eventual consistency?
}

type Client struct {
	Name            string
	Conn            gnet.Conn
	ProtocolVersion int32
	State           int
	OptionalData    interface{}
	FrameChannel    chan []byte
	Close           chan bool
}

func NewServer(ip string, port string, multicore bool, tick bool, reuse bool, sendBuf int, recvBuf int, readBufferCap int) Server {
	S := new(Server)
	gnet.Serve(S, "tcp://"+ip+port, gnet.WithMulticore(multicore), gnet.WithTicker(tick), gnet.WithLockOSThread(false), gnet.WithReusePort(reuse), gnet.WithSocketSendBuffer(sendBuf), gnet.WithSocketRecvBuffer(recvBuf), gnet.WithReadBufferCap(readBufferCap), gnet.WithTCPKeepAlive(5*time.Second))
	return *S
}

func (S *Server) OnInitComplete(Srv gnet.Server) (Action gnet.Action) {
	Log.Infof("HoneyGO is listening on %s (multi-cores: %t, SO_REUSE: %t, Timeout: %d, loops: %d) ", Srv.Addr.String(), Srv.Multicore, Srv.ReusePort, Srv.TCPKeepAlive, Srv.NumEventLoop)
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
	//S.ConnectedSockets.Store(Conn.RemoteAddr().String(), C)
	Conn.SetContext(C)
	Log.Debug(Conn.RemoteAddr().String())
	go C.React(C.FrameChannel, C.Close) //the goroutine that does packet logic
	return
}

func (S *Server) OnClosed(Conn gnet.Conn, err error) (Action gnet.Action) {
	Log.Infof("Socket with addr: %s is closing...\n", Conn.RemoteAddr().String())
	//S.ConnectedSockets.Delete(Conn.RemoteAddr().String())
	C := Conn.Context().(*Client)
	C.Close <- true
	close(C.FrameChannel)
	close(C.Close)
	Conn.SetContext(nil)
	Log.Infof("Socket with addr: %s is closed\n", Conn.RemoteAddr().String())
	return
}

func (S *Server) React(Frame []byte, Conn gnet.Conn) (Out []byte, Action gnet.Action) {
	ClientConn, tmp := Conn.Context().(*Client) //, tmp := S.ConnectedSockets.Load(Conn.RemoteAddr().String())
	if tmp == false {
		Conn.Close()
		Log.Critical("Client Connection context was not Client object!")
		return
	}
	if len(Frame) > 0 {
		ClientConn.FrameChannel <- Frame
	} else {
		Log.Warning("Frame is 0?!")
	}
	Conn.SetContext(ClientConn)
	Log.Critical("Sent FrameChannel")
	return
}

/*React - This continuously listens on FrameChan for frames and applies the logic
it listens continuously to make sure packets are in sequence by using a channel*/
func (ClientConn *Client) React(FrameChan chan []byte, Close chan bool) {
	//
	defer ClientConn.Conn.Close()
	for {
		select {
		case Frame := <-FrameChan:
			//Frame := <-FrameChan
			if len(Frame) == 0 {
				ClientConn.Conn.Close()
				return
			}
			//Get PacketSize and Data
			PacketSize, NR, err := npacket.DecodeVarInt(Frame)
			PacketDataSize := PacketSize - 1
			PacketID, NR2, err := npacket.DecodeVarInt(Frame[NR:])
			if err != nil {
				panic("error")
			}
			//Size check
			if PacketSize > 2097151 {
				ClientConn.Conn.Close()
			}
			//
			if DEBUG {
				Log.Debug("ClientConn ", ClientConn, "Conn: ", ClientConn.Conn.RemoteAddr().String())
				Log.Debug("PSize: ", PacketSize, "PDS: ", PacketDataSize, " PID: ", PacketID)
			}
			//Make GeneralPacket
			GP := new(npacket.GeneralPacket)
			GP.PacketSize = PacketSize
			GP.PacketID = PacketID
			GP.PacketData = Frame[NR2+NR:]
			Log.Debug("Frame: ", Frame)
			//Legacy Ping - drop conn
			if PacketSize == 0xFE && ClientConn.State == HANDSHAKE {
				ClientConn.Conn.Close()
			}
			//Packet Logic
			switch ClientConn.State {
			case HANDSHAKE:
				switch PacketID {
				case 0x00:
					HP := new(npacket.Handshake_0x00)
					HP.Packet = GP
					HP.Decode()
					ClientConn.ProtocolVersion = HP.ProtocolVersion
					ClientConn.State = int(HP.NextState)
					ClientConn.Conn.SetContext(ClientConn)
				}
			case STATUS:
				switch PacketID {
				case 0x00:
					Log.Debug("status 0x00_SB")
					if PacketSize == 1 {
						SP := new(npacket.Status_0x00_CB)
						SP.ProtocolVersion = ClientConn.ProtocolVersion
						writer := SP.Encode()
						ClientConn.Conn.AsyncWrite(writer.GetPacket())
						//Conn.Close()
						//return
					}
				case 0x01:
					Log.Debug("status 0x01_SB")
					StatP := new(npacket.Status_0x01_SB)
					GP.OptionalData = StatP.Ping
					StatP.Packet = GP
					StatP.Decode()
					StatPClient := new(npacket.Status_0x01_CB)
					StatPClient.Packet = GP
					StatPClient.Pong = StatP.Ping
					writer := StatPClient.Encode()
					ClientConn.Conn.AsyncWrite(writer.GetPacket())
					Log.Debug("WRITER NOTICE ME", writer.GetPacket())
				}
			}
		case <-Close:
			return
		}
	}
}
