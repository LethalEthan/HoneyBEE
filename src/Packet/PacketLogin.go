package Packet

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type VarInt int32

//LoginPacket Structure (0x01)
type LoginPacket struct {
	ClientSharedSecretLength int
	ClientSharedSecret       []byte
	ClientVerifyTokenLength  int
	ClientVerifyToken        []byte
}

var serverID = ""

func LoginPacketCreate(playername string, packet []byte, Conn *ClientConnection) /*reader *PacketReader) ([]byte, []byte) */ {
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
