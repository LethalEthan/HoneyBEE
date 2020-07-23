package server

import (
	"Packet"
	"VarTool"
	config "config"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"player"
	"strings"
	"time"

	logging "github.com/op/go-logging"
)

//Define used variables
var (
	//Encryption Stuff
	publicKey          *rsa.PublicKey  //Public Key - Used for authentication
	publicKeyBytes     []byte          //PublicKey in a byte array for packet delivery and Auth check
	privateKey         *rsa.PrivateKey //PrivateKey
	KeyLength          int             //Keylength used by Encryption Request
	ClientSharedSecret []byte          //CSS
	ClientVerifyToken  []byte          //CVT
	playername         string          //For Authentication
	Log                *logging.Logger //Pretty obvious
	CurrentStatus      *ServerStatus   //ServerStatus Object
	//Encryption stuff again
	//DEBUG                 = true            //Output Debug info?
	GotDaKeys             = false                     //Got dem keys?
	ClientSharedSecretLen = 128                       //Initialise CSSL
	ClientVerifyTokenLen  = 128                       //Initialise CVTL
	serverID              = ""                        //this isn't used by mc anymore
	ServerVerifyToken     = make([]byte, 4)           //Initialise a 4 element byte slice of cake
	PlayerMap             = make(map[string]string)   //Map Player to UUID
	PlayerConnMap         = make(map[net.Conn]string) //Map Connection to Player
	ConnPlayerMap         = make(map[uint32]net.Conn)
	//	EntityPlayerMap        = player.PlayerEntityMap    //= make(map[string]uint32)   //Map Player to EID
	GEID   uint32 = 2
	Config *config.Config
)

//REFERENCE: Play state goes from GameJoin.go -> SetDifficulty.go -> PlayerAbilities.go via goroutines
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
}

