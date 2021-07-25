package server

import (
	"HoneyGO/Packet"
	"HoneyGO/player"
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"time"

	"github.com/pquerna/ffjson/ffjson"
)

//Protocol things that haven't changed functionally but the protocol ID may have been changed

var (
	EncryptionError = errors.New("Encryption Failed")
	PacketError     = errors.New("Packet Error")
)

type DisconnectChat struct {
	Reason string `json:"text"`
}

//General

//AuthPlayer - Authenticate players against mojang session servers
func AuthPlayer(playername string, ClientSharedSecret []byte) (string, error) {
	var Auth string
	var err error
	if val, tmp := GetPlayerMap(playername); /*PlayerMap[playername]*/ tmp { //checks if map has the value
		Auth = val //Set auth to value
		return Auth, nil
	}
	for i := 0; i <= 3; i++ {
		//4 attempts to get UUID
		Auth, err = Authenticate(playername, serverID, ClientSharedSecret, publicKeyBytes)
		if err != nil {
			Log.Error(err /*"Authentication Failed, trying again"*/)
			time.Sleep(time.Second * 1)
		} else { //If no errors cache uuid in map
			SetPlayerMap(playername, Auth) //PlayerMap[playername] = Auth
			return Auth, nil
		}
	}
	Log.Error("Authentication Failed, trying again - if Authentication fails more than 3 times you're most likely rate limited by mojang")
	return "", err
}

//HANDSHAKE

//STATUS

//LOGIN
//SendLoginDisconnect - Packet 0x00 State: LOGIN
func SendLoginDisconnect(Connection *ClientConnection, Reason string) {
	Log.Debug("Login State, packetID 0x00")
	writer := Packet.CreatePacketWriter(0x00)
	R := new(DisconnectChat)
	R.Reason = Reason
	marshaledDC, err := ffjson.Marshal(R) //Sends status via json
	if err != nil {
		Log.Error(err.Error())
		CloseClientConnection(Connection)
		return
	}
	writer.WriteString(string(marshaledDC))
	SendData(Connection, writer)
	CloseClientConnection(Connection)
}

//CreateEncryptionRequest - Packet 0x01 State: LOGIN
func CreateEncryptionRequest(Connection *ClientConnection) {
	Connection.KeepAlive()
	if DEBUG {
		Log.Debug("Login State, packetID 0x01 Start")
	}
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

//HandleEncryptionResponse - Packet 0x02 State: LOGIN
func HandleEncryptionResponse(PH PacketHeader, Connection *ClientConnection) ([]byte, error) {
	//EncryptionResponse
	Log.Debug("Login State, packetID 0x01")
	Log.Debug("PacketSIZE: ", PH.packetSize)
	if PH.packetSize > 260 {
		return nil, PacketError
	}
	//--//
	ClientSharedSecret := PH.packet[2:130] //Find the first 128 bytes in the whole byte array
	ClientVerifyToken := PH.packet[132:]   //Find the second 128 bytes in whole byte array
	//Decrypt Shared Secret
	ClientSharedSecret, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ClientSharedSecret)
	if err != nil {
		Log.Error(err)
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
	//--//
	//Decrypt Verify Token
	ClientVerifyToken, err = rsa.DecryptPKCS1v15(rand.Reader, privateKey, ClientVerifyToken)
	if err != nil {
		Log.Error(err)
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
	for i := 0; i < len(ServerVerifyToken); i++ {
		if ServerVerifyToken[i] != ClientVerifyToken[i] {
			Log.Warning("Incorrect byte in CVT!")
			return nil, EncryptionError
		}
	}
	//--//
	//--Packet 0x01 S->C Start--//
	//--Authentication--//
	Auth, err := AuthPlayer(playername, ClientSharedSecret)
	if err != nil {
		Log.Error(err)
		SendLoginDisconnect(Connection, "Authentication Failure")
		CloseClientConnection(Connection)
	} else {
		Log.Debug(playername, "[", Auth, "]")
	}
	//--Packer 0x01 End--//

	//--Packet 0x02 S->C Start--//
	writer := Packet.CreatePacketWriter(0x02)
	Log.Debug("Playername: ", playername)
	writer.WriteString(Auth)
	writer.WriteString(playername)
	//time.Sleep(5000000) //DEBUG:Add delay -- remove me later
	SendData(Connection, writer)

	///Entity ID Handling///
	SetPCM(Connection.Conn, playername) //PlayerConnMap[Connection.Conn] = playername //link connection to player
	player.InitPlayer(playername, Auth /*, player.PlayerEntityMap[playername]*/, 1)
	PO, _ := player.GetPEM(playername)
	player.GetPlayerByID(PO)            //player.PlayerEntityMap[playername])
	EID, _ := player.GetPEM(playername) //player.PlayerEntityMap[playername]
	SetCPM(EID, Connection.Conn)        //ConnPlayerMap[EID] = Connection.Conn
	//--//
	CloseClientConnection(Connection)
	Disconnect(playername)
	/* Not enough play logic done so it's not needed for now
	Connection.State = PLAY
	PC := &player.ClientConnection{Connection.Conn, Connection.State, Connection.isClosed}
	player.CreateGameJoin(PC, PO) //player.PlayerEntityMap[playername])
	player.CreateSetDiff(PC)
	player.CreatePlayerAbilities(PC)
	Log.Debug("END")
	CloseClientConnection(Connection)
	Disconnect(playername)*/
	return ClientSharedSecret, nil
}
