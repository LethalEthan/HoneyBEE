package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

const ErrorHash = "00000000000000000000000000000000"

var MD5 string

func UIN() string {
	file, err := os.Open(os.Args[0])
	if err != nil {
		MD5 = ErrorHash
	}
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		MD5 = ErrorHash
	}
	//Get the 16 bytes hash
	hBytes := hash.Sum(nil)[:16]
	file.Close()
	MD5 = hex.EncodeToString(hBytes) //Convert bytes to string
	return MD5
}