func HandleConnection(Connection *ClientConnection) {
	if !GotDaKeys {
		GetKeyChain()
	}
	Config = config.GetConfig()
	DEBUG := Config.Server.DEBUG
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
		if DEBUG {
			Log.Debug("Packet Size: ", PH.packetSize)
			Log.Debug("Packet ID: ", PH.packetID, "State: ", Connection.State)
			Log.Debugf("Packet Contains: %v\n", PH.packet)
			Log.Debug("Direction: ", Connection.Direction) //TBD
			fmt.Print("")
		}
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
					//Hpacket := Packet.PacketOutbound(packetSize, reader)
					//Log.Warning("Hpackettt: ", HP)
					if err != nil {
						CloseClientConnection(Connection) //You have been terminated
						Log.Error(err.Error())
					}
					Connection.KeepAlive()
					Proto, _ := reader.ReadVarInt()
					Log.Debug("Protocol:", Proto)
					Connection.State = int(Hpacket.NextState)
					break
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
		case LOGIN: //Handle Login
			{
				switch PH.packetID {
				case 0x00:
					{
						//--Packet 0x00 S->C Start--//
						Log.Debug("Login State, packetID 0x00")
						Connection.KeepAlive()
						playername, _ = reader.ReadString()
						//PE := TranslatePacketStruct(Connection)
						//go Packet.CreateEncryptionRequest(PE)
						go CreateEncryptionRequest(Connection)
						break
					}
				case 0x01:
					{
						//--Packet 0x01 C->S Start--//
						//EncryptionResponse
						Connection.KeepAlive()
						Log.Debug("Login State, packetID 0x01")
						Log.Debug("PacketSIZE: ", PH.packetSize)
						ClientSharedSecretLen = 128           //Should always be 128
						ClientSharedSecret = PH.packet[2:130] //Find the 128 bytes in the whole byte array
						ClientVerifyToken = PH.packet[132:]   //Find the 128 bytes in whole byte array
						Connection.KeepAlive()
						//Decrypt Shared Secret
						decryptSS, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ClientSharedSecret)
						if err != nil {
							fmt.Print(err)
						}
						ClientSharedSecret = decryptSS //Set the decrypted value
						ClientSharedSecretLen = len(ClientSharedSecret)
						//Basic check to see whether it's 16 bytes
						if ClientSharedSecretLen != 16 {
							Log.Warning("Shared Secret Length is NOT 16 bytes :(")
						} else {
							Log.Info("ClientSharedSecret Recieved Successfully")
						}
						//Decrypt Verify Token
						decryptVT, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ClientVerifyToken)
						if err != nil {
							fmt.Print(err)
						}
						Connection.KeepAlive()
						ClientVerifyTokenLen = len(decryptVT)
						if ServerVerifyTokenLen != ClientVerifyTokenLen {
							Log.Warning("Encryption Failed!")
						} else {
							Log.Info("Encryption Successful!")
						}
						var Auth string
						//--Authentication Stuff--//
						//Authenticate Player
						//Check if playermap has any data -- UUID Caching
						if val, tmp := PlayerMap[playername]; tmp { //checks if map has the value
							Auth = val //Set auth to value
						} else { //If uuid isn't found, get it
							//2 attempts to get UUID
							Auth, err = Authenticate(playername, serverID, ClientSharedSecret, publicKeyBytes)
							if err != nil {
								Log.Error("Authentication Failed, trying second time")
								Auth, err = Authenticate(playername, serverID, ClientSharedSecret, publicKeyBytes)
								if err != nil {
									Log.Error("Authentication failed on second attempt, closing connection")
									CloseClientConnection(Connection)
								} else { //If no errors cache uuid in map
									PlayerMap[playername] = Auth
								}
							} else { //If no erros cache uuid in map
								PlayerMap[playername] = Auth
							}
						}
						Log.Debug(playername, "[", Auth, "]")
						//--Packer 0x01 End--//

						//--Packet 0x02 S->C Start--//
						writer := Packet.CreatePacketWriter(0x02)
						Log.Debug("Playername: ", playername)
						writer.WriteString(Auth)
						writer.WriteString(playername)
						//UUID Cache
						//DEBUG: REMOVE ME
						Log.Debug("PlayerMap: ", PlayerMap)
						Log.Debug("PlayerData:", PlayerMap[playername])
						time.Sleep(5000000) //DEBUG:Add delay -- remove me later
						SendData(Connection, writer)

						///Entity ID Handling///
						PlayerConnMap[Connection.Conn] = playername //link connection to player
						player.InitPlayer(playername, Auth /*, player.PlayerEntityMap[playername]*/, 1)
						player.GetPlayerByID(player.PlayerEntityMap[playername])
						EID := player.PlayerEntityMap[playername]
						ConnPlayerMap[EID] = Connection.Conn
						//go player.GCPlayer() //DEBUG: REMOVE ME LATER
						//--//
						Connection.State = PLAY
						//worldtime.
						PC := TranslatePlayerStruct(Connection) //Translates Server.ClientConnection -> player.ClientConnection
						//NOTE: goroutines are light weight threads that can be reused with the same stack created before,
						//this will be useful when multiple clients connect but with some added memory usage
						C := make(chan bool)
						go player.CreateGameJoin(PC, C, player.PlayerEntityMap[playername]) //Creates JoinGame packet AND SetDifficulty AND Player Abilities via go routines
						Log.Debug("END")
						//time.Sleep(60000000)
						//CloseClientConnection(Connection)
						break
					}
				case 0x02:
					{
						Log.Debug("Login State, packet 0x02")
						break
					}
				case 0x05:
					{
						Log.Debug("Packet 0x05")
						break
					}
				}
			}
			//Play will be handled by another package/function
		case PLAY:
			{
				switch PH.packetID {
				case 0x00:
					{
						Log.Debug("Play State, packet 0x00")
						break
					}
				case 0x01:
					{
						Log.Debug("Play State, Packet 0x01")
						break
					}
				case 0x02:
					{
						Log.Debug("Play State, Packet 0x02")
						break
					}
				case 0x03:
					{
						Log.Debug("Play State, Packet 0x03")
						break
					}
				case 0x04:
					{
						Log.Debug("Play State, Packet 0x04")
						break
					}
				case 0x05:
					{
						Log.Debug("Play State, Packet 0x05")
						break
					}
				default:
					for {
						Log.Fatal("Play Packet recieved")
					}
				}
			}
		default:
			Log.Debug("oo")
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

//Server ClientConn -> Player ClientConn
func TranslatePlayerStruct(Conn *ClientConnection) *player.ClientConnection {
	PC := new(player.ClientConnection)
	PC.Conn = Conn.Conn
	PC.State = Conn.State
	PC.Closed = Conn.isClosed
	return PC
}

