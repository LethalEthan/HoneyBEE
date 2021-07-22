package server

import (
	"HoneyGO/Packet"
	"net"
)

func GetKeyChain() {
	privateKey = Packet.GetPrivateKey()
	publicKeyBytes = Packet.GetPublicKeyBytes()
	publicKey = Packet.GetPublicKey()
	GotDaKeys = true
}

func GetRun() bool {
	RunMutex.Lock()
	r := Run
	RunMutex.Unlock()
	return r
}
func SetRun(v bool) {
	RunMutex.Lock()
	Run = v
	RunMutex.Unlock()
}

func GetPlayerMap(key string) (string, bool) {
	PlayerMapMutex.RLock()
	P, B := PlayerMap[key]
	PlayerMapMutex.RUnlock()
	return P, B
}

func SetPlayerMap(key string, value string) {
	PlayerMapMutex.Lock()
	PlayerMap[key] = value
	PlayerMapMutex.Unlock()
}

func GetCPM(key uint32) (net.Conn, bool) {
	ConnPlayerMutex.RLock()
	C, B := ConnPlayerMap[key]
	ConnPlayerMutex.RUnlock()
	return C, B
}

func SetCPM(key uint32, value net.Conn) {
	ConnPlayerMutex.Lock()
	ConnPlayerMap[key] = value
	ConnPlayerMutex.Unlock()
}

func GetPCM(key net.Conn) (string, bool) {
	PlayerConnMutex.RLock()
	P, B := PlayerConnMap[key]
	PlayerConnMutex.RUnlock()
	return P, B
}

func SetPCM(key net.Conn, value string) {
	PlayerConnMutex.Lock()
	PlayerConnMap[key] = value
	PlayerConnMutex.Unlock()
}
