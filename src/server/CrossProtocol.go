package server

import (
	"Packet"
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"fmt"

	logging "github.com/op/go-logging"
)

//Protocol things that haven't changed functionally but the protocol ID may have been changed

var (
	EncryptionError = errors.New("Encryption Failed")
	PacketError     = errors.New("Packet Error")
)

//func

//HANDSHAKE

//STATUS

//LOGIN
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

	//PacketWrite// NOTE: Later on the packet system will be redone in a more efficient manor where packets will be created in bulk
	writer := Packet.CreatePacketWriter(0x01)
	writer.WriteString("")                   //Empty;ServerID
	writer.WriteVarInt(int32(KeyLength))     //Key Byte array length
	writer.WriteArray(publicKeyBytes)        //Write Key byte Array
	writer.WriteVarInt(ServerVerifyTokenLen) //Always 4 on notchian servers
	rand.Read(ServerVerifyToken)             //Randomly Generate ServerVerifyToken
	writer.WriteArray(ServerVerifyToken)
	SendData(Connection, writer)
	Log.Debug("Encryption Request Sent")
}

func HandleEncryptionResponse(PH *PacketHeader) ([]byte, error) {
	//EncryptionResponse
	Log.Debug("Login State, packetID 0x01")
	Log.Debug("PacketSIZE: ", PH.packetSize)
	if PH.packetSize > 260 {
		return nil, PacketError
	}
	//ClientSharedSecretLen = 128           //Should always be 128 when encrypted
	ClientSharedSecret := PH.packet[2:130] //Find the first 128 bytes in the whole byte array
	ClientVerifyToken := PH.packet[132:]   //Find the second 128 bytes in whole byte array
	//Decrypt Shared Secret
	ClientSharedSecret, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ClientSharedSecret)
	if err != nil {
		fmt.Print(err)
		Log.Debug("Decryption of ClientSharedSecret failed")
		return nil, EncryptionError
	}
	ClientSharedSecretLen := len(ClientSharedSecret)
	//Basic check to see whether it's 16 bytes
	if ClientSharedSecretLen != 16 {
		Log.Warning("Shared Secret Length is NOT 16 bytes :(")
		return nil, EncryptionError
	} else {
		Log.Info("ClientSharedSecret Recieved Successfully")
	}
	//Decrypt Verify Token
	ClientVerifyToken, err = rsa.DecryptPKCS1v15(rand.Reader, privateKey, ClientVerifyToken)
	if err != nil {
		fmt.Print(err)
		Log.Debug("Decryption of ClientVerifyToken failed")
		return nil, EncryptionError
	}
	//Basic Length Check
	ClientVerifyTokenLen := len(ClientVerifyToken)
	if ServerVerifyTokenLen != ClientVerifyTokenLen {
		Log.Warning("Encryption Failed!")
		return nil, EncryptionError
	} else {
		Log.Info("Encryption Successful!")
	}
	//Compare each byte check
	for i := 0; i < 4; i++ {
		if ServerVerifyToken[i] != ClientVerifyToken[i] {
			Log.Warning("Incorrect byte in CVT!")
			return nil, EncryptionError
		}
	}
	return ClientSharedSecret, nil
}