//Server ClientConn -> Packet ClientConn
func TranslatePacketStruct(Conn *ClientConnection) *Packet.ClientConnection {
	PE := new(Packet.ClientConnection)
	PE.Conn = Conn.Conn
	PE.State = Conn.State
	PE.Closed = Conn.isClosed
	return PE
}

//readPacketHeader - Reads the packet Header for Packet ID and size info
func readPacketHeader(Conn *ClientConnection) ([]byte, int32, int32, error) {
	//Read Packet size
	packetSize, err := VarTool.ParseVarIntFromConnection(Conn.Conn)
	if err != nil {
		return nil, 0, 0, err //Return nothing on error
	}
	//Information used from wiki.vg
	//Handling Handshake
	if packetSize == 254 && Conn.State == HANDSHAKE {
		PreBufferSize := 30 //Create PreBuffer
		PreBuffer := make([]byte, PreBufferSize)
		Conn.Conn.Read(PreBuffer)
		//PostBuffer -> Big endian
		PostBufferSize := int(binary.BigEndian.Uint16(PreBuffer[25:]))
		PostBuffer := make([]byte, PostBufferSize)
		size := PreBufferSize + PostBufferSize
		Conn.Conn.Read(PostBuffer)
		return append(PreBuffer, PostBuffer...), int32(size), 0xFE, nil
	}

	packetID, err := VarTool.ParseVarIntFromConnection(Conn.Conn)

	if err != nil {
		return nil, 0, 0, err
	}

	if packetSize-1 == 0 {
		return nil, packetSize, packetID, nil
	}
	packet := make([]byte, packetSize-1)
	Conn.Conn.Read(packet)
	return packet, packetSize - 1, packetID, nil
}

var ErrorAuthFailed = errors.New("Authentication failed")

type jsonResponse struct {
	ID string `json:"id"`
}

func Authenticate(username string, serverID string, sharedSecret, publicKey []byte) (string, error) {
	//A hash is created using the shared secret and public key and is sent to the mojang sessionserver
	//The server returns the data about the player including the player's skin blob
	//Again I cannot thank enough wiki.vg, this is based off one of the linked java gists by Drew DeVault; thank you for the gist that I used to base this off
	sha := sha1.New()
	sha.Write([]byte(serverID))
	sha.Write(sharedSecret)
	sha.Write(publicKey)
	hash := sha.Sum(nil)

	negative := (hash[0] & 0x80) == 0x80
	if negative {
		twosCompliment(hash)
	}

	buf := hex.EncodeToString(hash)
	if negative {
		buf = "-" + buf
	}
	hashString := strings.TrimLeft(buf, "0")

	response, err := http.Get(fmt.Sprintf("https://sessionserver.mojang.com/session/minecraft/hasJoined?username=%s&serverId=%s", username, hashString))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	dec := json.NewDecoder(response.Body)
	res := &jsonResponse{}
	err = dec.Decode(res)
	if err != nil {
		return "", ErrorAuthFailed
	}

	if len(res.ID) != 32 {
		return "", ErrorAuthFailed
	}
	hyphenater := res.ID[0:8] + "-" + res.ID[8:12] + "-" + res.ID[12:16] + "-" + res.ID[16:20] + "-" + res.ID[20:]
	res.ID = hyphenater
	return res.ID, nil
}

func twosCompliment(p []byte) {
	carry := true
	for i := len(p) - 1; i >= 0; i-- {
		p[i] = ^p[i]
		if carry {
			carry = p[i] == 0xFF
			p[i]++
		}
	}
}

func Disconnect(Player string) {
	Log.Debug("Disconnecting Player: ", Player)
	player.Disconnect(Player)
}

