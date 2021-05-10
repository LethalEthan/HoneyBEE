package npacket

//Login_0x00_CB - Disconnect (login)
type Login_0x00_CB struct {
	Packet *GeneralPacket
	Reason string
}

//Login_0x01_CB - EncryptionRequest
type Login_0x01_CB struct {
	Packet         *GeneralPacket
	PublicKeyLen   int32
	PublicKey      []byte
	VerifyTokenLen int32
	VerifyToken    []byte
}

//Login_0x02_CB - Login Success
type Login_0x02_CB struct {
	UUID     string
	Username string
}

//Login_0x03_CB - SetCompression
type Login_0x03_CB struct {
	Packet    GeneralPacket
	Threshold int32
}

//Login_0x04_CB - LoadPluginRequest
type Login_0x04_CB struct {
	Packet    GeneralPacket
	MessageID int32
	Channel   string
	Data      []byte
}

///
///Server bound C->S
///

//Login_0x00_SB - Login Start
type Login_0x00_SB struct {
	Packet *GeneralPacket
	Name   string
}

//Login_0x01_SB - EncryptionResponse
type Login_0x01_SB struct {
	Packet          *GeneralPacket
	SharedSecretLen int32
	SharedSecret    []byte
	VerifyTokenLen  int32
	VerifyToken     []byte
}

//Login_0x02_SB - Login Plugin Request
type Login_0x02_SB struct {
	MessageID  int32
	Successful bool
	Data       []byte
}
