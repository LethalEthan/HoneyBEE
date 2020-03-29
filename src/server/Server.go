package server

import (
	"Packet"
	"VarTool"
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

	logging "github.com/op/go-logging"
)

//Define used variables
var (
	//Encryption Stuff
	publicKey          *rsa.PublicKey  //Pooblic Key
	publicKeyBytes     []byte          //PublicKey in a byte array for packet delivery and Auth check
	privateKey         *rsa.PrivateKey //Like Do I need to comment this?
	Encryption         bool            //TODO: Control via confighandler
	KeyLength          int             //Keylength used by Encryption Request
	ClientSharedSecret []byte          //Cantelope Melon
	ClientVerifyToken  []byte          //Lemons
	//HoneyGO, HoneyComb and should there be a HoneyPot that allows plugins and mods to work together? IDK if that's possible but we can try later
	playername          string                       //For Authentication
	Log                 *logging.Logger              //Pretty obvious
	CurrentStatus       *ServerStatus                //ServerStatus Object
	ClientConnectionMap map[string]*ClientConnection //ClientConnectionMap - Map of connections, duh
	//Encryption stuff again
	GotDaKeys             = false           //Got dem keys?
	ClientSharedSecretLen = 128             //Initialise CSSL
	ClientVerifyTokenLen  = 128             //Initialise CVTL
	serverID              = ""              //Apparently this isn't used by mc anymore
	ServerVerifyToken     = make([]byte, 4) //Initialise a 4 element byte slice of cake
)

//Making these comments is the only way to make this fun I'm sorry lol
const (
	MinecraftVersion         = "1.15.2" //Supported MC version
	MinecraftProtocolVersion = 578      //Supported MC protocol Version
	ServerVerifyTokenLen     = 4        //Should always be 4
)

func GetKeyChain() {
	privateKey = Packet.GetPrivateKey()
	publicKeyBytes = Packet.GetPublicKeyBytes()
	publicKey = Packet.GetPublicKey()
}

