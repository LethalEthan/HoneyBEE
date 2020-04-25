package Packet

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"time"

	logging "github.com/op/go-logging"
)

var (
	publicKeyBytes     []byte            //Key stored in byte array for packet delivery
	publicKey          *rsa.PublicKey    //PublicKey
	privateKey         *rsa.PrivateKey   //PrivateKey
	KeyLength          int               //Length of key array (should be 162)
	Encryption         bool              //Finish ConfigHandler
	Log                *logging.Logger   //Logger
	ClientSharedSecret []byte            //Used for Authentication
	ClientVerifyToken  []byte            //Used for Authentication
	ServerVerifyToken  = make([]byte, 4) //Initialise a 4 element byte slice
	DEBUG              = true            //Output debug info?
	err                error             //error interface
)

const (
	ServerVerifyTokenLen = 4
)

//KeyGen - Generates KeyChain
func KeyGen() {
	go keys()
}

//CreateEncryptionRequest - Creates the packet Encryption Request and sends to the client
func CreateEncryptionRequest(Connection *ClientConnection) /*, CH chan bool)*/ {
	Connection.KeepAlive()
	Log := logging.MustGetLogger("HoneyGO")
	Log.Debug("Login State, packetID 0x00")
	Encryption = true //TODO: Finish ConfigHandler

	//Encryption Request
	//--Packet 0x01 S->C Start --//
	if Encryption {
		Log.Debug("Login State, packetID 0x01 Start")
		KeyLength = len(publicKeyBytes)
		//Log.Debug("KeyLength: ", KeyLength)
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
		writer := CreatePacketWriter(0x01)
		writer.WriteString("")                   //Empty;ServerID
		writer.WriteVarInt(int32(KeyLength))     //Key Byte array length
		writer.WriteArray(publicKeyBytes)        //Write Key byte Array
		writer.WriteVarInt(ServerVerifyTokenLen) //Always 4 on notchian servers
		rand.Read(ServerVerifyToken)             // Randomly Generate ServerVerifyToken
		//Log.Debug("ServerVerifyToken: ", ServerVerifyToken)
		writer.WriteArray(ServerVerifyToken)
		SendData(Connection, writer)
		//Packet.LoginPacketCreate(packetSize, reader) //TBD
		//Log.Debug("Encryption Request: ", writer)
		Log.Debug("Encryption Request Sent")
		//CH := make(chan bool)
		//CH <- true
	}
}

func keys() {
	t := time.Now()
	privateKey, err = rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		Log.Error(err.Error())
	}
	privateKey.Precompute()
	publicKey = &privateKey.PublicKey
	publicKeyBytes, err = x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		panic(err)
	}
	Log := logging.MustGetLogger("HoneyGO")
	Log.Info("Took Keys(): ", time.Since(t))
	Log.Info("Key Generated!")
}

//SendData - Sends the data to the client
func SendData(Connection *ClientConnection, writer *PacketWriter) {
	Connection.Conn.Write(writer.GetPacket())
}

//To be able to retrieve the keychain because it runs within a goroutine
func GetPrivateKey() *rsa.PrivateKey {
	return privateKey
}

func GetPublicKey() *rsa.PublicKey {
	return publicKey
}

func GetPublicKeyBytes() []byte {
	return publicKeyBytes
}

func (Conn *ClientConnection) KeepAlive() {
	Conn.Conn.SetDeadline(time.Now().Add(time.Second * 10))
}
