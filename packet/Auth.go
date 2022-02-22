package packet

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
)

var ErrorAuthFailed = /*errors.New*/ ("Authentication failed")
var ErrorHash = "00000000000000000000000000000000"
var SHA string

type jsonResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func Authenticate(username, serverID string, sharedSecret, publicKey []byte) (uuid.UUID, *jsonResponse, error) {
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
		return uuid.Nil, nil, err
	}
	defer response.Body.Close()

	dec := json.NewDecoder(response.Body)
	res := new(jsonResponse)
	// res.Properties = new(properties)
	err = dec.Decode(res)
	if err != nil {
		return uuid.Nil, nil, err
	}
	if len(res.ID) != 32 {
		return uuid.Nil, nil, errors.New(ErrorAuthFailed)
	}
	// res.ID = res.ID[0:8] + "-" + res.ID[8:12] + "-" + res.ID[12:16] + "-" + res.ID[16:20] + "-" + res.ID[20:]
	UUID, err := uuid.Parse(res.ID)
	if err != nil {
		return uuid.Nil, nil, err
	}
	// T, err := UUID.MarshalText()
	Log.Debug("UUID from auth: ", res.ID, res)
	return UUID, res, nil
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

func Hash() string {
	file, err := os.Open(os.Args[0])
	if err != nil {
		SHA = ErrorHash
	}
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		SHA = ErrorHash
	}
	hBytes := hash.Sum(nil)
	file.Close()
	SHA = hex.EncodeToString(hBytes) //Convert bytes to string
	return SHA
}
