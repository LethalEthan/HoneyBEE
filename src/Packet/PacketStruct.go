package Packet

import "net"

//Future Plans for improving packet system
type UncompressedPacket interface {
	Length() int32
	PacketID() int32
}

//Future Plans for improving packet system
type CompressedPacket interface {
	Length() int32
	PacketID() int32
}

//Used PacketCreation
type ClientConnection struct {
	Conn     net.Conn
	State    int
	isClosed bool
}