func HandleConnection(Connection *ClientConnection) {
	if !GotDaKeys {
		GetKeyChain()
		Log.Debug("Got the keys!")
	}
	Log.Debug("Connection handler initiated")
	//Løøps
	for !Connection.isClosed {
		packet, packetSize, packetID, err := readPacketHeader(Connection)

		if err != nil {
			DestroyClientConnection(Connection)
			Log.Error("Connection Terminated: " + err.Error())
			return
		}
		//DEBUG: print out all values
		Log.Debug("Packet Size: ", packetSize)
		Log.Debug("Packet ID: ", packetID)
		Log.Debugf("Packet Contains: %v\n", packet)
		Log.Debug("")
		//Create Packet Reader
		reader := Packet.CreatePacketReader(packet)
		//Packet Handling
		switch Connection.State {
		case HANDSHAKE: //Handle Handshake
			switch packetID {
			case 0x00:
				{
					//--Packet 0x00 S->C Start--//
					Hpacket, err := Packet.HandshakePacketCreate(packetSize, reader)
					if err != nil {
						DestroyClientConnection(Connection) //You have been terminated
						print(err)
					}
					Connection.KeepAlive() //Ah, ah. ah, ah stayin alive, stayin alive!
					Connection.State = int(Hpacket.NextState)
					break
					//--Packet 0x00 End--//
				}

			case 0xFE:
				{
					//--Packet 0xFE Legacy Ping Request --//
					Log.Warning("Legacy Ping Request received! - Terminated")
					DestroyClientConnection(Connection)
					return
					//--Packet 0xFE End--//
				}
			}

		case STATUS: //Handle Status Request
			{
				switch packetID {
				case 0x00:
					{
						//--Packet 0x00 S->C Start--//
						Connection.KeepAlive()
						writer := Packet.CreatePacketWriter(0x00)
						marshaledStatus, err := json.Marshal(*CurrentStatus) //Sends status via json
						if err != nil {
							Log.Error(err.Error())
							DestroyClientConnection(Connection)
							return
						}
						writer.WriteString(string(marshaledStatus))
						SendData(Connection, writer)
						break
					}

				case 0x01:
					{
						//--Packet 0x01 S->C Start--//
						writer := Packet.CreatePacketWriter(0x01)
						Log.Debug("Status State, packetID 0x01")
						mirror, _ := reader.ReadLong()
						Log.Debug(mirror)
						writer.WriteLong(mirror)
						SendData(Connection, writer)
						DestroyClientConnection(Connection)
						break
						//--Packet 0x01 End--//
					}
				}
			}
		case LOGIN: //Handle Login
			{
				switch packetID {
				case 0x00:
					{
						//NOTE: Cannot be translated via a function due to the goroutine starting before the pointer is ready
						//Causing a crash
						playername, _ = reader.ReadString()
						PE := new(Packet.ClientConnection)
						PE.Conn = Connection.Conn
						PE.State = Connection.State
						go Packet.CreateEncryptionRequest(PE)
						Connection.KeepAlive()
						//--Packet 0x00 C->S Start--//
						Log.Debug("Login State, packetID 0x00")
						//NOTE for ethan:UCB
						break
					}

				case 0x01:
					{
						//--Packet 0x01 C->S Start--//
						Log.Debug("Login State, packetID 0x01")
						p := packet
						ClientSharedSecretLen = 128   //Should always be 128
						ClientSharedSecret = p[2:130] //Find the 128 bytes in the whole byte array
						ClientVerifyToken = p[132:]   //Find the 128 bytes in whole byte array
						Connection.KeepAlive()
						//Decrypt Shared Secret
						decryptSS, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ClientSharedSecret)
						if err != nil {
							fmt.Print(err)
						}
						//Set the decrypted value
						ClientSharedSecret = decryptSS
						ClientSharedSecretLen = len(ClientSharedSecret)
						//Basic check to see whether it's 16 bytes
						if ClientSharedSecretLen != 16 {
							Log.Warning("Shared Secret Length is NOT 16 bytes :(")
						} else {
							Log.Info("ClientSharedSecret Recieved Successfully")
						}
						Log.Info("ClientSharedSecret: ", ClientSharedSecret)

						//Decrypt Verify Token
						decryptVT, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ClientVerifyToken)
						if err != nil {
							fmt.Print(err)
						}

						ClientVerifyTokenLen = len(decryptVT)
						if ClientVerifyTokenLen != ServerVerifyTokenLen {
							Log.Warning("VerifyToken Mismatch!")
						}
						Log.Debug("VT:", decryptVT)
						if ServerVerifyTokenLen != ClientVerifyTokenLen {
							Log.Warning("Encryption Failed!")
						} else {
							Log.Info("Encryption Successful!")
						}
						Auth, err := Authenticate(playername, serverID, ClientSharedSecret, publicKeyBytes)
						Log.Debug("Auth: ", Auth)
						//--Packer 0x01 End--//

						//--Packet 0x02 S->C Start--//
						writer := Packet.CreatePacketWriter(0x02)
						//UUID, err := Authenticate(playername, serverID, ClientSharedSecret, publicKeyBytes)
						Log.Debug("Playername: ", playername)
						if err != nil {
							Log.Error(err)
						}
						writer.WriteString(Auth)
						writer.WriteString(playername)
						SendData(Connection, writer)

						Connection.State = PLAY
						//--Packet 0x02 End--//
						//writer2 := Packet..CreatePacketWriter(0x04) //Compression Not needed, FOR NOW
						Connection.KeepAlive()
						PC := TranslatePlayerStruct(Connection) //Translates Server.ClientConnection -> player.ClientConnection
						//NOTE: goroutines are light weight threads that can be reused with the same stack created before,
						//this will be useful when multiple clients connect but with some added memory usage
						go player.CreateGameJoin(PC) //Creates JoinGame packet AND SetDifficulty via go routines
						//SendData(Connection, writer26)
						break
					}

				case 0x02:
					{
						Log.Debug("Login State, packet 0x02")
						break
					}
				}
			}
		case PLAY:
			{
				switch packetID {
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
					Log.Debug("Packet recieved")
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

func TranslatePlayerStruct(Conn *ClientConnection) *player.ClientConnection {
	PC := new(player.ClientConnection)
	PC.Conn = Conn.Conn
	PC.State = Conn.State
	return PC
}

func TranslatePacketStruct(Conn *ClientConnection) *Packet.ClientConnection {
	PE := new(Packet.ClientConnection)
	PE.Conn = Conn.Conn
	PE.State = Conn.State
	return PE
}

//readPacketHeader - Reads the packet Header and ensures that the packet size is correct, info from wiki.vg
func readPacketHeader(Conn *ClientConnection) ([]byte, int32, int32, error) {
	packetSize, err := VarTool.ParseVarIntFromConnection(Conn.Conn)
	if err != nil {
		return nil, 0, 0, err
	}
	//Information used from wiki.vg
	if packetSize == 254 && Conn.State == HANDSHAKE {
		PreBufferSize := 29
		PreBuffer := make([]byte, PreBufferSize)
		Conn.Conn.Read(PreBuffer)
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

// //keys - Generates a random key that is sent to the client in a byte array
// func keys() (keybytes []byte) {
// 	var err error
// 	privateKey, err = rsa.GenerateKey(rand.Reader, 1024)
// 	if err != nil {
// 		Log.Error(err.Error())
// 	}
// 	privateKey.Precompute()
// 	//privateKey = privateKey
// 	publicKey = &privateKey.PublicKey
// 	publicKeyBytes, err = x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return publicKeyBytes
// }

var ErrorAuthFailed = errors.New("Authentication failed")

type jsonResponse struct {
	ID string `json:"id"`
}

func Authenticate(username string, serverID string, sharedSecret, publicKey []byte) (string, error) {
	//A hash is created using the shared secret and public key and is sent to the mojang sessionserver
	//The server returns the data about the player including the player's skin blob
	//Again I cannot thank enough wiki.vg, this is based off one of the linked java gists by Drew DeVault thank you for the gist that I used to base this off
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
