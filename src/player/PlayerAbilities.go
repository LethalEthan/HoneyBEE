package player

import (
	"Packet"

	logging "github.com/op/go-logging"
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

func CreatePlayerAbilities(Conn *ClientConnection) {
	Log := logging.MustGetLogger("HoneyGO")
	Log.Debug("Packet Play, 0x32 Created")
	Conn.KeepAlive()
	T := byte(Invulnerable)
	TT := []byte{T}
	PA := &PlayerAbilities{TT, 0.05, 0.1}
	writer := Packet.CreatePacketWriter(0x32)
	writer.WriteArray(PA.Flags)
	writer.WriteFloat(PA.FlyingSpeed)
	writer.WriteFloat(PA.FOVModifier)
	Log.Debug("Conn state: ", Conn.State)
	SendData(Conn, writer)
	Log.Debug("PlayerAbilities sent")
	CanContinue = true
}
