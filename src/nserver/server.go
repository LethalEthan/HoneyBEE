package nserver

import (
	"Packet"
	"fmt"
	"npacket"
	"server"

	logging "github.com/op/go-logging"
	"github.com/panjf2000/gnet"
	"github.com/pquerna/ffjson/ffjson"
)

var Log = logging.MustGetLogger("HoneyGO")
var DEBUG = true

type Server struct {
	*gnet.EventServer
	//ConnectedSockets sync.Map //PROPOSAL: Create map standard that can be used concurrently and uses eventual consistency?
}

type Client struct {
	Name            string
	Conn            gnet.Conn
	ProtocolVersion int32
	State           int
	OptionalData    interface{}
}

func NewServer() {
	S := new(Server)
	//Conn := "tcp://" + config.yadda
	fmt.Println("bruh")
	gnet.Serve(S, "tcp://:25560", gnet.WithMulticore(true), gnet.WithTicker(false), gnet.WithLockOSThread(false))
	fmt.Println("bruh")
}

func (S *Server) OnInitComplete(Srv gnet.Server) (Action gnet.Action) {
	Log.Infof("HoneyGO is listening on %s (multi-cores: %t, loops: %d) ", Srv.Addr.String(), Srv.Multicore, Srv.NumEventLoop)
	return gnet.None
}

func (S *Server) OnOpened(Conn gnet.Conn) (Out []byte, Action gnet.Action) {
	Log.Infof("Socket with addr: %s has been opened...\n", Conn.RemoteAddr().String())
	C := new(Client)
	C.Name = Conn.RemoteAddr().String()
	C.Conn = Conn
	C.State = HANDSHAKE
	//S.ConnectedSockets.Store(Conn.RemoteAddr().String(), C)
	Conn.SetContext(C)
	Log.Debug(Conn.RemoteAddr().String())
	return
}

func (S *Server) OnClosed(Conn gnet.Conn, err error) (Action gnet.Action) {
	Log.Infof("Socket with addr: %s is closing...\n", Conn.RemoteAddr().String())
	//S.ConnectedSockets.Delete(Conn.RemoteAddr().String())
	return
}

func (S *Server) React(Frame []byte, Conn gnet.Conn) (Out []byte, Action gnet.Action) {
	//Get PacketSize and Data
	ClientConn := Conn.Context().(*Client) //, tmp := S.ConnectedSockets.Load(Conn.RemoteAddr().String())
	PacketSize, NR, err := npacket.DecodeVarInt(Frame)
	PacketDataSize := PacketSize - 1
	PacketID, NR2, err := npacket.DecodeVarInt(Frame[NR:])
	if err != nil {
		panic("error")
	}
	if DEBUG {
		Log.Debug("ClientConn ", ClientConn, "Conn: ", Conn.RemoteAddr().String())
		Log.Debug("PSize: ", PacketSize, "PDS: ", PacketDataSize, " PID: ", PacketID)
	}
	//
	GP := new(npacket.GeneralPacket)
	GP.PacketSize = PacketSize
	GP.PacketID = PacketID
	GP.PacketData = Frame[NR2+NR:]
	Log.Debug("Frame: ", Frame)
	//
	if PacketSize == 0xFE && ClientConn.State == HANDSHAKE {
		Conn.Close()
	}
	//
	switch ClientConn.State {
	case HANDSHAKE:
		switch PacketID {
		case 0x00:
			HP := new(npacket.Handshake_0x00)
			HP.Packet = GP
			HP.Decode()
			ClientConn.State = int(HP.NextState)
			Conn.SetContext(ClientConn)
		}
	case STATUS:
		switch PacketID {
		case 0x00:
			Log.Debug("status 0x00_SB")
			if PacketDataSize == 1 {
				writer := Packet.CreatePacketWriter(0x00)
				server.CreateStatusObject(ClientConn.ProtocolVersion, "1.16.5")
				marshaledStatus, err := ffjson.Marshal(server.CurrentStatus) //Sends status via json
				if err != nil {
					Log.Error(err)
					Conn.Close()
					return
				}
				writer.WriteString(string(marshaledStatus))
				Conn.AsyncWrite(writer.Data)
				//} else {
				//Conn.Close()
				return
			}
		case 0x01:
			Log.Debug("status 0x01_SB")
			StatP := new(npacket.Status_0x01_SB)
			StatP.Packet = GP
			StatP.Decode()
			StatPClient := new(npacket.Status_0x01_CB)
			StatPClient.Packet.OptionalData = StatP.Ping
			StatPClient.Pong = StatP.Ping
			PW := StatPClient.Encode()
			Conn.AsyncWrite(PW.Data)
		}
	}
	//PR := npacket.CreatePacketReader(Frame)
	return
}
