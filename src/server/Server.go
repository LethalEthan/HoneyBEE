package server

import (
	"Packet"
	"VarTool"
	config "config"
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"net"
	"player"
	"sync"

	//"time"
	"world"

	logging "github.com/op/go-logging"
	"github.com/pquerna/ffjson/ffjson"
)

//Define used variables
var (
	//Encryption Stuff
	publicKey      *rsa.PublicKey  //Public Key - Used for authentication
	publicKeyBytes []byte          //PublicKey in a byte array for packet delivery and Auth check
	privateKey     *rsa.PrivateKey //PrivateKey
	KeyLength      int             //Keylength used by Encryption Request
	playername     string          //For Authentication
	Log            *logging.Logger //Pretty obvious
	CurrentStatus  *ServerStatus   //ServerStatus Object
	//Encryption stuff again
	//DEBUG                 = true            //Output Debug info?
	GotDaKeys = false //Got dem keys?
	//	ClientSharedSecretLen = 128                       //Initialise CSSL
	//	ClientVerifyTokenLen  = 128                       //Initialise CVTL
	serverID                 = ""                      //this isn't used by mc anymore
	ServerVerifyToken        = make([]byte, 4)         //Initialise a 4 element byte slice of cake
	PlayerMap                = make(map[string]string) //Map Player to UUID
	PlayerMapMutex           = sync.RWMutex{}
	PlayerConnMap            = make(map[net.Conn]string) //Map Connection to Player
	PlayerConnMutex          = sync.RWMutex{}
	ConnPlayerMap            = make(map[uint32]net.Conn) //Map EID to Connection
	ConnPlayerMutex          = sync.RWMutex{}
	GEID              uint32 = 2
	Config            *config.Config
	DEBUG             bool
	KC                = false
	//	pv                 Version
	AvailableProtocols []int32
	//Run control
	Run               bool
	RunMutex          = sync.Mutex{}
	GCPShutdown       = make(chan bool)
	MMODE             bool //Maintenance mode
	ServerInitialised bool
	ServerREINIT      bool
)

const (
	PrimaryMinecraftVersion               = "1.16.3" //Primary Supported MC version
	PrimaryMinecraftProtocolVersion int32 = 753      //Primary Supported MC protocol Version
	ServerVerifyTokenLen                  = 4        //Should always be 4 on notchian servers
)

type PacketHeader struct {
	packet     []byte
	packetSize int32
	packetID   int32
	protocol   int32
	ClientConn *ClientConnection
}

/*type Version interface {
	MCDEFAULT(Conn *ClientConnection)
	MC1_15_2(Conn *ClientConnection)
	MC1_16(Conn *ClientConnection)
	MC1_16_1(Conn *ClientConnection)
	MC1_16_2(Conn *ClientConnection)
	MC1_16_3(Conn *ClientConnection)
}

var GCPShutdown = make(chan bool)*/

func Init() { //this can't be the standard go function init since the logger isn't initialised by the time it's called
	Log.Debug("Server initialising")
	if !GotDaKeys {
		GetKeyChain()
	}
	Config = config.GetConfig()
	DEBUG = Config.Server.DEBUG
	if len(Config.Server.Protocol.AvailableProtocols) != 0 || Config.Server.Protocol.AvailableProtocols != nil {
		AvailableProtocols = Config.Server.Protocol.AvailableProtocols
	}
	StatusSemaphore.Start()
	StatusSemaphore.FlushAndSetSemaphore(StatusCache)
	CurrentStatus = CreateStatusObject(PrimaryMinecraftProtocolVersion, PrimaryMinecraftVersion)
	player.Init()
	ProtocolToVersionInit()
	go world.Init()
	go player.GCPlayer(GCPShutdown)
	if DEBUG {
		Log.Debug("Server initialised")
	}
	ServerInitialised = true
	Run = true
}

func ServerReload() {
	Config = config.GetConfig()
	DEBUG = Config.Server.DEBUG
	player.Init()
	go world.Init()
	GCPShutdown <- true
	go player.GCPlayer(GCPShutdown)
	if DEBUG {
		Log.Debug("Server re-initialised")
	}
	ServerREINIT = true
}