//--Encryption Response/Request--//
func CreateEncryptionResponse(Connection *ClientConnection, PH *PacketHeader) {
	//--Packet 0x01 C->S Start--//
	Connection.KeepAlive()
	Log.Debug("Login State, packetID 0x01")
	Log.Debug("PacketSIZE: ", PH.packetSize)
	ClientSharedSecretLen = 128           //Should always be 128
	ClientSharedSecret = PH.packet[2:130] //Find the 128 bytes in the whole byte array
	ClientVerifyToken = PH.packet[132:]   //Find the 128 bytes in whole byte array
	Connection.KeepAlive()
	//Decrypt Shared Secret
	decryptSS, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ClientSharedSecret)
	if err != nil {
		fmt.Print(err)
	}
	ClientSharedSecret = decryptSS //Set the decrypted value
	ClientSharedSecretLen = len(ClientSharedSecret)
	//Basic check to see whether it's 16 bytes
	if ClientSharedSecretLen != 16 {
		Log.Warning("Shared Secret Length is NOT 16 bytes :(")
	} else {
		Log.Info("ClientSharedSecret Recieved Successfully")
	}
	//Decrypt Verify Token
	decryptVT, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ClientVerifyToken)
	if err != nil {
		fmt.Print(err)
	}
	Connection.KeepAlive()
	ClientVerifyTokenLen = len(decryptVT)
	if ServerVerifyTokenLen != ClientVerifyTokenLen {
		Log.Warning("Encryption Failed!")
	} else {
		Log.Info("Encryption Successful!")
	}
	var Auth string
	//Authenticate Player
	//Check if playermap has any data
	if val, tmp := PlayerMap[playername]; tmp { //checks if map has the value
		Auth = val //Set auth to value
	} else { //If uuid isn't found, get it
		//2 attempts to get UUID
		for i := 0; i <= 2; i++ {
			Auth, err = Authenticate(playername, serverID, ClientSharedSecret, publicKeyBytes)
			if err != nil {
				Log.Error("Authentication Failed, trying ", i, " time")
				//Auth, err = Authenticate(playername, serverID, ClientSharedSecret, publicKeyBytes)
			} else {
				PlayerMap[playername] = Auth
				break
			}
		}
	}
	Log.Debug(playername, "[", Auth, "]")
	//--Packer 0x01 End--//

	//--Packet 0x02 S->C Start--//
	writer := Packet.CreatePacketWriter(0x02)
	Log.Debug("Playername: ", playername)
	writer.WriteString(Auth)
	writer.WriteString(playername)
	//UUID Cache
	//DEBUG: REMOVE ME
	Log.Debug("PlayerMap: ", PlayerMap)
	Log.Debug("PlayerData:", PlayerMap[playername])
	//time.Sleep(5000000) //Debug:Add delay
	SendData(Connection, writer)

	///Entity ID Handling///
	PlayerConnMap[Connection.Conn] = playername //link connection to player
	player.InitPlayer(playername, Auth /*, player.PlayerEntityMap[playername]*/, 1)
	player.GetPlayerByID(player.PlayerEntityMap[playername])
	//go player.GCPlayer() //DEBUG: REMOVE ME LATER
	//--//
	Connection.State = PLAY
	PC := TranslatePlayerStruct(Connection) //Translates Server.ClientConnection -> player.ClientConnection
	//NOTE: goroutines are light weight threads that can be reused with the same stack created before,
	//this will be useful when multiple clients connect but with some added memory usage
	C := make(chan bool)
	go player.CreateGameJoin(PC, C, player.PlayerEntityMap[playername]) //Creates JoinGame packet AND SetDifficulty AND Player Abilities via go routines
	Log.Debug("END")
	//time.Sleep(60000000)
	//CloseClientConnection(Connection)
}

func CreateEncryptionRequest(Connection *ClientConnection) {
	Connection.KeepAlive()
	Log := logging.MustGetLogger("HoneyGO")
	Log.Debug("Login State, packetID 0x00")

	//Encryption Request
	//--Packet 0x01 S->C Start --//
	Log.Debug("Login State, packetID 0x01 Start")
	KeyLength = len(publicKeyBytes)
	//KeyLength Checks
	if KeyLength > 162 {
		Log.Warning("Key is bigger than expected!")
	}
	if KeyLength < 162 {
		Log.Warning("Key is smaller than expected!")
	} else {
		Log.Debug("Key Generated Successfully")
	}

	//PacketWrite - // NOTE: Later on the packet system will be redone in a more efficient manor where packets will be created in bulk
	writer := Packet.CreatePacketWriter(0x01)
	writer.WriteString("")                   //Empty;ServerID
	writer.WriteVarInt(int32(KeyLength))     //Key Byte array length
	writer.WriteArray(publicKeyBytes)        //Write Key byte Array
	writer.WriteVarInt(ServerVerifyTokenLen) //Always 4 on notchian servers
	rand.Read(ServerVerifyToken)             // Randomly Generate ServerVerifyToken
	writer.WriteArray(ServerVerifyToken)
	SendData(Connection, writer)
	Log.Debug("Encryption Request Sent")
}
