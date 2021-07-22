package Packet

import (
	"fmt"
	"net"
)

type PacketHeader struct {
	packet     []byte
	packetSize int32
}

type PacketHandler struct {
	packet     []byte
	packetSize int32
	Connection net.Conn
}

//PacketInbound - Handles packets C->S
type PacketInbound interface {
	//Handshake() //(*HandshakePacket, error)
}

//PacketOutbound - Handles packets S->C
type PacketOutbound interface {
	//Handshake() *HandshakePacket
}

func (PH PacketHandler) Handshake(PackHand *PacketHandler) /*(*HandshakePacket, error)*/ {
	//HP, err := HandshakePacketCreate(length, reader)
	// if err != nil {
	// 	Log.Debug(err)
	// 	return HP, err
	// }
	// Log.Warning("HP:", HP)
	// return HP, nil
	reader := CreatePacketReader(PH.packet)
	HP := new(HandshakePacket)
	HP.length = PH.packetSize
	HP.ProtocolVersion, err = reader.ReadVarInt()
	if err != nil {
		//return nil, err
	}
	//Log.Debug("Protocol Version: ", HP.ProtocolVersion)
	HP.ServerAddress = ""
	HP.ServerPort, err = reader.ReadUnsignedShort()
	if err != nil {
		//return nil, err
	}
	HP.NextState, err = reader.ReadVarInt()
	if err != nil {
		//return nil, err
	}
	print("Handshake: ", HP)
	//return HP, nil
}

func CreatePackHand(packet []byte, packetSize int32, Conn net.Conn) *PacketHandler {
	PackHand := new(PacketHandler)
	PackHand.packet = packet
	PackHand.packetSize = packetSize
	PackHand.Connection = Conn
	return PackHand
}

func HandleInb(PacketInbound interface{}) {
	describe(PacketInbound)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
