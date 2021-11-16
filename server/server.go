package server

import (
	"HoneyBEE/packet"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/panjf2000/gnet"
)

type Client struct {
	RemoteAddr      string
	PlayerName      string
	UUID            uuid.UUID
	Conn            gnet.Conn
	ProtocolVersion int32
	State           int
	Closed          bool
	ClosedMutex     sync.Mutex
	OptionalData    interface{}
}

type ServerCodec struct{}

///
///Optimise me, try not to use channels, maybe try a zero sized struct{} channel in worst case
///Do not use a seperate go routine for every client as per gnet I can use ants a go pool but this needs testing
///

func (S *Server) React(Frame []byte, Conn gnet.Conn) (Out []byte, Action gnet.Action) {
	//CC, tmp := S.ConnectedSockets.Load(Conn.RemoteAddr().String())
	Log.Debug("React hit!")
	ClientConn, tmp := Conn.Context().(*Client) //Get the client object from conn context
	if !tmp {
		Conn.Close()
		Log.Critical("Client Connection context was not Client object!")
		return nil, gnet.Close
	}
	if len(Frame) > 0 {
		Out = Frame //ClientConn.FrameChannel <- Frame
		Action = gnet.None
		return
	} else {
		Log.Warning("Frame is 0?!")
		return nil, gnet.Close
	}
	Conn.SetContext(ClientConn)
	//Out = Frame
	return nil, gnet.Close
}

func (SC *ServerCodec) Encode(c gnet.Conn, buf []byte) (out []byte, err error) {
	//Log.Debug("Encode hit! " /*buf*/)
	return buf, nil
}

