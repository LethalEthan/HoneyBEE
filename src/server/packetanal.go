package server

import (
	"Packet"
	"VarTool"
	"fmt"
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

func AnalProbe(Client net.Conn, Server net.Conn) {
	// CConn := CreateClientConnection(Client, HANDSHAKE)
	// SConn := CreateClientConnection(Server, HANDSHAKE)
	Log.Info("I'm here")
	fmt.Print(Server.RemoteAddr())
	//var err error
	//err = nil
	go AnalClient(Client, Data)
	go AnalServer(Server, Data)
}

func AnalServer(Server net.Conn, SSend chan net.Conn) {
	SPH := new(PacketHeader)
	var err error
	SPH.packet, SPH.packetSize, SPH.packetID, err = MITMreadPacketHeader(Server) //readPacketHeader(SConn)
	if SPH.packet != nil {
		if err != nil {
			//CloseClientConnection(CConn)
			//Log.Error("Connection Terminated: " + err.Error())
			//return
			Log.Info("b")
		}
		Log.Info("-----Server-----")
		MITMDisplayPacketInfo(*SPH)
		Log.Info("-----Server-----")
		writer := Packet.CreatePacketWriter(SPH.packetID)
		PBUFF := append(writer.CreateVarLong(int64(SPH.packetSize)), SPH.packet...)
		writer.WriteArray(PBUFF)
		//MITMSendData(Client, writer)
		switch State {
		case HANDSHAKE:
			switch SPH.packetID {
			case 0x00:
				//Server.Close()
			}
		}
		//Client.SetDeadline(time.Now().Add(time.Duration(1000000000 * 10)))
		//MITMSendData(SConn, PBUFF)
	}
}

func AnalClient(Client net.Conn, CSend chan net.Conn) {
	CPH := new(PacketHeader)
	var err error
	CPH.packet, CPH.packetSize, CPH.packetID, err = MITMreadPacketHeader(Client) //readPacketHeader(CConn)
	if CPH.packet != nil {
		if err != nil {
			//CloseClientConnection(CConn)
			//Log.Error("Connection Terminated: " + err.Error())
			//return
			Log.Info("b")
		}
		Log.Info("-----Client-----")
		MITMDisplayPacketInfo(*CPH)
		Log.Info("-----Client-----")
		writer := Packet.CreatePacketWriter(CPH.packetID)
		PBUFF := append(writer.CreateVarLong(int64(CPH.packetSize)), CPH.packet...)
		writer.WriteArray(PBUFF)
		//MITMSendData(Server, writer)
		Client.SetDeadline(time.Now().Add(time.Duration(1000000000 * 10)))
		//MITMSendData(SConn, PBUFF)
	}
}

func MITMSendData(Connection net.Conn, writer *Packet.PacketWriter) {
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
	if DEBUG {
		Log.Debug("Packet Size: ", PH.packetSize)
		Log.Debug("Packet ID: ", PH.packetID, "State: ", State)
		Log.Debugf("Packet Contains: %v\n", PH.packet)
		Log.Debug("Protocol: ", PH.protocol)
		fmt.Print("\n")
	}
}
