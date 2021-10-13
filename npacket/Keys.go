package npacket

import (
	"HoneyBEE/config"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"time"
)

var (
	privateKey      *rsa.PrivateKey
	publicKey       *rsa.PublicKey
	publicKeySlice  []byte
	privateKeySlice []byte
	VerifyToken     = make([]byte, 4)
)

const (
	VerifyTokenLen = 4
)

func Keys() {
	var t time.Time
	if config.GConfig.Server.DEBUG {
		t = time.Now()
	}
	privateKey, err = rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		Log.Error(err.Error())
	}
	privateKey.Precompute()
	publicKey = &privateKey.PublicKey
	publicKeySlice, err = x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		panic(err)
	}
	rand.Read(VerifyToken)
	if config.GConfig.Server.DEBUG {
		Log.Info("Took Keys(): ", time.Since(t))
	}
	Log.Info("Key Generated!")
}

func Auth(username string, sharedSecret []byte) string {
	PlayerUUID, autherr := Authenticate(username, "", sharedSecret, publicKeySlice)
	if autherr != nil {
		Log.Error("Auth Fail!")
		return ""
	}
	return PlayerUUID
}
