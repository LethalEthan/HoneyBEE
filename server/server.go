package server

import (
	"HoneyBEE/packet"

	"github.com/panjf2000/gnet"
)

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

///
///Optimise me, try not to use channels, maybe try a zero sized struct{} channel in worst case
///Do not use a seperate go routine for every client as per gnet I can use ants a go pool but this needs testing
///

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
	//Log.Critical("Sent FrameChannel")
	return
}

/*React - This continuously listens on FrameChan for frames and applies the logic
it listens continuously to make sure packets are in sequence by using a channel*/
func (ClientConn *Client) React(FrameChan chan []byte, Close chan bool) {
	//
	for {
		select {
		case Frame := <-FrameChan:
			//Log.Debug("RECV Frame")
			if len(Frame) == 0 {
				ClientConn.Conn.Close()
				return
			}
			//Get PacketSize and Data
			PacketSize, NR, err := packet.DecodeVarInt(Frame) //NR = Numread, used to note the position in the frame where it read to
			PacketDataSize := PacketSize - 1
			PacketID, NR2, err := packet.DecodeVarInt(Frame[NR:]) //NR2 is the second numread so the Decoder later on will correctly
			if err != nil {
				panic(err)
			}
			/*Size check - packets cannot be bigger than this which can lead to the server and client crashing
			also known as book banning or any item/block that is used to overload the packet limit*/
			if PacketSize > 2097151 {
				ClientConn.Conn.Close() //Disconnect the client until  I find a solution, my idea is a custom mod or client that raises the packet size to a Long
			}
			//
			if DEBUG {
				Log.Debug("ClientConn ", ClientConn, "Conn: ", ClientConn.Conn.RemoteAddr().String())
				Log.Debug("PSize: ", PacketSize, "PDS: ", PacketDataSize, " PID: ", PacketID, "State: ", ClientConn.State)
			}
			//Make GeneralPacket
			GP := new(packet.GeneralPacket)
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
					HP := new(packet.Handshake_0x00)
					HP.Packet = GP
					err := HP.Decode()
					if err != nil {
						ClientConn.Conn.Close()
					}
					ClientConn.ProtocolVersion = HP.ProtocolVersion
					ClientConn.State = int(HP.NextState)
					ClientConn.Conn.SetContext(ClientConn)
				}
			case STATUS:
				switch PacketID {
				case 0x00:
					Log.Debug("status 0x00_SB")
					if PacketSize == 1 {
						SP := new(packet.Stat_Response)
						SP.ProtocolVersion = ClientConn.ProtocolVersion
						writer, err := SP.Encode()
						if err != nil {
							panic(err)
						}
						ClientConn.Conn.AsyncWrite(writer.GetPacket())
					} else {
						Log.Warning("Packet is bigger than expected")
					}
				case 0x01:
					Log.Debug("status 0x01_SB")
					StatP := new(packet.Stat_Ping)
					GP.OptionalData = StatP.Ping
					StatP.Packet = GP
					StatP.Decode()
					StatPClient := new(packet.Stat_Pong)
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
					LoginStart := new(packet.Login_0x00_SB)
					LoginStart.Packet = GP
					LoginStart.Decode()
					Log.Info("Name decoded: ", LoginStart.Name)
					ClientConn.PlayerName = LoginStart.Name
					LERQ := new(packet.Login_0x01_CB)
					PW := LERQ.Encode()
					ClientConn.Conn.AsyncWrite(PW.GetPacket())
					Log.Debug("Sent 0x01_CB")
				case 0x01:
					Log.Debug("Play 0x01_SB") //Login Response
					LERSP := new(packet.Login_0x01_SB)
					LERSP.Packet = GP
					LERSP.Decode()
					//Login Success
					LS := new(packet.Login_0x02_CB)
					LS.UUID = packet.Auth(ClientConn.PlayerName, LERSP.SharedSecret)
					if LS.UUID == "" {
						ClientConn.Conn.Close()
					}
					LS.Username = ClientConn.PlayerName
					PW := LS.Encode(ClientConn.PlayerName)
					ClientConn.Conn.AsyncWrite(PW.GetPacket()) //Send Login Success
					Log.Debug("Sent 0x02_CB")
					Log.Debug(PW.GetData())
					//Set to Play state
					ClientConn.State = PLAY
					ClientConn.Conn.SetContext(ClientConn) //Set conn context
					//JoinGame
					JG := new(packet.JoinGame_CB)
					PW = JG.Encode(LS.UUID, LS.Username, 0, ClientConn.Conn)
					ClientConn.Conn.AsyncWrite(PW.GetPacket())
					/*
						//Plugin Message - brand
						PW = packet.CreatePacketWriterWithCapacity(0x04, 24)
						PW.WriteVarInt(0x0F)
						PW.WriteIdentifier("minecraft:brand")
						PW.WriteArray([]byte("HoneyBEE!"))
						ClientConn.Conn.AsyncWrite(PW.GetPacket())
						//
						PM := new(packet.PluginMessage_CB)
						PW = PM.Encode()
						ClientConn.Conn.AsyncWrite(PW.GetPacket())
						Log.Debug("Sent PM: ", PW.GetPacket())
						//
						PW = packet.CreatePacketWriterWithCapacity(0x0E, 10)
						PW.WriteByte(0)
						PW.WriteBoolean(false)
						ClientConn.Conn.AsyncWrite(PW.GetPacket())
						Log.Debug("Sent Diff")
						//
						PW = packet.CreatePacketWriterWithCapacity(0x32, 10)
						PW.WriteByte(0x01)
						PW.WriteFloat(0.05)
						PW.WriteBoolean(false)
						ClientConn.Conn.AsyncWrite(PW.GetPacket())
						Log.Debug("Sent PA")
						//
						PW = packet.CreatePacketWriterWithCapacity(0x58, 24)
						PW.WriteLong(100000)
						PW.WriteLong(1000)
						ClientConn.Conn.AsyncWrite(PW.GetPacket())
						Log.Debug("TIME: ", PW.GetPacket())
						Log.Debug("Sent Time")*/
				default:
					Log.Debug("Recieved some shit")
				case 0x02:
					Log.Debug("Play 0x02_SB")
				case 0x03:
					Log.Debug("Recived 0x03")
				case 0x04:
					Log.Debug("Recived 0x04")
				}
			case PLAY:
				switch PacketID {
				case 0x00:
					Log.Info("Recieved 0x00")
				case 0x05:
					Log.Info("Recieved 0x05")
				case 0x0A:
					Log.Info("Recieved 0x0A")
				default:
					Log.Debug("Recieved packet play:", PacketID)
				}
			default:
				Log.Info("PP")
			}
		case <-Close:
			Log.Debug("Closing client conn routine")
			return
		}
	}
}