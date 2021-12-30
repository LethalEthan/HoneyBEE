package packet

import (
	"crypto/rand"
	"crypto/rsa"

	"github.com/google/uuid"
)

//Login_0x00_CB - Disconnect (login)
type Login_0x00_CB struct {
	Reason string
}

//Login_0x01_CB - EncryptionRequest
type Login_0x01_CB struct {
	PublicKeyLen   int32
	PublicKey      []byte
	VerifyTokenLen int32
	VerifyToken    []byte
}

//Login_0x02_CB - Login Success
type Login_0x02_CB struct {
	UUID     uuid.UUID
	Username string
}

//Login_0x03_CB - SetCompression
type Login_0x03_CB struct {
	Threshold int32
}

//Login_0x04_CB - LoadPluginRequest
type Login_0x04_CB struct {
	MessageID int32
	Channel   string
	Data      []byte
}

///
///Server bound C->S
///

//Login_0x00_SB - Login Start
type Login_0x00_SB struct {
	Name string
}

//Login_0x01_SB - EncryptionResponse
type Login_0x01_SB struct {
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

func (LS *Login_0x00_SB) Decode(PR *PacketReader) error {
	var err error
	if LS.Name, err = PR.ReadString(); err != nil {
		Log.Error(err)
		return err
	}
	return nil
}

func (LERSP *Login_0x01_SB) Decode(PR *PacketReader) error {
	var err error
	LERSP.SharedSecretLen, _, err = PR.ReadVarInt()
	if err != nil {
		Log.Error(err)
		return err
	}
	LERSP.SharedSecret, err = PR.ReadByteArray(int(LERSP.SharedSecretLen))
	if err != nil {
		Log.Error(err)
		return err
	}
	LERSP.VerifyTokenLen, _, err = PR.ReadVarInt()
	if err != nil {
		Log.Error(err)
		return err
	}
	LERSP.VerifyToken, err = PR.ReadByteArray(int(LERSP.VerifyTokenLen))
	if err != nil {
		Log.Error(err)
		return err
	}
	// Decrypt using server private key
	LERSP.SharedSecret, err = rsa.DecryptPKCS1v15(rand.Reader, privateKey, LERSP.SharedSecret)
	if err != nil {
		Log.Error(err)
		return err
	}
	LERSP.VerifyToken, err = rsa.DecryptPKCS1v15(rand.Reader, privateKey, LERSP.VerifyToken)
	if err != nil {
		Log.Error(err)
		return err
	}
	LERSP.VerifyTokenLen = int32(len(LERSP.VerifyToken))
	LERSP.SharedSecretLen = int32(len(LERSP.SharedSecret))
	return nil
}

func (LPR *Login_0x02_SB) Decode(PR *PacketReader) error {
	var err error
	var NR byte
	LPR.MessageID, NR, err = PR.ReadVarInt()
	if err != nil {
		Log.Error(err)
		return err
	}
	LPR.Successful, err = PR.ReadBoolean()
	if err != nil {
		Log.Error(err)
		return err
	}
	LPR.Data, err = PR.ReadByteArray(len(PR.data) - int(NR) - 1)
	if err != nil {
		Log.Error(err)
		return err
	}
	return nil
}

func (LERQ *Login_0x01_CB) Encode(PW *PacketWriter) {
	LERQ.VerifyTokenLen = 4
	LERQ.PublicKeyLen = int32(len(publicKeySlice))
	PW.ResetData(0x01)
	PW.WriteString("")                         // serverID
	PW.WriteVarInt(int32(len(publicKeySlice))) // publickey length
	PW.WriteArray(publicKeySlice)              // publickey
	PW.WriteVarInt(4)                          // verifytoken len
	LERQ.VerifyToken = make([]byte, 4)         // 4 byte token
	rand.Read(LERQ.VerifyToken)                // generate random bytes
	PW.WriteArray(LERQ.VerifyToken)            // verify token
}

func (LoginSucc *Login_0x02_CB) Encode(PW *PacketWriter) error {
	PW.ResetData(0x02)
	PW.WriteUUID(LoginSucc.UUID)
	PW.WriteString(LoginSucc.Username)
	tmp, err := LoginSucc.UUID.MarshalText()
	if err != nil {
		Log.Error(err)
		return err
	}
	Log.Info("Username:", LoginSucc.Username, "UUID:", string(tmp))
	Log.Debug("Sent Login success")
	return nil
}

func (SC *Login_0x03_CB) Encode(PW *PacketWriter, Threshold int32) []byte {
	PW.ResetData(0x03)
	PW.WriteVarInt(Threshold)
	return []byte{0}
}

func (LPR *Login_0x04_CB) Encode() []byte {
	//PW := CreatePacketWriter(0x04)
	// LPR.MessageID = 8
	// LPR.Channel = "Honey"
	// LPR.Data = []byte{0, 0, 0}
	return []byte{0}
}
