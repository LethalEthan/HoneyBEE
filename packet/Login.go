package packet

import (
	"crypto/rand"
	"crypto/rsa"

	"github.com/panjf2000/gnet"
)

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
	Packet     *GeneralPacket
	MessageID  int32
	Successful bool
	Data       []byte
}

func (LS *Login_0x00_SB) Decode() {
	var err error
	PR := CreatePacketReader(LS.Packet.PacketData)
	LS.Name, err = PR.ReadString()
	if err != nil {
		panic(err)
	}
}

func (LERSP *Login_0x01_SB) Decode() {
	var err error
	PR := CreatePacketReader(LERSP.Packet.PacketData)
	LERSP.SharedSecretLen, _, err = PR.ReadVarInt()
	LERSP.SharedSecret, err = PR.ReadByteArray(LERSP.SharedSecretLen)
	LERSP.VerifyTokenLen, _, err = PR.ReadVarInt()
	LERSP.VerifyToken, err = PR.ReadByteArray(LERSP.VerifyTokenLen)
	LERSP.SharedSecret, err = rsa.DecryptPKCS1v15(rand.Reader, privateKey, LERSP.SharedSecret)
	LERSP.VerifyToken, err = rsa.DecryptPKCS1v15(rand.Reader, privateKey, LERSP.VerifyToken)
	if err != nil {
		panic(err)
	}
}

func (LPR *Login_0x02_SB) Decode() {
	var err error
	var NR byte
	PR := CreatePacketReader(LPR.Packet.PacketData)
	LPR.MessageID, NR, err = PR.ReadVarInt()
	LPR.Successful, err = PR.ReadBoolean()
	LPR.Data, err = PR.ReadByteArray(int32(len(LPR.Packet.PacketData) - int(NR) - 1))
	if err != nil {
		panic(err)
	}
}

func (LERQ *Login_0x01_CB) Encode() *PacketWriter {
	PW := CreatePacketWriter(0x01)
	PW.WriteString("")
	PW.WriteVarInt(int32(len(publicKeySlice)))
	PW.WriteArray(publicKeySlice)
	PW.WriteVarInt(int32(len(privateKeySlice)))
	PW.WriteArray(privateKeySlice)
	return PW
}

func (LoginSucc *Login_0x02_CB) Encode() *PacketWriter {
	PW := CreatePacketWriter(0x02)
	PW.WriteString(LoginSucc.UUID)
	PW.WriteString(LoginSucc.Username)
	Log.Info("Username:", LoginSucc.Username, "UUID:", LoginSucc.UUID)
	return PW
}

func (SC *Login_0x03_CB) Encode(Conn *gnet.Conn) *PacketWriter {
	PW := CreatePacketWriter(0x03)
	SC.Threshold = -1
	return PW
}

func (LPR *Login_0x04_CB) Encode(Conn gnet.Conn) {
	PW := CreatePacketWriter(0x04)
	LPR.MessageID = 8
	LPR.Channel = "Honey"
	LPR.Data = []byte{0, 0, 0}
	Conn.AsyncWrite(PW.GetPacket())
}
