package player

import (
	"Packet"
)

type PlayerAbilities struct {
	Flags       []byte
	FlyingSpeed float32
	FOVModifier float32
}

var CanContinue bool

const (
	Invulnerable           = 0x01
	Flying                 = 0x02
	AllowFlying            = 0x04
	CreativeModeInstaBreak = 0x08
)

func CreatePlayerAbilities(Conn *ClientConnection, C chan bool) {
	writer := Packet.CreatePacketWriter(0x32)
	log.Debug("Packet Play, 0x32 Created")
	Conn.KeepAlive()
	T := byte(Invulnerable)
	TT := []byte{T}
	PA := &PlayerAbilities{TT, 0.05, 0.1}
	writer.WriteArray(PA.Flags)
	writer.WriteFloat(PA.FlyingSpeed)
	writer.WriteFloat(PA.FOVModifier)
	log.Debug("Conn state: ", Conn.State)
	wait := <-C
	log.Debug("PA:", wait)
	SendData(Conn, writer)
	log.Debug("PlayerAbilities sent")
	CanContinue = true
}
