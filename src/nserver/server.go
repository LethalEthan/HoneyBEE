package nserver

import (
	"npacket"
	"sync"
	"time"

	logging "github.com/op/go-logging"
	"github.com/panjf2000/gnet"
)

var Log = logging.MustGetLogger("HoneyGO")
var DEBUG = true
var GlobalServer *Server

type Server struct {
	*gnet.EventServer
	//pool *goroutine.Pool
	ConnectedSockets sync.Map //PROPOSAL: Create map standard that can be used concurrently, writes and reads; queued based on time sent
}

type Client struct {
	Name            string
	Conn            gnet.Conn
	ProtocolVersion int32
	PlayerName      string
	State           int
	OptionalData    interface{}
	FrameChannel    chan []byte
	Close           chan bool
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

func (S *Server) React(Frame []byte, Conn gnet.Conn) (Out []byte, Action gnet.Action) {
	//CC, tmp := S.ConnectedSockets.Load(Conn.RemoteAddr().String())
	ClientConn, tmp := Conn.Context().(*Client) //Get the client object from conn context
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
	for {
		select {
		case Frame := <-FrameChan:
			Log.Debug("RECV Frame")
			if len(Frame) == 0 {
				ClientConn.Conn.Close()
				return
			}
			//Get PacketSize and Data
			PacketSize, NR, err := npacket.DecodeVarInt(Frame) //NR = Numread, used to note the position in the frame where it read to
			PacketDataSize := PacketSize - 1
			PacketID, NR2, err := npacket.DecodeVarInt(Frame[NR:]) //NR2 is the second numread so the Decoder later on will correctly
			if err != nil {
				panic(err)
			}
			/*Size check - packets cannot be bigger than this which can lead to the server and client crashing
			also known as book banning or any item/block that is used to overload the packet limit*/
			if PacketSize > 2097151 {
				ClientConn.Conn.Close()
			}
			//
			if DEBUG {
				Log.Debug("ClientConn ", ClientConn, "Conn: ", ClientConn.Conn.RemoteAddr().String())
				Log.Debug("PSize: ", PacketSize, "PDS: ", PacketDataSize, " PID: ", PacketID, "NR: ", NR, "NR2: ", NR2)
			}
			//Make GeneralPacket
			GP := new(npacket.GeneralPacket)
			GP.PacketSize = PacketSize
			GP.PacketID = PacketID
			GP.PacketData = Frame[NR2+NR:] //Uses the Numreads from earlier to correctly know where the data starts since VarInts are variable in size
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
					if PacketDataSize == 0 {
						Log.Critical("Packet ordering is whack yo, the bees flew into the glass")
						ClientConn.Conn.Close()
					}
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
						SP := new(npacket.Stat_Response)
						SP.ProtocolVersion = ClientConn.ProtocolVersion
						writer := SP.Encode()
						ClientConn.Conn.AsyncWrite(writer.GetPacket())
					}
				case 0x01:
					Log.Debug("status 0x01_SB")
					StatP := new(npacket.Stat_Ping)
					GP.OptionalData = StatP.Ping
					StatP.Packet = GP
					StatP.Decode()
					StatPClient := new(npacket.Stat_Pong)
					StatPClient.Packet = GP
					StatPClient.Pong = StatP.Ping
					writer := StatPClient.Encode(StatP.Ping)
					ClientConn.Conn.AsyncWrite(writer.GetPacket())
					Log.Debug("WRITER NOTICE ME", writer.GetPacket())
				}
			case LOGIN:
				switch PacketID {
				case 0x00:
					Log.Debug("Play 0x00_SB")
					LoginStart := new(npacket.Login_0x00_SB)
					LoginStart.Packet = GP
					LoginStart.Decode()
					Log.Info("Name decoded: ", LoginStart.Name)
					ClientConn.PlayerName = LoginStart.Name
					LERQ := new(npacket.Login_0x01_CB)
					PW := LERQ.Encode()
					ClientConn.Conn.AsyncWrite(PW.GetPacket())
					Log.Debug("Sent 0x01_CB")
				case 0x01:
					Log.Debug("Play 0x01_SB")
					LERSP := new(npacket.Login_0x01_SB)
					LERSP.Packet = GP
					LERSP.Decode()
					LS := new(npacket.Login_0x02_CB)
					LS.UUID = npacket.Auth(ClientConn.PlayerName, LERSP.SharedSecret)
					LS.Username = ClientConn.PlayerName
					PW := LS.Encode(ClientConn.PlayerName)
					ClientConn.Conn.AsyncWrite(PW.GetPacket())
					Log.Debug("Sent 0x02_CB")
					JG := new(npacket.JoinGame_CB)
					JG.Encode(ClientConn.Conn)
				case 0x02:
					Log.Debug("Play 0x02_SB")
				}
			case PLAY:
			switch PacketID {
			case 0x00:

				}
			}
		case <-Close:
			Log.Debug("Closing client conn routine")
			return
		}
	}
}

func (S *Server) Shutdown() {
	S.EventServer = nil
	GlobalServer = nil
	return
}
