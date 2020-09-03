package server

import (
	"Packet"
	"VarTool"
	config "config"
	"crypto/rsa"
	"encoding/json"
	"event"
	"fmt"
	"io/ioutil"
	"net"
	"player"

	logging "github.com/op/go-logging"
)

//Define used variables
var (
	//Encryption Stuff
	publicKey      *rsa.PublicKey  //Public Key - Used for authentication
	publicKeyBytes []byte          //PublicKey in a byte array for packet delivery and Auth check
	privateKey     *rsa.PrivateKey //PrivateKey
	KeyLength      int             //Keylength used by Encryption Request
	//	ClientSharedSecret []byte          //CSS
	//	ClientVerifyToken  []byte          //CVT
	playername    string          //For Authentication
	Log           *logging.Logger //Pretty obvious
	CurrentStatus *ServerStatus   //ServerStatus Object
	//Encryption stuff again
	//DEBUG                 = true            //Output Debug info?
	GotDaKeys = false //Got dem keys?
	//	ClientSharedSecretLen = 128                       //Initialise CSSL
	//	ClientVerifyTokenLen  = 128                       //Initialise CVTL
	serverID                 = ""                        //this isn't used by mc anymore
	ServerVerifyToken        = make([]byte, 4)           //Initialise a 4 element byte slice of cake
	PlayerMap                = make(map[string]string)   //Map Player to UUID
	PlayerConnMap            = make(map[net.Conn]string) //Map Connection to Player
	ConnPlayerMap            = make(map[uint32]net.Conn) //Map EID to Connection
	GEID              uint32 = 2
	Config            *config.Config
	DEBUG             bool
	KC                = false
)

const (
	MinecraftVersion               = "1.15.2" //Supported MC version
	MinecraftProtocolVersion int32 = 578      //Supported MC protocol Version
	ServerVerifyTokenLen           = 4        //Should always be 4 on notchian servers
)

func GetKeyChain() {
	privateKey = Packet.GetPrivateKey()
	publicKeyBytes = Packet.GetPublicKeyBytes()
	publicKey = Packet.GetPublicKey()
	GotDaKeys = true
}

type PacketHeader struct {
	packet     []byte
	packetSize int32
	packetID   int32
	protocol   int32
}

type Version interface {
	MCDEFAULT(Conn *ClientConnection)
	MC1_15_2(Conn *ClientConnection)
	MC1_16(Conn *ClientConnection)
	MC1_16_1(Conn *ClientConnection)
	MC1_16_2(Conn *ClientConnection)
}

func Init() {
	if !GotDaKeys {
		GetKeyChain()
	}
	Config = config.GetConfig()
	DEBUG = Config.Server.DEBUG
	CurrentStatus = CreateStatusObject(578, "1.15.2")
	player.Init()
	go player.GCPlayer()
	go ProtocolToVersionInit()
}

func (PH *PacketHeader) MCDEFAULT(Conn *ClientConnection) {
	Log.Debug("Unsupported Version Handling")
	HandleUnsupported(Conn, *PH)
	return
}

func (PH *PacketHeader) MC1_15_2(Conn *ClientConnection) {
	Log.Debug("1.15.2")
	Handle_MC1_15_2(Conn, *PH)
	return
}

func (PH *PacketHeader) MC1_16(Conn *ClientConnection) {

}

func (PH *PacketHeader) MC1_16_1(Conn *ClientConnection) {
	Log.Debug("1.16.1")
	Handle_MC1_16_1(Conn, *PH)
	return
}

func (*PacketHeader) MC1_16_2(Conn *ClientConnection) {
	Log.Debug("1.16.2")
}

