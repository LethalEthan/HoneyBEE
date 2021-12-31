package server

import (
	"HoneyBEE/config"
	"HoneyBEE/packet"
	"HoneyBEE/player"
	"HoneyBEE/utils"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/panjf2000/gnet"
)

type Client struct {
	RemoteAddr string
	Player     player.PlayerObject
	PW         packet.PacketWriter
	PR         packet.PacketReader
	// Conn            gnet.Conn
	ProtocolVersion int32
	State           int
	Read            chan []byte
	lerq            packet.Login_0x01_CB
	decryptstream   *CFB8
	encryptstream   *CFB8
	onlinemode      bool
	OptionalData    interface{}
}

type Read struct{}

func (S *Server) React(Frame []byte, Conn gnet.Conn) (Out []byte, Action gnet.Action) {
	CO, load := S.ConnectedSockets.Load(Conn.RemoteAddr().String())
	if !load {
		Log.Debug("Could not load object")
		Conn.Close()
		return
	}
	Conn.ResetBuffer()
	CC := CO.(Client)
	// Log.Debug("React hit!")
	if len(Frame) > 0 {
		CC.Read <- Frame
		return nil, gnet.None
	} else {
		Log.Warning("Frame is 0?!")
		CC.Read <- nil
		return nil, gnet.Close
	}
}

