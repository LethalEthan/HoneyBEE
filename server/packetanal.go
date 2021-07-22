package server

import (
	"Packet"
	"VarTool"
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

//var PBUFF = make(map[int32][]byte)
var State int
var close bool
var Data = make(chan net.Conn)

type ServerMITM struct {
	PacketSize int32
	PacketID   int32
	packet     []byte
	Conn       net.Conn
}

type ClientMITM struct {
	PacketSize int32
	packetID   int32
	packet     []byte
	Conn       net.Conn
}

type MITM interface {
	Client()
	Server()
}

func AnalProbe(Client net.Conn, ServerC string) {
	// CConn := CreateClientConnection(Client, HANDSHAKE)
	// SConn := CreateClientConnection(Server, HANDSHAKE)
	Log.Info("I'm here")
	//fmt.Print(Server.RemoteAddr())
	//var err error
	//err = nil
	Server, err := net.Dial("tcp", ServerC)
	if err != nil {
		panic("Could not connect to backend")
	}
	CC := CreateClientConnection(Client, HANDSHAKE)
	SC := CreateClientConnection(Server, HANDSHAKE)
	go Handle(*CC, *SC)
	//Client.SetDeadline(time.Now().Add(time.Duration(1000000000 * 5)))
	//Server.SetDeadline(time.Now().Add(time.Duration(1000000000 * 5)))
	// CB := make(chan bool, 1)
	// SB := make(chan bool, 1)
	//buff := make([]byte, 50)
	// ClientConn := bufio.NewReader(Client)
	// //sbuff := make([]byte, 50)
	// ServerConn := bufio.NewReader(Server)
	// go AnalClient(Client, Server, SB)
	// go AnalServer(Server, Client, SB)
	go Handler(Client, Server /*, CB, SB*/)
	//CB <- true
	return
}

func Handle(C ClientConnection, S ClientConnection) {

}

func Handler(Client net.Conn, Server net.Conn /*, CB chan bool, SB chan bool*/) {
	ClientBuffer := bufio.NewReader(Client)
	//ServerConn := bufio.NewReader(Server)
	Log.Info("Decode plength")
	for {
		result, err := VarTool.ReadVarIntFromBufIO(*ClientBuffer)
		if err != nil {
			Log.Debug("ReadFromBuffer", err)
			return
		}
		Log.Critical("Packet length is: ", result /*, result2*/)
		pbuffer := make([]byte, result+1)
		_, err = io.ReadFull(ClientBuffer, pbuffer)
		Log.Critical("pbuffer now contains: ", pbuffer)
		if err != nil {
			Log.Debug("After ReadFull", err)
		}
		return
	}
	// pbruh, _ := VarTool.ByteDecodeVarInt(PLength)
	// Log.Critical("VARINT: ", pbruh)
	// packet, err := io.ReadFull(ClientConn, buff)
	//PLength := Packet.CreatePacketReader(buff)
}

func AnalServer(Server net.Conn, Client net.Conn, SB chan bool) {
	SPH := new(PacketHeader)
	//var err error
	for {
		Server.SetDeadline(time.Now().Add(time.Duration(1000000000 * 5)))
		//err = nil
		_ = <-SB
		packet, packetSize, packetID, err := MITMreadPacketHeader(Server) //readPacketHeader(SConn)
		SPH = &PacketHeader{packet, packetSize, packetID, 0, nil}
		//if /*SPH.packet != nil*/ err {
		if err != nil {
			//CloseClientConnection(CConn)
			//Log.Error("Connection Terminated: " + err.Error())
			//return
			Server.Close()
			//Client.Close()
			Log.Info("Server Conn closed: ", err)
			return
		}
		Log.Info("-----Server-----")
		MITMDisplayPacketInfo(*SPH)
		Log.Info("-----Server-----")
		writer := Packet.CreatePacketWriter(SPH.packetID)
		//PBUFF := append(writer.CreateVarLong(int64(SPH.packetSize)), SPH.packet...)
		writer.AppendByteSlice(SPH.packet)
		MITMSendData(Client /*SPH.packet*/, writer)
		SB <- true
		// switch State {
		// case HANDSHAKE:
		// 	switch SPH.packetID {
		// 	case 0x00:
		// 		//Server.Close()
		// 	}
		// }
		//Client.SetDeadline(time.Now().Add(time.Duration(1000000000 * 10)))
		//MITMSendData(SConn, PBUFF)
		//}
	}
}

func AnalClient(Client net.Conn, Server net.Conn, SB chan bool) {
	CPH := new(PacketHeader)
	//var err error
	fr := false
	for {
		Client.SetDeadline(time.Now().Add(time.Duration(1000000000 * 5)))
		if fr {
			_ = <-SB
		}
		packet, packetSize, packetID, err := MITMreadPacketHeader(Client) //readPacketHeader(CConn)
		if err != nil {
			//CloseClientConnection(CConn)
			//Log.Error("Connection Terminated: " + err.Error())
			//return
			Log.Info("Client Conn closed: ", err)
			//Server.Close()
			Client.Close()
			return
		}
		CPH = &PacketHeader{packet, packetSize, packetID, 0, nil}
		//if /*CPH.packet != nil*/ err != nil {
		Log.Info("-----Client-----")
		MITMDisplayPacketInfo(*CPH)
		Log.Info("-----Client-----")
		writer := Packet.CreatePacketWriter(CPH.packetID)
		//PBUFF := append(writer.CreateVarLong(int64(CPH.packetSize)), CPH.packet...)
		writer.AppendByteSlice(CPH.packet)
		fr = true
		MITMSendData(Server /*CPH.packet*/, writer)
		SB <- true
		//MITMSendData(SConn, PBUFF)
		//}
	}
}

func MITMSendData(Connection net.Conn /* b []byte*/, writer *Packet.PacketWriter) {
	Connection.Write(writer.GetPacket())
}

func MITMreadPacketHeader(Conn net.Conn) ([]byte, int32, int32, error) {
	//Information used from wiki.vg
	//Read Packet size
	packetSize, err := VarTool.ParseVarIntFromConnection(Conn)
	if err != nil {
		return nil, 0, 0, err //Return nothing on error
	}
	//Handling Legacy Handshake
	if packetSize == 254 && State == HANDSHAKE {
		return nil, 254, 0xFE, nil
	}
	packetID, err := VarTool.ParseVarIntFromConnection(Conn)
	if err != nil {
		return nil, 0, 0, err //Return nothing on error
	}
	//Don't bother
	if packetSize-1 == 0 {
		return nil, packetSize, packetID, nil
	}
	packet := make([]byte, packetSize-1)
	Conn.Read(packet)
	return packet, packetSize - 1, packetID, nil
}

func MITMDisplayPacketInfo(PH PacketHeader) {
	//DEBUG: output debug info
	DEBUG = true
	if DEBUG {
		Log.Debug("Packet Size: ", PH.packetSize)
		Log.Debug("Packet ID: ", PH.packetID, "State: ", State)
		Log.Debugf("Packet Contains: %v\n", PH.packet)
		Log.Debug("Protocol: ", PH.protocol)
		fmt.Print("\n")
	}
}

// func ReadPacket(ClientConn bufio.Reader) (int32, error) {
// 	buff := make([]byte, 1)
// 	var result int32
// 	var numRead byte
// 	var err error
// 	err = nil
// 	for {
// 		buff[0], err = ClientConn.ReadByte()
// 		Log.Info("byte = ", buff)
// 		if err != nil {
// 			Log.Info("error")
// 			return 0, err
// 		}
// 		err = nil
// 		val := int32((buff[0] & 0x7F))
// 		result |= (val << (7 * numRead))
//
// 		numRead++
//
// 		if numRead > 5 {
// 			Log.Info("varint termination error")
// 			return 0, fmt.Errorf("varint was over five bytes without termination")
// 		}
//
// 		if buff[0]&0x80 == 0 {
// 			break
// 		}
// 	}
// 	return result, nil
// }