func HandleConnection(Connection *ClientConnection) {
	if !KC {
		GetKeyChain()
		KC = true
	}
	Log.Info("Connection handler initiated")
	//Løøps
	PH := new(PacketHeader)
	for !Connection.isClosed {
		//packet, packetSize, packetID, err := readPacketHeader(Connection)
		var err error
		PH.packet, PH.packetSize, PH.packetID, err = readPacketHeader(Connection)
		if err != nil {
			CloseClientConnection(Connection)
			Log.Error("Connection Terminated: " + err.Error())
			return
		}
		//DEBUG: output debug info
		DisplayPacketInfo(*PH, Connection)
		//Create Packet Reader
		reader := Packet.CreatePacketReader(PH.packet)
		//Packet Handling
		switch Connection.State {
		case HANDSHAKE: //Handle Handshake
			switch PH.packetID {
			case 0x00:
				{
					//--Packet 0x00 S->C Start--//
					Hpacket, err := Packet.HandshakePacketCreate(PH.packetSize, reader)
					if err != nil {
						CloseClientConnection(Connection) //You have been terminated
						Log.Error(err.Error())
					}
					Connection.KeepAlive()
					//Proto, _ := reader.ReadVarInt()
					//Log.Debug("Protocol:", Proto)
					Connection.State = int(Hpacket.NextState)
					PH.protocol = Hpacket.ProtocolVersion
					var pv Version
					pv = PH
					switch Hpacket.ProtocolVersion {
					case 578:
						pv.MC1_15_2(Connection)
						return
					case 735:
						pv.MC1_16(Connection)
					case 736:
						pv.MC1_16_1(Connection)
						return
					case 751:
						pv.MC1_16_2(Connection)
						return
					default:
						Log.Warning("Unsupported protocol:", Hpacket.ProtocolVersion, "("+ProtocolToVer[Hpacket.ProtocolVersion]+")", "- sending status and closing connection!")
						pv.MCDEFAULT(Connection)
						CloseClientConnection(Connection)
						return
					}
					Log.Debug("Bruhmo: ", pv)
					Log.Debug("Bruh")
					return
					//--Packet 0x00 End--//
				}
			case 0xFE:
				{
					//--Packet 0xFE Legacy Ping Request --//
					Log.Warning("Legacy Ping Request received! - Terminated")
					CloseClientConnection(Connection)
					return
					//--Packet 0xFE End--//
				}
			}

		case STATUS: //Handle Status Request
			{
				switch PH.packetID {
				case 0x00:
					{
						//--Packet 0x00 S->C Start--//
						Connection.KeepAlive()
						writer := Packet.CreatePacketWriter(0x00)
						marshaledStatus, err := json.Marshal(*CurrentStatus) //Sends status via json
						if err != nil {
							Log.Error(err.Error())
							CloseClientConnection(Connection)
							return
						}
						writer.WriteString(string(marshaledStatus))
						SendData(Connection, writer)
						break
					}
				case 0x01:
					{
						//--Packet 0x01 S->C Start--//
						Connection.KeepAlive()
						writer := Packet.CreatePacketWriter(0x01)
						Log.Debug("Status State, packetID 0x01")
						mirror, _ := reader.ReadLong()
						Log.Debug("Mirror:", mirror)
						writer.WriteLong(mirror)
						SendData(Connection, writer)
						CloseClientConnection(Connection)
						break
						//--Packet 0x01 End--//
					}
				}
			}
		}
	}
}

//SendData - Sends the data to the client
func SendData(Connection *ClientConnection, writer *Packet.PacketWriter) {
	Connection.Conn.Write(writer.GetPacket())
}

//getPacketData - Reads incoming packets and returns data in a byte array to which ever function requires it
func getPacketData(Conn net.Conn) ([]byte, error) {
	return ioutil.ReadAll(Conn)
}

