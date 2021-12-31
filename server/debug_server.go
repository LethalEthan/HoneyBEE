package server

import (
	"HoneyBEE/config"
	"HoneyBEE/packet"
	"HoneyBEE/player"
	"math/rand"
	"net"
	"runtime"
	"time"

	"github.com/google/uuid"
)

var server net.Listener

type ClientC struct {
	Name            string
	Conn            net.Conn
	ProtocolVersion int32
	State           int
	onlinemode      bool
	encryptstream   *CFB8
	decryptstream   *CFB8
	lerq            *packet.Login_0x01_CB
	Player          player.PlayerObject
	OptionalData    interface{}
}

func DebugServer() {
	var err error
	server, err = net.Listen("tcp", "127.0.0.1:25560")
	if err != nil {
		panic(err)
	}
	Log.Info("Generating Key chain")
	packet.GenerateKeys()
	Log.Info("Creating Entries")
	packet.CreateEntries()
	for {
		conn, err := server.Accept()
		if err != nil {
			panic(err)
		}
		go Start(conn)
	}
}

func Start(conn net.Conn) {
	runtime.LockOSThread()
	Log.Debug("Started Debug server handler")
	conn.SetReadDeadline(time.Now().Add(time.Second * 10))
	conn.SetDeadline(time.Now().Add(time.Second * 10))
	var err error
	var n int
	var PR = packet.CreatePacketReader([]byte{0})
	var PW = packet.CreatePacketWriterWithCapacity(0x00, 2048)
	//var r []byte
	data := make([]byte, 2097152)
	ClientConn := new(ClientC)
	ClientConn.Conn = conn
	ClientConn.State = HANDSHAKE
	for {
	start:
		Log.Debug("Reading from conn...")
		n, err = conn.Read(data)
		Log.Debug("Read from conn")
		if err != nil {
			Log.Debug("Closing because client: ", err)
			conn.Close()
			return
		}
		if n <= 0 {
			Log.Critical("Read 0 bytes from client!")
			conn.Close()
			return
		}
		r := data[0:n]
		if len(r) <= 0 {
			panic("You done fucked up")
		}
		if ClientConn.onlinemode {
			Log.Debug("decrypted stream")
			ClientConn.decryptstream.XORKeyStream(r, r)
		}
		PacketSize, NR, err := packet.DecodeVarInt(r) //NR = Numread, used to note the position in the frame where it read to
		PacketDataSize := PacketSize - 1
		PacketID, NR2, err2 := packet.DecodeVarInt(r[NR:]) //NR2 is the second numread so the Decoder later on will correctly
		if err != nil || err2 != nil {
			panic(err)
		}
		//Check size
		if PacketSize > 2097151 {
			Log.Debug("Packet size greater than 3 byte varint")
			conn.Close()
			return
		}
		PR.SetData(r[NR2+NR:])
		Log.Debug("ClientConn ", ClientConn, "Conn: ", ClientConn.Conn.RemoteAddr().String())
		Log.Debug("PSize: ", PacketSize, "PDS: ", PacketDataSize, " PID: ", PacketID, "State: ", ClientConn.State)
		Log.Debug("Frame: ", r)
		//Legacy Ping - drop conn
		if PacketSize == 0xFE && ClientConn.State == HANDSHAKE {
			Log.Error("Legacy ping")
			conn.Close()
			return
		}
		//Packet Logic
		switch ClientConn.State {
		case HANDSHAKE:
			switch PacketID {
			case 0x00:
				if PacketDataSize == 0 {
					Log.Critical("Packet ordering is whack yo, the bees flew into the glass")
					conn.Close()
					return
				}
				HP := new(packet.Handshake_0x00)
				err := HP.Decode(&PR)
				if err != nil {
					conn.Close()
					return
				}
				ClientConn.ProtocolVersion = HP.ProtocolVersion
				ClientConn.State = int(HP.NextState)
				Log.Debug("NextState: ", HP.NextState)
				goto start
			}
		case STATUS:
			switch PacketID {
			case 0x00:
				Log.Debug("status 0x00_SB")
				if PacketSize == 1 {
					SP := new(packet.Stat_Response)
					SP.ProtocolVersion = ClientConn.ProtocolVersion
					SP.Encode(&PW)
					if err := ClientConn.SendData(conn, PW.GetPacket()); err != nil {
						Log.Error(err)
						conn.Close()
						return
					}
				} else {
					Log.Warning("Packet is bigger than expected")
					conn.Close()
					return
				}
			case 0x01:
				Log.Debug("status 0x01_SB")
				StatP := new(packet.Stat_Ping)
				if err := StatP.Decode(&PR); err != nil {
					Log.Error(err)
					conn.Close()
					return
				}
				StatPClient := new(packet.Stat_Pong)
				StatPClient.Pong = StatP.Ping
				StatPClient.Encode(&PW)
				if err := ClientConn.SendData(conn, PW.GetPacket()); err != nil {
					Log.Error(err)
					conn.Close()
					return
				}
				Log.Debug("WRITER NOTICE ME", PW.GetPacket())
			}
		case LOGIN:
			switch PacketID {
			case 0x00:
				Log.Debug("0x00_SB - Login start")
				LoginStart := new(packet.Login_0x00_SB)
				if err := LoginStart.Decode(&PR); err != nil {
					Log.Error(err)
					conn.Close()
					return
				}
				Log.Info("Name decoded: ", LoginStart.Name)
				ClientConn.Player.PlayerName = LoginStart.Name
				//Encryption Request
				LERQ := new(packet.Login_0x01_CB)
				LERQ.Encode(&PW)
				if err := ClientConn.SendData(conn, PW.GetPacket()); err != nil {
					Log.Error(err)
					conn.Close()
					return
				}
				Log.Debug("Sent 0x01_CB - Encryption Request")
				ClientConn.lerq = LERQ
			case 0x01:
				Log.Debug("Recieved 0x01_SB - Encryption response")
				LERSP := *new(packet.Login_0x01_SB)
				if err := LERSP.Decode(&PR); err != nil {
					Log.Error(err)
					conn.Close()
					return
				}
				if LERSP.SharedSecretLen != 16 {
					Log.Error("Shared secret length is incorrect! recieved: ", LERSP.SharedSecretLen)
					conn.Close()
					return
				}
				if ClientConn.lerq.VerifyTokenLen != LERSP.VerifyTokenLen {
					Log.Error("Client verify token is incorrect, possible tampering recieved: ", LERSP.VerifyTokenLen)
					conn.Close()
					return
				}
				for i := 0; i < len(ClientConn.lerq.VerifyToken); i++ {
					if ClientConn.lerq.VerifyToken[i] != LERSP.VerifyToken[i] {
						Log.Error("Incorrect verify token, closing connection")
						conn.Close()
						return
					}
				}
				Log.Debug("Shared secret: ", LERSP.SharedSecret)
				ClientConn.encryptstream, ClientConn.decryptstream, err = CreateStreamCipher(LERSP.SharedSecret)
				if err != nil {
					Log.Error(err)
					conn.Close()
					return
				}
				// Login Success
				LS := new(packet.Login_0x02_CB)
				LS.UUID, LS.Username, err = packet.Auth(ClientConn.Player.PlayerName, LERSP.SharedSecret)
				if err != nil {
					Log.Debug(err)
					conn.Close()
					return
				}
				if LS.UUID == uuid.Nil {
					Log.Error("UUID is nil!")
					conn.Close()
					return
				}
				Log.Debug("UUID: ", LS.UUID)
				ClientConn.Player.UUID = LS.UUID
				ClientConn.Player.PlayerName = LS.Username // We reset the playername to what the sessionserver says
				ClientConn.onlinemode = true
				err := LS.Encode(&PW)
				if err != nil { // Encode Login Success
					Log.Error("Could not encode login success: ", err)
					conn.Close()
					return
				}
				// Send Login Success
				if err := ClientConn.SendData(conn, PW.GetPacket()); err != nil {
					Log.Error(err)
					conn.Close()
					return
				}
				Log.Debug("Sent 0x02_CB - Login Success")
				ClientConn.State = PLAY
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
				JG.Encode(&PW)
				if err := ClientConn.SendData(conn, PW.GetPacket()); err != nil { //JoinGamePacket
					Log.Error(err)
					conn.Close()
					return
				}

				Log.Debug("Sent Join Game - 0x26 - P")

				PW.ResetData(0x18) // Server brand
				PW.WriteString("minecraft:brand")
				PW.WriteString("HoneyBEE")
				Log.Debug("Sent server brand")
				if err := ClientConn.SendData(conn, PW.GetPacket()); err != nil {
					Log.Error(err)
					conn.Close()
					return
				}

				// PW.ResetData(0x0E) //Difficulty
				// PW.WriteByte(0)
				// PW.WriteBoolean(true)
				// Log.Debug("Sent Difficulty")
				// ClientConn.SendData(conn,PW.GetPacket())

				// PW.ResetData(0x32) //Player Abilities
				// PW.WriteByte(0x0F)
				// PW.WriteFloat(0.05)
				// PW.WriteFloat(0.1)
				// Log.Debug("Sent Player Abilities")
				// ClientConn.SendData(conn,PW.GetPacket())

				PW.ResetData(0x48) //Held Item Change
				PW.WriteByte(0)
				Log.Debug("Sent held item change")
				if err := ClientConn.SendData(conn, PW.GetPacket()); err != nil {
					Log.Error(err)
					conn.Close()
					return
				}

				PW.ResetData(0x65) //Declare Recipes - investigate
				PW.WriteVarInt(0)
				Log.Debug("Sent recipe")
				if err := ClientConn.SendData(conn, PW.GetPacket()); err != nil {
					Log.Error(err)
					conn.Close()
					return
				}

				// Log.Debug("Sent Tags")
				// if err := ClientConn.SendData(conn, TagsPacket); err != nil {
				// 	Log.Error(err)
				// 	conn.Close()
				// 	return
				// }

				PW.ResetData(0x1B) // Entity Status - Disable reduced debug mode
				PW.WriteInt(2)
				PW.WriteByte(DRDB)
				Log.Debug("Sent Entity Status")
				if err := ClientConn.SendData(conn, PW.GetPacket()); err != nil {
					Log.Error(err)
					conn.Close()
					return
				}

				PW.ResetData(0x1B) // Entity Status - set op level 4
				PW.WriteInt(2)
				PW.WriteByte(OP4)
				Log.Debug("Sent Entity Status")
				if err := ClientConn.SendData(conn, PW.GetPacket()); err != nil {
					Log.Error(err)
					conn.Close()
					return
				}

				// PW.ResetData(0x12) //Declare commands
				// PW.WriteVarInt(0)
				// PW.WriteVarInt(0)
				// Log.Debug("Sent Declare commands")
				// ClientConn.SendData(conn,PW.GetPacket())

				// Log.Debug("Sent Unlock Recipes")
				// if err := ClientConn.SendData(conn, UnlockRecipesPacket); err != nil { // Unlock Recipes
				// 	Log.Error(err)
				// 	conn.Close()
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
				if err := ClientConn.SendData(conn, PW.GetPacket()); err != nil {
					Log.Error(err)
					conn.Close()
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
				// ClientConn.SendData(conn,PW.GetPacket())

				PW.ResetData(0x36) //PlayerInfo
				PW.WriteVarInt(0)
				PW.WriteVarInt(1)
				PW.WriteUUID(ClientConn.Player.UUID)
				PW.WriteString(ClientConn.Player.PlayerName)
				PW.WriteVarInt(0)      // num properties
				PW.WriteVarInt(1)      // gamemode
				PW.WriteVarInt(0)      // ping
				PW.WriteBoolean(false) // has display name
				Log.Debug("Sent Player info")
				if err := ClientConn.SendData(conn, PW.GetPacket()); err != nil {
					Log.Error(err)
					conn.Close()
					return
				}

				PW.ResetData(0x49) // Update view position
				PW.WriteVarInt(0)
				PW.WriteVarInt(0)
				Log.Debug("Sent view position")
				if err := ClientConn.SendData(conn, PW.GetPacket()); err != nil {
					Log.Error(err)
					conn.Close()
					return
				}

				PW.ResetData(0x4A) // Update view distance
				PW.WriteVarInt(12)
				Log.Debug("Sent view distance")
				if err := ClientConn.SendData(conn, PW.GetPacket()); err != nil {
					Log.Error(err)
					conn.Close()
					return
				}

				// PW.ResetData(0x4D) // Entity metadata
				// PW.WriteVarInt(2)
				// Log.Debug("Sent Entity Metadata")
				// ClientConn.SendData(conn,PW.GetPacket())

				if err := ClientConn.ChunkLoad(conn); err != nil { // Load chunks
					Log.Error(err)
					conn.Close()
					return
				}

				PW.ResetData(0x4B) // Spawn Position
				PW.WritePosition(0, 64, 0)
				PW.WriteFloat(0.0)
				Log.Debug("Sent Spawn position")
				if err := ClientConn.SendData(conn, PW.GetPacket()); err != nil {
					Log.Error(err)
					conn.Close()
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
				if err := ClientConn.SendData(conn, PW.GetPacket()); err != nil {
					Log.Error(err)
					conn.Close()
					return
				}

				PW.ResetData(0x58) // Time Update
				PW.WriteLong(0)
				PW.WriteLong(-12000)
				Log.Debug("Sent Time update")
				if err := ClientConn.SendData(conn, PW.GetPacket()); err != nil {
					Log.Error(err)
					conn.Close()
					return
				}

				PW.ResetData(0x14) // Window Items
				PW.WriteUByte(0)
				PW.WriteVarInt(1)
				PW.WriteVarInt(46)
				PW.WriteArray(make([]byte, 46))
				PW.WriteByte(0)
				Log.Debug("Sent Window Items")
				if err := ClientConn.SendData(conn, PW.GetPacket()); err != nil {
					Log.Error(err)
					conn.Close()
					return
				}

				PW.ResetData(0x52) // Update Health
				PW.WriteFloat(20.0)
				PW.WriteVarInt(20)
				PW.WriteFloat(5.0)
				Log.Debug("Sent Update Health")
				if err := ClientConn.SendData(conn, PW.GetPacket()); err != nil {
					Log.Error(err)
					conn.Close()
					return
				}

				PW.ResetData(0x51) // Set expereince
				PW.WriteFloat(1.0)
				PW.WriteVarInt(1)
				PW.WriteVarInt(7)
				Log.Debug("Sent Set Experience")
				if err := ClientConn.SendData(conn, PW.GetPacket()); err != nil {
					Log.Error(err)
					conn.Close()
					return
				}
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
		Log.Debug("Reached end of switch")
		conn.SetReadDeadline(time.Now().Add(time.Second * 10))
		conn.SetDeadline(time.Now().Add(time.Second * 10))
	}
}

func (C *ClientC) SendData(c net.Conn, data []byte) error {
	if C.onlinemode {
		Log.Debug("Sending online packet")
		C.encryptstream.XORKeyStream(data, data)
		_, err := c.Write(data)
		return err
	} else {
		Log.Debug("Sending offline")
		_, err := c.Write(data)
		return err
	}
}

func (Client *ClientC) ChunkLoad(c net.Conn) error {
	//Send chunk
	if err := Client.SendData(c, CreateLightData(0, 0)); err != nil {
		c.Close() //send
		return err
	}
	if err := Client.SendData(c, CreateChunk(0, 0)); err != nil { //CreateChunk(0, 0)
		c.Close() //send
		return err
	}
	var X int32 //j
	var Z int32 //i
	for Z = 0; Z < 12; Z++ {
		for X = 0; X < 12; X++ {
			if X != 0 || Z != 0 {
				if err := Client.SendData(c, CreateLightData(X, -Z)); err != nil {
					c.Close()
					return err
				}
				if err := Client.SendData(c, CreateChunk(X, Z)); err != nil { // Send chunk
					c.Close()
					return err
				}
				if Z == 0 {
					if err := Client.SendData(c, CreateLightData(-X, Z)); err != nil {
						c.Close()
						return err
					}
					if err := Client.SendData(c, CreateChunk(-X, Z)); err != nil {
						c.Close()
						return err
					}
				}
				if X == 0 {
					if err := Client.SendData(c, CreateLightData(X, -Z)); err != nil {
						c.Close()
						return err
					}
					if err := Client.SendData(c, CreateChunk(X, -Z)); err != nil {
						c.Close()
						return err
					}
				}
				if X != 0 && Z != 0 {
					if err := Client.SendData(c, CreateLightData(X, -Z)); err != nil {
						c.Close()
						return err
					}
					if err := Client.SendData(c, CreateChunk(X, -Z)); err != nil {
						c.Close()
						return err
					}
					if err := Client.SendData(c, CreateLightData(-X, Z)); err != nil {
						c.Close()
						return err
					}
					if err := Client.SendData(c, CreateChunk(-X, Z)); err != nil {
						c.Close()
						return err
					}
					if err := Client.SendData(c, CreateLightData(-X, -Z)); err != nil {
						c.Close()
						return err
					}
					if err := Client.SendData(c, CreateChunk(-X, -Z)); err != nil {
						c.Close()
						return err
					}
				}
			}
		}
	}
	Log.Debug("Sent chunks!")
	PW := packet.CreatePacketWriterWithCapacity(0x20, 128) //Initialise World Border
	PW.WriteDouble(0.0)
	PW.WriteDouble(0.0)
	PW.WriteDouble(348.0)
	PW.WriteDouble(348.0)
	PW.WriteVarLong(0)
	PW.WriteVarInt(29999984)
	PW.WriteVarInt(0)
	PW.WriteVarInt(0)
	Log.Debug("Sent Init World Border")
	if err := Client.SendData(c, PW.GetPacket()); err != nil {
		c.Close()
		return err
	}
	return nil
}
