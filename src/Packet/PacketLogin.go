package Packet

//type VarInt int32

//LoginPacket Structure (0x01)
type LoginPacket struct {
	ClientSharedSecretLength int
	ClientSharedSecret       []byte
	ClientVerifyTokenLength  int
	ClientVerifyToken        []byte
}

//--!!UNUSED!!--//
//var serverID = ""

/*func LoginPacketCreate(playername string, packet []byte, Conn *ClientConnection) reader *PacketReader) ([]byte, []byte) {
	Conn.KeepAlive()
	lp := new(LoginPacket)
	lp.ClientSharedSecretLength = 128
	lp.ClientSharedSecret = packet[2:(lp.ClientSharedSecretLength + 2)] //[2:130]
	lp.ClientVerifyToken = packet[(lp.ClientSharedSecretLength + 4):]   //[132:]
	//DecryptSS
	decryptSS, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, lp.ClientSharedSecret)
	if err != nil {
		fmt.Print(err)
	}
	lp.ClientSharedSecret = decryptSS
	lp.ClientSharedSecretLength = len(decryptSS)
	Log.Debug("DecryptSS: ", decryptSS)
	//Basic check to see whether it's 16 bytes
	if lp.ClientSharedSecretLength != 16 {
		Log.Warning("Shared Secret Length is NOT 16 bytes :(")

	} else {
		Log.Info("ClientSharedSecret Recieved Successfully")
	}
	Conn.KeepAlive()
	//DecryptVT
	decryptVT, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, lp.ClientVerifyToken)
	if err != nil {
		fmt.Print(err)
	}
	lp.ClientVerifyToken = decryptVT
	lp.ClientVerifyTokenLength = len(decryptVT)
	if ServerVerifyTokenLen != lp.ClientVerifyTokenLength {
		Log.Warning("Encryption Failed!")
	} else {
		Log.Info("Encryption Successful!")
	}
	Authenticate(playername, serverID, ClientSharedSecret, publicKeyBytes)
	// writer := CreatePacketWriter(0x02)
	// SendData(Conn, writer)
	//return lp.ClientSharedSecret, publicKeyBytes
}

var ErrorAuthFailed = errors.New("Authentication failed")

type jsonResponse struct {
	ID string `json:"id"`
}
*/