func HandleUnsupported(Connection *ClientConnection, PH PacketHeader) {
	Log.Info("Connection handler for Unsupported MC initiated")
	CurrentStatus = CreateStatusObject(PH.protocol, "Unsupported")
	if publicKey == nil || privateKey == nil {
		panic("Keys have been thanos snapped")
	}
	for !Connection.isClosed {
		var err error
		PH.packet, PH.packetSize, PH.packetID, err = readPacketHeader(Connection)
		if err != nil {
			CloseClientConnection(Connection)
			Log.Error("Connection Terminated: " + err.Error())
			return
		}
		DisplayPacketInfo(PH, Connection)
		//Create Packet Reader
		reader := Packet.CreatePacketReader(PH.packet)
		//Packet Handling
		switch Connection.State {
		case STATUS: //Handle Status Request
			{
				switch PH.packetID {
				case 0x00:
					{
						//--Packet 0x00 S->C Start--// Request Status
						Connection.KeepAlive()
						writer := Packet.CreatePacketWriter(0x00)
						marshaledStatus, err := json.Marshal(*CurrentStatus) //Sends status via json
						if err != nil {
							Log.Error(err.Error())
							CloseClientConnection(Connection)
							return
						}
						writer.WriteString(string(marshaledStatus))
						SendData(Connection, writer)
					}
				case 0x01:
					{
						//--Packet 0x01 S->C Start--// Ping
						Connection.KeepAlive()
						writer := Packet.CreatePacketWriter(0x01)
						Log.Debug("Status State, packetID 0x01")
						mirror, _ := reader.ReadLong()
						Log.Debug("Mirror:", mirror)
						writer.WriteLong(mirror)
						SendData(Connection, writer)
						CloseClientConnection(Connection)
						//--Packet 0x01 End--//
					}
				}
			}
		case LOGIN:
			switch PH.packetID {
			default:
				{
					CloseClientConnection(Connection)
				}
			}
		case PLAY:
			switch PH.packetID {
			default:
				{
					CloseClientConnection(Connection)
				}
			}
		}
	}
}

//Server ClientConn -> Player ClientConn - Will be removed later
// func TranslatePlayerStruct(Conn *ClientConnection) *player.ClientConnection {
// 	PC := new(player.ClientConnection)
// 	PC.Conn = Conn.Conn
// 	PC.State = Conn.State
// 	PC.Closed = Conn.isClosed
// 	return PC
// }

//Server ClientConn -> Packet ClientConn - Will be removed later
// func TranslatePacketStruct(Conn *ClientConnection) *Packet.ClientConnection {
// 	PE := new(Packet.ClientConnection)
// 	PE.Conn = Conn.Conn
// 	PE.State = Conn.State
// 	PE.Closed = Conn.isClosed
// 	return PE
// }

//readPacketHeader - Reads the packet Header for Packet ID and size info
func readPacketHeader(Conn *ClientConnection) ([]byte, int32, int32, error) {
	//Information used from wiki.vg
	//Read Packet size
	packetSize, err := VarTool.ParseVarIntFromConnection(Conn.Conn)
	if err != nil {
		return nil, 0, 0, err //Return nothing on error
	}
	//Handling Legacy Handshake
	if packetSize == 254 && Conn.State == HANDSHAKE {
		return nil, 254, 0xFE, nil
	}
	packetID, err := VarTool.ParseVarIntFromConnection(Conn.Conn)
	if err != nil {
		return nil, 0, 0, err //Return nothing on error
	}
	//Don't bother
	if packetSize-1 == 0 {
		return nil, packetSize, packetID, nil
	}
	packet := make([]byte, packetSize-1)
	Conn.Conn.Read(packet)
	return packet, packetSize - 1, packetID, nil
}

func DisplayPacketInfo(PH PacketHeader, Conn *ClientConnection) {
	//DEBUG: output debug info
	if DEBUG {
		Log.Debug("Packet Size: ", PH.packetSize)
		Log.Debug("Packet ID: ", PH.packetID, "State: ", Conn.State)
		Log.Debugf("Packet Contains: %v\n", PH.packet)
		Log.Debug("Protocol: ", PH.protocol)
		Log.Debug("Direction: ", Conn.Direction) //TBD
		fmt.Print("\n")
	}
}

//Authentication moved to Auth.go

func Disconnect(Player string) {
	Log.Debug("Disconnecting Player: ", Player)
	player.Disconnect(Player)
	var p event.Event = event.Player(Player)
	p.PlayerDisconnect()
	EID := player.PlayerEntityMap[Player]
	Tmp := ConnPlayerMap[EID]
	Tmp.Close()
	//P := player.GetPlayerByName(Player)
	//event.PlayerDisconnect(Player)
	//
	// Tmp.Close()
}

//CreateEncryptionRequest moved to CrossProtocol.go