func (Client *Client) ClientReact(c gnet.Conn) {
	var PR = &Client.PR
	var PW = &Client.PW
	for {
		Frame := <-Client.Read //Block until react fires
		if Frame == nil || len(Frame) == 0 {
			Log.Debug("Frame nil, closing")
			Client = nil
			c.Close()
			return
		}
		if Client.onlinemode {
			PR.SetData(Client.DecryptPacket(Frame))
		} else {
			PR.SetData(Frame)
		}
		//Get PacketSize and Data
		PacketSize, NR, err := PR.ReadVarInt() //NR = Numread, used to note the position in the frame where it read to
		if err != nil {
			Log.Error(err)
			c.Close()
			return
		}
		PacketDataSize := PacketSize - 1
		PacketID, _, err := PR.ReadVarInt()
		if err != nil {
			Log.Error(err)
			c.Close()
			return
		}
		if int(PacketSize+int32(NR)) != len(Frame) {
			Log.Debug("Frame: ", Frame)
			Log.Debug("PacketSize", PacketSize)
			Log.Debug("Frame length", len(Frame))
			Log.Debug("Grouped packets")
		}
		//Packets cannot be bigger than a 3 byte varint :(
		if PacketSize > 2097151 {
			Log.Error("packet size greater than 3 byte varint")
			c.Close() //Disconnect the client until I find a solution, my idea is a custom mod or client that raises the packet size to a VarLong because why not
			return
		}
		// Legacy Ping - drop conn
		if PacketSize == 0xFE && Client.State == HANDSHAKE {
			Log.Error("Legacy ping recieved, closing...")
			c.Close()
			return
		}
		// Packet Logic
		switch Client.State {
		case HANDSHAKE:
			switch PacketID {
			case 0x00:
				if PacketDataSize == 0 {
					Log.Error("Packet ordering is whack yo, the bees flew into the glass")
					c.Close()
					return
				}
				HP := new(packet.Handshake_0x00)
				if err := HP.Decode(PR); err != nil {
					Log.Error("Error while decoding HP: ", err)
					c.Close()
					return
				}
				Client.ProtocolVersion = HP.ProtocolVersion
				if Client.ProtocolVersion != utils.PrimaryMinecraftProtocolVersion {
					Log.Debug("Closing becase client does not match server protocol version!")
					c.Close()
				}
				Client.State = int(HP.NextState)
			}
		case STATUS:
			switch PacketID {
			case 0x00:
				Log.Debug("status 0x00_SB")
				if PacketSize == 1 {
					SP := new(packet.Stat_Response)
					SP.ProtocolVersion = Client.ProtocolVersion
					SP.Encode(PW)
					if err := Client.SendData(c, PW.GetPacket()); err != nil {
						Log.Error(err)
						c.Close()
						return
					}
				} else {
					Log.Error("Status pong packet is bigger than expected")
					c.Close()
					return
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
				if err := Client.SendData(c, PW.GetPacket()); err != nil {
					Log.Error(err)
					c.Close()
					return
				}
			}
		case LOGIN:
			switch PacketID {
			case 0x00:
				// Login Start
				Log.Debug("0x00_SB - Login start")
				LoginStart := new(packet.Login_0x00_SB)
				if err := LoginStart.Decode(PR); err != nil {
					Log.Error(err)
					c.Close()
					return
				}
				Log.Info("Name decoded: ", LoginStart.Name)
				Client.Player.PlayerName = LoginStart.Name
				//Login Encryption Request
				LERQ := *new(packet.Login_0x01_CB)
				LERQ.Encode(PW)
				if err := Client.SendData(c, PW.GetPacket()); err != nil {
					Log.Error(err)
					c.Close()
					return
				}
				Client.lerq = LERQ
				Log.Debug("Sent 0x01_CB - Encryption Request")
			case 0x01:
				Log.Debug("Recieved 0x01_SB - Encryption response")
				LERSP := *new(packet.Login_0x01_SB)
				if err := LERSP.Decode(PR); err != nil {
					Log.Error(err)
					c.Close()
					return
				}
				if LERSP.SharedSecretLen != 16 {
					Log.Error("Shared secret length is incorrect! recieved: ", LERSP.SharedSecretLen)
					c.Close()
					return
				}
				if Client.lerq.VerifyTokenLen != LERSP.VerifyTokenLen {
					Log.Error("Client verify token is incorrect, possible tampering recieved: ", LERSP.VerifyTokenLen)
					c.Close()
					return
				}
				for i := 0; i < len(Client.lerq.VerifyToken); i++ {
					if Client.lerq.VerifyToken[i] != LERSP.VerifyToken[i] {
						Log.Error("Incorrect verify token, closing connection")
						c.Close()
						return
					}
				}
				Log.Debug("Shared secret: ", LERSP.SharedSecret)
				Client.encryptstream, Client.decryptstream, err = CreateStreamCipher(LERSP.SharedSecret)
				if err != nil {
					Log.Error(err)
					c.Close()
					return
				}
				// Login Success
				LS := new(packet.Login_0x02_CB)
				LS.UUID, LS.Username, err = packet.Auth(Client.Player.PlayerName, LERSP.SharedSecret)
				if err != nil {
					Log.Debug(err)
					c.Close()
					return
				}
				if LS.UUID == uuid.Nil {
					Log.Error("UUID is nil!")
					c.Close()
					return
				}
				Log.Debug("UUID: ", LS.UUID)
				Client.Player.UUID = LS.UUID
				Client.Player.PlayerName = LS.Username // We reset the playername to what the sessionserver says
				Client.onlinemode = true
				err := LS.Encode(PW)
				if err != nil { // Encode Login Success
					Log.Error("Could not encode login success: ", err)
					c.Close()
					return
				}
				// Send Login Success
				if err := Client.SendData(c, PW.GetPacket()); err != nil {
					Log.Error(err)
					c.Close()
					return
				}
				Log.Debug("Sent 0x02_CB - Login Success")
				Client.State = PLAY
				// c.SetContext(Client) // Set conn context
				// JoinGame - Play 0x26
				JG := new(packet.JoinGame_CB)
				rand.Seed(time.Now().UnixNano())
				JG.EntityID = rand.Int31n(5000-4) + 4
				JG.IsHardcore = false
				JG.Gamemode = 1
				JG.PreviousGamemode = -1
				JG.WorldCount = 4
				JG.WorldNames = []packet.Identifier{"minecraft:overworld", "minecraft:overworld_caves", "minecraft:the_nether", "minecraft:the_end"}
				JG.WorldName = "minecraft:overworld"
				JG.HashedSeed = 0
				JG.MaxPlayers = 0
				JG.ViewDistance = int32(config.GConfig.Performance.ViewDistance)
				JG.SimulationDistance = int32(config.GConfig.Performance.SimulationDistance)
				JG.ReducedDebugInfo = false
				JG.EnableRespawnScreen = true
				JG.IsDebug = false
				JG.IsFlat = true
				JG.Encode(PW)
				if err := Client.SendData(c, PW.GetPacket()); err != nil { //JoinGamePacket
					Log.Error(err)
					c.Close()
					return
				}

				Log.Debug("Sent Join Game - 0x26 - P")

				PW.ResetData(0x18) // Server brand
				PW.WriteString("minecraft:brand")
				PW.WriteString("HoneyBEE")
				Log.Debug("Sent server brand")
				if err := Client.SendData(c, PW.GetPacket()); err != nil {
					Log.Error(err)
					c.Close()
					return
				}

				// PW.ResetData(0x0E) //Difficulty
				// PW.WriteByte(0)
				// PW.WriteBoolean(true)
				// Log.Debug("Sent Difficulty")
				// Client.SendData(c,PW.GetPacket())

				// PW.ResetData(0x32) //Player Abilities
				// PW.WriteByte(0x0F)
				// PW.WriteFloat(0.05)
				// PW.WriteFloat(0.1)
				// Log.Debug("Sent Player Abilities")
				// Client.SendData(c,PW.GetPacket())

				PW.ResetData(0x48) //Held Item Change
				PW.WriteByte(0)
				Log.Debug("Sent held item change")
				if err := Client.SendData(c, PW.GetPacket()); err != nil {
					Log.Error(err)
					c.Close()
					return
				}

				PW.ResetData(0x65) //Declare Recipes - investigate
				PW.WriteVarInt(0)
				Log.Debug("Sent recipe")
				if err := Client.SendData(c, PW.GetPacket()); err != nil {
					Log.Error(err)
					c.Close()
					return
				}

				// Log.Debug("Sent Tags")
				// if err := Client.SendData(c, TagsPacket); err != nil {
				// 	Log.Error(err)
				// 	c.Close()
				// 	return
				// }

				PW.ResetData(0x1B) // Entity Status - Disable reduced debug mode
				PW.WriteInt(2)
				PW.WriteByte(DRDB)
				Log.Debug("Sent Entity Status")
				if err := Client.SendData(c, PW.GetPacket()); err != nil {
					Log.Error(err)
					c.Close()
					return
				}

				PW.ResetData(0x1B) // Entity Status - set op level 4
				PW.WriteInt(2)
				PW.WriteByte(OP4)
				Log.Debug("Sent Entity Status")
				if err := Client.SendData(c, PW.GetPacket()); err != nil {
					Log.Error(err)
					c.Close()
					return
				}

				// PW.ResetData(0x12) //Declare commands
				// PW.WriteVarInt(0)
				// PW.WriteVarInt(0)
				// Log.Debug("Sent Declare commands")
				// Client.SendData(c,PW.GetPacket())

				// Log.Debug("Sent Unlock Recipes")
				// if err := Client.SendData(c, UnlockRecipesPacket); err != nil { // Unlock Recipes
				// 	Log.Error(err)
				// 	c.Close()
				// 	return
				// }

				PW.ResetData(0x38) // Player pos and look
				PW.WriteDouble(0.0)
				PW.WriteDouble(64.0)
				PW.WriteDouble(0.0)
				PW.WriteFloat(0.0)
				PW.WriteFloat(0.0)
				PW.WriteUByte(0)
				PW.WriteVarInt(0)
				PW.WriteBoolean(true)
				Log.Debug("Sent Player pos and look")
				if err := Client.SendData(c, PW.GetPacket()); err != nil {
					Log.Error(err)
					c.Close()
					return
				}

				// PW.ResetData(0x0F) //Chat Message
				// CM := new(packet.ChatMessage_CB)
				// CM.Chat = new(jsonstruct.ChatComponent)
				// CM.Chat.Text = "Hello"
				// CMB := CM.Chat.MarshalChatComponent()
				// PW.WriteArray(CMB)
				// PW.WriteByte(1)
				// PW.WriteUUID(LS.UUID)
				// Log.Debug("Sent Chat message")
				// Client.SendData(c,PW.GetPacket())

				PW.ResetData(0x36) //PlayerInfo
				PW.WriteVarInt(0)
				PW.WriteVarInt(1)
				PW.WriteUUID(Client.Player.UUID)
				PW.WriteString(Client.Player.PlayerName)
				PW.WriteVarInt(0)      // num properties
				PW.WriteVarInt(1)      // gamemode
				PW.WriteVarInt(0)      // ping
				PW.WriteBoolean(false) // has display name
				Log.Debug("Sent Player info")
				if err := Client.SendData(c, PW.GetPacket()); err != nil {
					Log.Error(err)
					c.Close()
					return
				}

				PW.ResetData(0x49) // Update view position
				PW.WriteVarInt(0)
				PW.WriteVarInt(0)
				Log.Debug("Sent view position")
				if err := Client.SendData(c, PW.GetPacket()); err != nil {
					Log.Error(err)
					c.Close()
					return
				}

				PW.ResetData(0x4A) // Update view distance
				PW.WriteVarInt(12)
				Log.Debug("Sent view distance")
				if err := Client.SendData(c, PW.GetPacket()); err != nil {
					Log.Error(err)
					c.Close()
					return
				}

				// PW.ResetData(0x4D) // Entity metadata
				// PW.WriteVarInt(2)
				// Log.Debug("Sent Entity Metadata")
				// Client.SendData(c,PW.GetPacket())

				if err := Client.ChunkLoad(c); err != nil { // Load chunks
					Log.Error(err)
					c.Close()
					return
				}

				PW.ResetData(0x4B) // Spawn Position
				PW.WritePosition(0, 64, 0)
				PW.WriteFloat(0.0)
				Log.Debug("Sent Spawn position")
				if err := Client.SendData(c, PW.GetPacket()); err != nil {
					Log.Error(err)
					c.Close()
					return
				}

				PW.ResetData(0x38) // Player pos and look
				PW.WriteDouble(0.0)
				PW.WriteDouble(64.0)
				PW.WriteDouble(0.0)
				PW.WriteFloat(0.0)
				PW.WriteFloat(0.0)
				PW.WriteByte(0)
				PW.WriteVarInt(0)
				PW.WriteBoolean(true)
				Log.Debug("Sent Player pos and look")
				if err := Client.SendData(c, PW.GetPacket()); err != nil {
					Log.Error(err)
					c.Close()
					return
				}

				PW.ResetData(0x58) // Time Update
				PW.WriteLong(0)
				PW.WriteLong(-12000)
				Log.Debug("Sent Time update")
				if err := Client.SendData(c, PW.GetPacket()); err != nil {
					Log.Error(err)
					c.Close()
					return
				}

				PW.ResetData(0x14) // Window Items
				PW.WriteUByte(0)
				PW.WriteVarInt(1)
				PW.WriteVarInt(46)
				PW.WriteArray(make([]byte, 46))
				PW.WriteByte(0)
				Log.Debug("Sent Window Items")
				if err := Client.SendData(c, PW.GetPacket()); err != nil {
					Log.Error(err)
					c.Close()
					return
				}

				PW.ResetData(0x52) // Update Health
				PW.WriteFloat(20.0)
				PW.WriteVarInt(20)
				PW.WriteFloat(5.0)
				Log.Debug("Sent Update Health")
				if err := Client.SendData(c, PW.GetPacket()); err != nil {
					Log.Error(err)
					c.Close()
					return
				}

				PW.ResetData(0x51) // Set expereince
				PW.WriteFloat(1.0)
				PW.WriteVarInt(1)
				PW.WriteVarInt(7)
				Log.Debug("Sent Set Experience")
				if err := Client.SendData(c, PW.GetPacket()); err != nil {
					Log.Error(err)
					c.Close()
					return
				}
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
				OnGround, err := PR.ReadBoolean()
				if err != nil {
					Log.Error(err)
				}
				Log.Debugf("X: %.10F Y: %.10F Z: %.10F ONGROUND: %t", X, Y, Z, OnGround)
				PW.ResetData(0x21)
				PW.WriteLong(138412846)
				if err := Client.SendData(c, PW.GetPacket()); err != nil {
					Log.Error(err)
					c.Close()
					return
				}
			case 0x12:
				Log.Debug("Recieved 0x11 player position and rotation")
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
				Yaw, err := PR.ReadFloat()
				if err != nil {
					Log.Error(err)
				}
				Pitch, err := PR.ReadFloat()
				if err != nil {
					Log.Error(err)
				}
				OnGround, err := PR.ReadBoolean()
				if err != nil {
					Log.Error(err)
				}
				Log.Debugf("X: %.10F Y: %.10F Z: %.10F Yaw: %.10F Pitch: %.10F ONGROUND: %t", X, Y, Z, Yaw, Pitch, OnGround)
			case 0x13:
				Yaw, err := PR.ReadFloat()
				if err != nil {
					Log.Error(err)
				}
				Pitch, err := PR.ReadFloat()
				if err != nil {
					Log.Error(err)
				}
				OnGround, err := PR.ReadBoolean()
				if err != nil {
					Log.Error(err)
				}
				Log.Debugf("Yaw: %.10F Pitch: %.10F ONGROUND: %t", Yaw, Pitch, OnGround)
			case 0x14:
				OnGround, err := PR.ReadBoolean()
				if err != nil {
					Log.Error(err)
				}
				Log.Debugf("ONGROUND: %t", OnGround)
			default:
				Log.Debug("Recieved packet play:", PacketID)
			}
		default:
			Log.Info("PP")
		}

	}
}
