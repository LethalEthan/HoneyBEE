package mitm

import (
	"net"
	"time"
)

var Listen net.Listener
var ClientConn net.Conn
var CCounter int
var ClientPackets = make(map[int][]byte)

func StartClient() {
	var err error
	Listen, err = net.Listen("tcp", "127.0.0.1:25565")
	if err != nil {
		panic(err)
	}
	for {
		ClientConn, err = Listen.Accept()
		if err != nil {
			panic(err)
		}
		Start()
		go HandleClient(ClientConn)
	}
}

func HandleClient(conn net.Conn) {
	Log.Debug("Started net handler client")
	conn.SetReadDeadline(time.Now().Add(time.Second * 15))
	conn.SetDeadline(time.Now().Add(time.Second * 15))
	var err error
	//var state int
	var n int
	var r []byte
	data := make([]byte, 2097152)
	for {
		n, err = conn.Read(data)
		if err != nil {
			Log.Debug("Closing because client: ", err)
			CloseClient()
			CloseServer()
			return
		}
		if n <= 0 {
			Log.Critical("Read 0 bytes from client!")
			//goto start
		}
		r = data[0:n]
		if len(r) <= 0 {
			panic("You done fucked up")
		}
		//Log.Critical("Sending: ", r)
		n, err := ServerConn.Write(r)
		if err != nil {
			Log.Critical("SC W: ", err)
			CloseClient()
			CloseServer()
			return
		}
		if n != len(r) || n <= 0 {
			Log.Debug("Bruh HC")
		}
		conn.SetReadDeadline(time.Now().Add(time.Second * 15))
		conn.SetDeadline(time.Now().Add(time.Second * 15))
		ServerConn.SetDeadline(time.Now().Add(time.Second * 15))
		ServerConn.SetReadDeadline(time.Now().Add(time.Second * 15))
	}
}
