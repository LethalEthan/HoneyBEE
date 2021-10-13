package server

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
		// 	pr := npacket.CreatePacketReader(r)
		// 	_, _, err := pr.ReadVarInt()
		// 	ID, _, err := pr.ReadVarInt()
		// 	if err != nil {
		// 		err = nil
		// 		goto skip
		// 	}
		// 	if ID == 0x00 && state == 0 {
		// 		PV, _, err := pr.ReadVarInt()
		// 		_, err = pr.ReadString()
		// 		SP, err := pr.ReadUnsignedShort()
		// 		NS, _, err := pr.ReadVarInt()
		// 		if err != nil {
		// 			Log.Critical("BRUHHHHH")
		// 			state = 5
		// 			goto skip
		// 		}
		// 		state = int(NS)
		// 		pw := npacket.CreatePacketWriter(0x00)
		// 		pw.WriteVarInt(PV)
		// 		pw.WriteString("play.moonglade.pw")
		// 		pw.WriteUnsignedShort(SP)
		// 		pw.WriteVarInt(NS)
		// 		r = pw.GetPacket()
		// 	}
		// skip:
		Log.Critical("Sending: ", r)
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
