package server

import (
	"HoneyBEE/packet"
	"HoneyBEE/player"
	"time"

	"github.com/google/uuid"
	"github.com/panjf2000/gnet"
)

type Client struct {
	RemoteAddr      string
	Player          player.PlayerObject
	PW              packet.PacketWriter
	PR              packet.PacketReader
	Conn            gnet.Conn
	ProtocolVersion int32
	State           int
	Read            chan []byte
	OptionalData    interface{}
}

type Read struct{}

func (S *Server) React(Frame []byte, Conn gnet.Conn) (Out []byte, Action gnet.Action) {
	CC, _ := Conn.Context().(*Client) //S.ConnectedSockets.Load(Conn.RemoteAddr().String())
	Log.Debug("React hit!")
	if len(Frame) > 0 {
		CC.Read <- Frame
		return nil, gnet.None
	} else {
		Log.Warning("Frame is 0?!")
		return nil, gnet.Close
	}
}

func (Client *Client) ClientReact(c gnet.Conn) {
	var PR = &Client.PR //packet.CreatePacketReader([]byte{0})
	var PW = &Client.PW
	for {
		Frame := <-Client.Read //Block until react fires
		if Frame == nil {
			Log.Debug("Frame nil, closing")
			return
		}
		// Log.Debug("FRAME: ", Frame)
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
		//Log.Debug("Client ", Client, "Conn: ", c.RemoteAddr().String())
		//Log.Debug("PSize: ", PacketSize, "PDS: ", PacketDataSize, " PID: ", PacketID, "State: ", Client.State)
		// Legacy Ping - drop conn
		if PacketSize == 0xFE && Client.State == HANDSHAKE {
			c.Close()
			return
		}
		// Packet Logic
		switch Client.State {
		case HANDSHAKE:
			switch PacketID {
			case 0x00:
				if PacketDataSize == 0 {
					Log.Critical("Packet ordering is whack yo, the bees flew into the glass")
					c.Close()
					return
				}
				HP := new(packet.Handshake_0x00)
				if err := HP.Decode(PR); err != nil {
					Log.Critical("Error while decoding HP: ", err)
					c.Close()
					return
				}
				Client.ProtocolVersion = HP.ProtocolVersion
				if Client.ProtocolVersion != 756 {
					c.Close()
				}
				Client.State = int(HP.NextState)
				c.SetContext(Client)
			}
		case STATUS:
			switch PacketID {
			case 0x00:
				Log.Debug("status 0x00_SB")
				if PacketSize == 1 {
					SP := new(packet.Stat_Response)
					SP.ProtocolVersion = Client.ProtocolVersion
					SP.Encode(PW)
					c.AsyncWrite(PW.GetPacket())
				} else {
					Log.Warning("Packet is bigger than expected")
				}
			case 0x01:
				Log.Debug("status 0x01_SB")
				StatP := new(packet.Stat_Ping)
				if err := StatP.Decode(PR); err != nil {
					Log.Error(err)
					c.Close()
					return
				}
				StatPClient := new(packet.Stat_Pong)
				StatPClient.Pong = StatP.Ping
				StatPClient.Encode(PW)
				c.AsyncWrite(PW.GetPacket())
			}
		case LOGIN:
			switch PacketID {
			case 0x00:
				Log.Debug("0x00_SB - Login start")
				LoginStart := new(packet.Login_0x00_SB)
				if err := LoginStart.Decode(PR); err != nil {
					Log.Error(err)
					c.Close()
					return
				}
				Log.Info("Name decoded: ", LoginStart.Name)
				Client.Player.PlayerName = LoginStart.Name
				// LERQ := new(packet.Login_0x01_CB)
				// LERQ.Encode()
				// c.AsyncWrite(PW.GetPacket())
				// Log.Debug("Sent 0x01_CB - Encryption Request")
				PWC := packet.CreatePacketWriter(0x03)
				PWC.WriteVarInt(-1)
				c.AsyncWrite(PWC.GetPacket())
				PW := packet.CreatePacketWriter(0x02)                                       //LS.Encode()
				Client.Player.UUID = uuid.MustParse("523c4206-f96b-43ad-a220-9835508444d6") //"f41f9506-e8f7-4a2b-bd4d-4db421620bff")
				PW.WriteUUID(uuid.MustParse("523c4206-f96b-43ad-a220-9835508444d6"))        //"f41f9506-e8f7-4a2b-bd4d-4db421620bff"))
				PW.WriteString(Client.Player.PlayerName)
				c.AsyncWrite(PW.GetPacket()) //Send Login Success
				Log.Debug("Data: ", PW.GetPacket())
				Log.Debug("Sent 0x02_CB - Login Success")
				Client.State = PLAY
				c.SetContext(Client)
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
				PW.WriteUUID(Client.Player.UUID)
				PW.WriteString(Client.Player.PlayerName)
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

				ChunkLoad(c)       //Load chunks
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
			case 0x01:
				// Log.Debug("Sent 0x01_SB - Encryption response")
				// LERSP := new(packet.Login_0x01_SB)
				// LERSP.Packet = GP
				// LERSP.Decode()
				// //Login Success
				// LS := new(packet.Login_0x02_CB)
				// LS.UUID = packet.Auth(Client.PlayerName, LERSP.SharedSecret)
				// if LS.UUID == uuid.Nil {
				// 	c.Close()
				// 	return
				// }
				// //time.Sleep(3 * time.Second)
				// LS.Username = Client.PlayerName
				// PW := packet.CreatePacketWriter(0x02) //LS.Encode()
				// PW.WriteUUID(uuid.MustParse("f41f9506-e8f7-4a2b-bd4d-4db421620bff"))
				// PW.WriteString(Client.PlayerName)
				// c.AsyncWrite(PW.GetPacket()) //Send Login Success
				// Log.Debug("Sent 0x02_CB - Login Success")
				// time.Sleep(3 * time.Second)
				// Client.UUID = LS.UUID
				// Client.State = PLAY
				// c.SetContext(Client) //Set conn context

				JG := new(packet.JoinGame_CB) //JoinGame - Play 0x26
				PW := JG.Encode(Client.Player.PlayerName, 0)
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
				PW.WriteUUID(Client.Player.UUID)
				PW.WriteString(Client.Player.PlayerName)
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

				//Client.ChunkLoad() //Load chunks

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
			case 0x11:
				Log.Debug("Recieved 0x11 player position")
				X, err := PR.ReadDouble()
				if err != nil {
					Log.Error(err)
				}
				Y, err := PR.ReadDouble()
				if err != nil {
					Log.Error(err)
				}
				Z, err := PR.ReadDouble()
				if err != nil {
					Log.Error(err)
				}
				OnGround, err := PR.ReadByte()
				if err != nil {
					Log.Error(err)
				}
				Log.Debugf("X: %.10F Y: %.10F Z: %.10F ONGROUND: %t", X, Y, Z, OnGround)
			default:
				Log.Debug("Recieved packet play:", PacketID)
			}
		default:
			Log.Info("PP")
		}
	}
}
