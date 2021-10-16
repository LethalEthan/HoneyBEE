package server

import (
	"net"
	"time"
)

var ServerConn net.Conn
var ServerPackets = make(map[int][]byte)
var SCounter int

func Start() {
	var err error
	ServerConn, err = net.Dial("tcp", "127.0.0.1:25567") //mc-central.net:25565") //192.168.0.56:25570") //"127.0.0.1:25570")
	if err != nil {
		Log.Critical("Error dialing")
		panic(err)
	}
	Log.Debug("Connected handling...")
	go HandleServer(ServerConn)
}

func HandleServer(conn net.Conn) {
	Log.Debug("Started net handler server")
	conn.SetReadDeadline(time.Now().Add(time.Second * 15))
	conn.SetDeadline(time.Now().Add(time.Second * 15))
	var err error
	//var state int
	var n int
	var r []byte
	var data []byte
	data = make([]byte, 2097152)
	for {
		// size, ba, err := VarTool.ParseVarIntFromConnection(conn)
		// if err != nil {
		// 	Log.Debug("Closing because VI: ", err)
		// 	return
		// }
		// if size > 2097151 || size < 0 {
		// 	Log.Debug("Closing because size is greater than 2097151", " Got size: ", size)
		// 	CloseClient()
		// 	CloseServer()
		// 	return
		// }
		// err = nil
	start:
		n, err = conn.Read(data)
		if err != nil {
			if err.Error() == "EOF" {
				goto start
			}
			Log.Debug("Closing because server: ", err)
			CloseClient()
			CloseServer()
			return
		}
		//data = append(ba, data...)
		if n <= 0 {
			//goto start
			Log.Critical("Read 0 bytes from server!")
			CloseServer()
			CloseClient()
			return
		}
		r = data[:n]
		//Log.Debug("DATAMITM: ", r)
		n, err := ClientConn.Write(r)
		if err != nil {
			Log.Critical("CC W: ", err)
			CloseClient()
			CloseServer()
			return
		}
		if n != len(r) || n <= 0 {
			Log.Critical("Bruh HS")
		}
		conn.SetReadDeadline(time.Now().Add(time.Second * 15))
		conn.SetDeadline(time.Now().Add(time.Second * 15))
		ServerConn.SetDeadline(time.Now().Add(time.Second * 15))
		ServerConn.SetReadDeadline(time.Now().Add(time.Second * 15))
	}
}

func CloseServer() {
	ServerConn.Close()
	Log.Debug("Closed Server")
}

func CloseClient() {
	ClientConn.Close()
	Log.Debug("Closed Client")
}