func (PH PacketHeader) MapVersion(Conn *ClientConnection) {
	switch PH.protocol {
	case 578:
		Handle_MC1_15_2(Conn, PH)
		return
	case 735:
		Handle_MC1_16(Conn, PH)
		return
	case 736:
		Handle_MC1_16_1(Conn, PH)
		return
	case 751:
		Handle_MC1_16_2(Conn, PH)
		return
	case 753:
		Handle_MC1_16_3(Conn, PH)
		return
	default:
		Log.Warning("Unsupported protocol:", PH.protocol, "("+ProtocolToVer[PH.protocol]+")", "- sending status and closing connection!")
		HandleUnsupported(Conn, PH, false)
		CloseClientConnection(Conn)
		return
	}
}

//Get/Set moved to GetSet.go

func HandleConnection(Connection *ClientConnection) {
	if !KC {
		GetKeyChain()
		KC = true
	}
	Log.Info("Connection handler initiated")
	//Løøps
	PH := new(PacketHeader)
	var err error
	if Connection != nil {
		for !Connection.isClosed && GetRun() {
			//_ = <-Ticker
			//packet, packetSize, packetID, err := readPacketHeader(Connection)
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
						if err != nil || Hpacket == nil {
							CloseClientConnection(Connection) //You have been terminated
							Log.Error(err.Error())
						}
						Connection.KeepAlive()
						Connection.State = int(Hpacket.NextState)
						PH.protocol = Hpacket.ProtocolVersion
						if Config.DEBUGOPTS.Maintenance {
							MMODE = true
							HandleUnsupported(Connection, *PH, true)
						} else {
							PH.MapVersion(Connection)
						}
						//PH.MapVersion(Connection) //ADD ME BACK
						/*
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
							case 753:
								pv.MC1_16_3(Connection)
								return
							default:
								Log.Warning("Unsupported protocol:", Hpacket.ProtocolVersion, "("+ProtocolToVer[Hpacket.ProtocolVersion]+")", "- sending status and closing connection!")
								pv.MCDEFAULT(Connection)
								CloseClientConnection(Connection)
								return
							}*/
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
							if PH.packetSize == 1 {
								writer := Packet.CreatePacketWriter(0x00)
								marshaledStatus, err := ffjson.Marshal(CurrentStatus) //Sends status via json
								if err != nil {
									Log.Error(err)
									CloseClientConnection(Connection)
									return
								}
								writer.WriteString(string(marshaledStatus))
								SendData(Connection, writer)
							} else {
								CloseClientConnection(Connection)
								return
							}
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
							//--Packet 0x01 End--//
						}
					}
				}
			default:
				{
					Log.Critical("Unkown packet: ", PH.packet, "PHSize: ", PH.packetSize)
					Log.Critical("Contains: ", PH.packet)
				}
			}
		}
	} else {
		Log.Error("Connection pointer is null, skipping!")
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

func HandleUnsupported(Connection *ClientConnection, PH PacketHeader, MMODE bool) {
	Log.Info("Connection handler for Unsupported MC initiated")
	if MMODE != true {
		CurrentStatus = CreateStatusObject(PH.protocol, "Unsupported")
	} else {
		CurrentStatus = CreateStatusObject(PH.protocol, "Maintenance!")
	}
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
						marshaledStatus, err := ffjson.Marshal(CurrentStatus) //Sends status via json
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
					if MMODE != true {
						SendLoginDisconnect(Connection, "Unsupported Protocol, Available protocols are: 1.15.2+")
					} else {
						SendLoginDisconnect(Connection, "Maintenance Mode active!")
					}
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

///
/// Get/Set functions that utilise mutexes moved to GetSet.go
///

///
/// Authentication moved to Auth.go
///

func Disconnect(Player string) {
	Log.Debug("Disconnecting Player: ", Player)
	player.Disconnect(Player)
	// var p event.Event = event.Player(Player)
	// p.PlayerDisconnect()
	EID, _ := player.GetPEM(Player) //PlayerEntityMap[Player]
	Tmp, _ := GetCPM(EID)           //ConnPlayerMap[EID]
	Tmp.Close()
	//P := player.GetPlayerByName(Player)
	//event.PlayerDisconnect(Player)
	//
	// Tmp.Close()
}

///
/// CreateEncryptionRequest moved to CrossProtocol.go

func SendPacket(Connection *ClientConnection, writer *Packet.PacketWriter) {
	Connection.Conn.Write(writer.GetPacket())
}

func ReadPacketHeader(Conn *ClientConnection) ([]byte, int32, int32, error) {
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
