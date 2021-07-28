package server

import (
	"HoneyGO/Packet"
	"net"

	logging "github.com/op/go-logging"
)

var Log = logging.MustGetLogger("HoneyGO")

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

func AnalProbe(Client net.Conn, ServerC string) {
}

func Handler(Client net.Conn, Server net.Conn /*, CB chan bool, SB chan bool*/) {

}

func AnalServer(Server net.Conn, Client net.Conn, SB chan bool) {

}

func AnalClient(Client net.Conn, Server net.Conn, SB chan bool) {
}

func MITMSendData(Connection net.Conn /* b []byte*/, writer *Packet.PacketWriter) {
	Connection.Write(writer.GetPacket())
}

//func MITMreadPacketHeader(Conn net.Conn) ([]byte, int32, int32, error) {
//Information used from wiki.vg
//Read Packet size
// packetSize, err := VarTool.ParseVarIntFromConnection(Conn)
// if err != nil {
// 	return nil, 0, 0, err //Return nothing on error
// }
//Handling Legacy Handshake
// if packetSize == 254 && State == HANDSHAKE {
// 	return nil, 254, 0xFE, nil
// }
// packetID, err := VarTool.ParseVarIntFromConnection(Conn)
// if err != nil {
// 	return nil, 0, 0, err //Return nothing on error
// }
//Don't bother
// if packetSize-1 == 0 {
// 	return nil, packetSize, packetID, nil
// }
// packet := make([]byte, packetSize-1)
// Conn.Read(packet)
// return packet, packetSize - 1, packetID, nil
//}

// func MITMDisplayPacketInfo(PH PacketHeader) {
// 	//DEBUG: output debug info
// 	DEBUG = true
// 	if DEBUG {
// 		Log.Debug("Packet Size: ", PH.packetSize)
// 		Log.Debug("Packet ID: ", PH.packetID, "State: ", State)
// 		Log.Debugf("Packet Contains: %v\n", PH.packet)
// 		Log.Debug("Protocol: ", PH.protocol)
// 		fmt.Print("\n")
// 	}
// }
