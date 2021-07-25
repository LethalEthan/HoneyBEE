package Packet

import (
	"HoneyGO/config"
	"fmt"
	"net"
)

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
	Conn      net.Conn
	State     int
	Closed    bool
	Direction string
}

var SConfig *config.Config

type PaxI interface {
	Test()
}

func (HP HandshakePacket) Test() {
	fmt.Print(HP)
}

func (PH PacketHeader) Length() {
	fmt.Print(PH.packetSize)
}