func (SC *ServerCodec) Decode(c gnet.Conn) (out []byte, err error) {
	Log.Debug("Decode hit!")
	Frame := c.Read()
	c.ResetBuffer()
	ClientConn := c.Context().(*Client)
	var PR = packet.CreatePacketReader([]byte{0})
	//Get PacketSize and Data
	PacketSize, NR, err := packet.DecodeVarInt(Frame) //NR = Numread, used to note the position in the frame where it read to
	if err != nil {
		panic(err)
	}
	PacketDataSize := PacketSize - 1
	PacketID, NR2, err := packet.DecodeVarInt(Frame[NR:]) //NR2 is the second numread so the Decoder later on will correctly
	if err != nil {
		panic(err)
	}
	PR.Setdata(Frame[NR2+NR:])
	/*Size check - packets cannot be bigger than this which can lead to the server and client crashing
	also known as book banning or any item/block that is used to overload the packet limit*/
	if PacketSize > 2097151 {
		c.Close() //Disconnect the client until  I find a solution, my idea is a custom mod or client that raises the packet size to a Long
		return
	}
	//
	//Log.Debug("ClientConn ", ClientConn, "Conn: ", c.RemoteAddr().String())
	//Log.Debug("PSize: ", PacketSize, "PDS: ", PacketDataSize, " PID: ", PacketID, "State: ", ClientConn.State)
	//Make GeneralPacket
	GP := &packet.GeneralPacket{PacketSize, PacketID, &PR, nil} //new(packet.GeneralPacket)
	//Legacy Ping - drop conn
	if PacketSize == 0xFE && ClientConn.State == HANDSHAKE {
		c.Close()
		return
	}
	//Packet Logic
	switch ClientConn.State {
	case HANDSHAKE:
		switch PacketID {
		case 0x00:
			if PacketDataSize == 0 {
				Log.Critical("Packet ordering is whack yo, the bees flew into the glass")
				c.Close()
				return
			}
			HP := new(packet.Handshake_0x00)
			HP.Packet = GP
			err := HP.Decode()
			if err != nil {
				Log.Critical("Error while decoding HP: ", err)
				c.Close()
				return nil, err
			}
			ClientConn.ProtocolVersion = HP.ProtocolVersion
			ClientConn.State = int(HP.NextState)
			c.SetContext(ClientConn)
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
				c.AsyncWrite(writer.GetPacket())
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
			c.AsyncWrite(writer.GetPacket())
			//Log.Debug("WRITER NOTICE ME", writer.GetPacket())
		}
	case LOGIN:
		switch PacketID {
		case 0x00:
			Log.Debug("0x00_SB - Login start")
			LoginStart := new(packet.Login_0x00_SB)
			LoginStart.Packet = GP
			LoginStart.Decode()
			Log.Info("Name decoded: ", LoginStart.Name)
			ClientConn.PlayerName = LoginStart.Name
			// LERQ := new(packet.Login_0x01_CB)
			// PW := LERQ.Encode()
			// c.AsyncWrite(PW.GetPacket())
			// Log.Debug("Sent 0x01_CB - Encryption Request")
			PWC := packet.CreatePacketWriter(0x03)
			PWC.WriteVarInt(-1)
			c.AsyncWrite(PWC.GetPacket())
			PW := packet.CreatePacketWriter(0x02)                                    //LS.Encode()
			ClientConn.UUID = uuid.MustParse("523c4206-f96b-43ad-a220-9835508444d6") //"f41f9506-e8f7-4a2b-bd4d-4db421620bff")
			PW.WriteUUID(uuid.MustParse("523c4206-f96b-43ad-a220-9835508444d6"))     //"f41f9506-e8f7-4a2b-bd4d-4db421620bff"))
			PW.WriteString(ClientConn.PlayerName)
			c.AsyncWrite(PW.GetPacket()) //Send Login Success
			Log.Debug("Data: ", PW.GetPacket())
			Log.Debug("Sent 0x02_CB - Login Success")
			ClientConn.State = PLAY
			ClientConn.Conn.SetContext(ClientConn)
			time.Sleep(1 * time.Second)
			//
			//JG := new(packet.JoinGame_CB) //JoinGame - Play 0x26
			//PW := JG.Encode(LoginStart.Name, 0)
			c.AsyncWrite(JoinGamePacket) //PW.GetPacket())

			Log.Debug("Sent Join Game - 0x26 - P")

			PW.ResetData(0x18) //Server brand
			PW.WriteString("minecraft:brand")
			PW.WriteString("HoneyBEE")
			Log.Debug("Sent server brand")
			c.AsyncWrite(PW.GetPacket())

			// PW.ResetData(0x0E) //Difficulty
			// PW.WriteByte(0)
			// PW.WriteBoolean(true)
			// Log.Debug("Sent Difficulty")
			// c.AsyncWrite(PW.GetPacket())

			// PW.ResetData(0x32) //Player Abilities
			// PW.WriteByte(0x0F)
			// PW.WriteFloat(0.05)
			// PW.WriteFloat(0.1)
			// Log.Debug("Sent Player Abilities")
			// c.AsyncWrite(PW.GetPacket())

			PW.ResetData(0x48) //Held Item Change
			PW.WriteByte(0)
			Log.Debug("Sent held item change")
			c.AsyncWrite(PW.GetPacket())

			PW.ResetData(0x65) //Declare Recipes - investigate
			PW.WriteVarInt(0)
			Log.Debug("Sent recipe")
			c.AsyncWrite(PW.GetPacket())

			Log.Debug("Sent Tags")
			c.AsyncWrite(TagsPacket)

			PW.ResetData(0x1B) //Entity Status
			PW.WriteInt(2)
			PW.WriteByte(OP4)
			Log.Debug("Sent Entity Status")
			c.AsyncWrite(PW.GetPacket())

			// PW.ResetData(0x12) //Declare commands
			// PW.WriteVarInt(0)
			// PW.WriteVarInt(0)
			// Log.Debug("Sent Declare commands")
			// c.AsyncWrite(PW.GetPacket())

			Log.Debug("Sent Unlock Recipes")
			c.AsyncWrite(UnlockRecipesPacket) //Unlock Recipes

			PW.ResetData(0x38) //Player pos and look
			PW.WriteDouble(0.0)
			PW.WriteDouble(64.0)
			PW.WriteDouble(0.0)
			PW.WriteFloat(0.0)
			PW.WriteFloat(0.0)
			PW.WriteUnsignedByte(0)
			PW.WriteVarInt(0)
			PW.WriteBoolean(true)
			Log.Debug("Sent Player pos and look")
			c.AsyncWrite(PW.GetPacket())

			// PW.ResetData(0x0F) //Chat Message
			// CM := new(packet.ChatMessage_CB)
			// CM.Chat = new(jsonstruct.ChatComponent)
			// CM.Chat.Text = "Hello"
			// CMB := CM.Chat.MarshalChatComponent()
			// PW.WriteArray(CMB)
			// PW.WriteByte(1)
			// PW.WriteUUID(LS.UUID)
			// Log.Debug("Sent Chat message")
			// c.AsyncWrite(PW.GetPacket())

			PW.ResetData(0x36) //PlayerInfo
			PW.WriteVarInt(0)
			PW.WriteVarInt(1)
			PW.WriteUUID(ClientConn.UUID)
			PW.WriteString(ClientConn.PlayerName)
			PW.WriteVarInt(0)
			PW.WriteVarInt(1)
			PW.WriteVarInt(0)
			PW.WriteBoolean(false)
			Log.Debug("Sent Player info")
			c.AsyncWrite(PW.GetPacket())

			PW.ResetData(0x49) //Update view position
			PW.WriteVarInt(0)
			PW.WriteVarInt(0)
			Log.Debug("Sent view position")
			c.AsyncWrite(PW.GetPacket())

			PW.ResetData(0x4A) //Update view distance
			PW.WriteVarInt(12)
			Log.Debug("Sent view distance")
			c.AsyncWrite(PW.GetPacket())

			// PW.ResetData(0x4D) //Entity metadata
			// PW.WriteVarInt(2)
			// Log.Debug("Sent Entity Metadata")
			// c.AsyncWrite(PW.GetPacket())

			PW.ResetData(0x58) //Time Update
			PW.WriteLong(0)
			PW.WriteLong(-12000)
			Log.Debug("Sent Time update")
			c.AsyncWrite(PW.GetPacket())

			PW.ResetData(0x4B) //Spawn Position
			PW.WritePosition(0, 0, 64)
			PW.WriteFloat(0.0)
			Log.Debug("Sent Spawn position")
			c.AsyncWrite(PW.GetPacket())

			go ClientConn.ChunkLoad() //Load chunks
			PW.ResetData(0x38)        //Player pos and look
			PW.WriteDouble(0.0)
			PW.WriteDouble(64.0)
			PW.WriteDouble(0.0)
			PW.WriteFloat(0.0)
			PW.WriteFloat(0.0)
			PW.WriteByte(0)
			PW.WriteVarInt(0)
			PW.WriteBoolean(true)
			Log.Debug("Sent Player pos and look")
			c.AsyncWrite(PW.GetPacket())

			PW.ResetData(0x14) //Window Items
			PW.WriteUnsignedByte(0)
			PW.WriteVarInt(1)
			PW.WriteVarInt(46)
			PW.WriteArray(make([]byte, 46))
			PW.WriteByte(0)
			Log.Debug("Sent Window Items")
			c.AsyncWrite(PW.GetPacket())

			PW.ResetData(0x52) //Update Health
			PW.WriteFloat(20.0)
			PW.WriteVarInt(20)
			PW.WriteFloat(5.0)
			Log.Debug("Sent Update Health")
			c.AsyncWrite(PW.GetPacket())

			PW.ResetData(0x51) //Set expereince
			PW.WriteFloat(1.0)
			PW.WriteVarInt(1)
			PW.WriteVarInt(7)
			Log.Debug("Sent Set Experience")
			c.AsyncWrite(PW.GetPacket())
		case 0x01:
			// Log.Debug("Sent 0x01_SB - Encryption response")
			// LERSP := new(packet.Login_0x01_SB)
			// LERSP.Packet = GP
			// LERSP.Decode()
			// //Login Success
			// LS := new(packet.Login_0x02_CB)
			// LS.UUID = packet.Auth(ClientConn.PlayerName, LERSP.SharedSecret)
			// if LS.UUID == uuid.Nil {
			// 	c.Close()
			// 	return
			// }
			// //time.Sleep(3 * time.Second)
			// LS.Username = ClientConn.PlayerName
			// PW := packet.CreatePacketWriter(0x02) //LS.Encode()
			// PW.WriteUUID(uuid.MustParse("f41f9506-e8f7-4a2b-bd4d-4db421620bff"))
			// PW.WriteString(ClientConn.PlayerName)
			// c.AsyncWrite(PW.GetPacket()) //Send Login Success
			// Log.Debug("Sent 0x02_CB - Login Success")
			// time.Sleep(3 * time.Second)
			// ClientConn.UUID = LS.UUID
			// ClientConn.State = PLAY
			// c.SetContext(ClientConn) //Set conn context

			JG := new(packet.JoinGame_CB) //JoinGame - Play 0x26
			PW := JG.Encode(ClientConn.PlayerName, 0)
			c.AsyncWrite(PW.GetPacket())
			Log.Debug("Sent Join Game - 0x26 - P")

			// PW.ResetData(0x48) //Server brand
			// PW.WriteString("minecraft:brand")
			// PW.WriteString("HoneyBEE")
			// Log.Debug("Sent held item change")
			// c.AsyncWrite(PW.GetPacket())

			// PW.ResetData(0x0E) //Difficulty
			// PW.WriteByte(0)
			// PW.WriteBoolean(true)
			// Log.Debug("Sent Difficulty")
			// c.AsyncWrite(PW.GetPacket())

			// PW.ResetData(0x32) //Player Abilities
			// PW.WriteByte(0x0F)
			// PW.WriteFloat(0.05)
			// PW.WriteFloat(0.1)
			// Log.Debug("Sent Player Abilities")
			// c.AsyncWrite(PW.GetPacket())

			PW.ResetData(0x48) //Held Item Change
			PW.WriteByte(0)
			Log.Debug("Sent held item change")
			c.AsyncWrite(PW.GetPacket())

			PW.ResetData(0x65) //Declare Recipes - investigate
			PW.WriteVarInt(0)
			Log.Debug("Sent recipe")
			c.AsyncWrite(PW.GetPacket())

			Log.Debug("Sent Tags")
			c.AsyncWrite(TagsPacket)

			PW.ResetData(0x1B) //Entity Status
			PW.WriteInt(2)
			PW.WriteUnsignedByte(28)
			Log.Debug("Sent Entity Status")
			c.AsyncWrite(PW.GetPacket())

			// PW.ResetData(0x12) //Declare commands
			// PW.WriteVarInt(0)
			// PW.WriteVarInt(0)
			// Log.Debug("Sent Declare commands")
			// c.AsyncWrite(PW.GetPacket())

			Log.Debug("Sent Unlock Recipes")
			c.AsyncWrite(UnlockRecipesPacket) //Unlock Recipes

			PW.ResetData(0x38) //Player pos and look
			PW.WriteDouble(0.0)
			PW.WriteDouble(64.0)
			PW.WriteDouble(0.0)
			PW.WriteFloat(0.0)
			PW.WriteFloat(0.0)
			PW.WriteUnsignedByte(0)
			PW.WriteVarInt(0)
			PW.WriteBoolean(true)
			Log.Debug("Sent Player pos and look")
			c.AsyncWrite(PW.GetPacket())

			// PW.ResetData(0x0F) //Chat Message
			// CM := new(packet.ChatMessage_CB)
			// CM.Chat = new(jsonstruct.ChatComponent)
			// CM.Chat.Text = "Hello"
			// CMB := CM.Chat.MarshalChatComponent()
			// PW.WriteArray(CMB)
			// PW.WriteByte(1)
			// PW.WriteUUID(LS.UUID)
			// Log.Debug("Sent Chat message")
			// c.AsyncWrite(PW.GetPacket())

			PW.ResetData(0x36) //PlayerInfo
			PW.WriteVarInt(0)
			PW.WriteVarInt(1)
			PW.WriteUUID(ClientConn.UUID)
			PW.WriteString(ClientConn.PlayerName)
			PW.WriteVarInt(0)
			PW.WriteVarInt(1)
			PW.WriteVarInt(0)
			PW.WriteBoolean(false)
			Log.Debug("Sent Player info")
			c.AsyncWrite(PW.GetPacket())

			PW.ResetData(0x49) //Update view position
			PW.WriteVarInt(0)
			PW.WriteVarInt(0)
			Log.Debug("Sent view position")
			c.AsyncWrite(PW.GetPacket())

			PW.ResetData(0x4A) //Update view distance
			PW.WriteVarInt(12)
			Log.Debug("Sent view distance")
			c.AsyncWrite(PW.GetPacket())

			PW.ResetData(0x4D) //Update view position
			PW.WriteVarInt(2)
			Log.Debug("Sent Entity Metadata")
			c.AsyncWrite(PW.GetPacket())

			//ClientConn.ChunkLoad() //Load chunks

			PW.ResetData(0x58) //Time Update
			PW.WriteLong(0)
			PW.WriteLong(-12000)
			Log.Debug("Sent Time update")
			c.AsyncWrite(PW.GetPacket())

			PW.ResetData(0x4B) //Spawn Position
			PW.WritePosition(0, 0, 64)
			PW.WriteFloat(0.0)
			Log.Debug("Sent Spawn position")
			c.AsyncWrite(PW.GetPacket())

			PW.ResetData(0x38) //Player pos and look
			PW.WriteDouble(0.0)
			PW.WriteDouble(64.0)
			PW.WriteDouble(0.0)
			PW.WriteFloat(0.0)
			PW.WriteFloat(0.0)
			PW.WriteByte(0)
			PW.WriteVarInt(0)
			PW.WriteBoolean(true)
			Log.Debug("Sent Player pos and look")
			c.AsyncWrite(PW.GetPacket())

			PW.ResetData(0x14) //Window Items
			PW.WriteUnsignedByte(0)
			PW.WriteVarInt(1)
			PW.WriteVarInt(46)
			PW.WriteArray(make([]byte, 46))
			PW.WriteByte(0)
			Log.Debug("Sent Window Items")
			c.AsyncWrite(PW.GetPacket())

			PW.ResetData(0x52) //Update Health
			PW.WriteFloat(20.0)
			PW.WriteVarInt(20)
			PW.WriteFloat(5.0)
			Log.Debug("Sent Update Health")
			c.AsyncWrite(PW.GetPacket())

			PW.ResetData(0x51) //Set expereince
			PW.WriteFloat(1.0)
			PW.WriteVarInt(1)
			PW.WriteVarInt(7)
			Log.Debug("Sent Set Experience")
			c.AsyncWrite(PW.GetPacket())

		case 0x02:
			Log.Debug("Play 0x02_SB")
		case 0x03:
			Log.Debug("Recived 0x03")
		case 0x04:
			Log.Debug("Recived 0x04")
		default:
			Log.Debug("Recieved packet login")
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
	return nil, nil
}
