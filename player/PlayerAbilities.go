package player

import (
	"HoneyGO/Packet"
	"time"
)

type PlayerAbilities struct {
	Flags       int8
	FlyingSpeed float32
	FOVModifier float32
}

var start time.Time

const (
	Invulnerable           = 0x01
	Flying                 = 0x02
	AllowFlying            = 0x04
	CreativeModeInstaBreak = 0x08
	DEBUG                  = false
)

func CreatePlayerAbilities(Conn *ClientConnection) {
	if DEBUG {
		start = time.Now()
	}
	writer := Packet.CreatePacketWriter(0x32)
	//log.Debug("Packet Play, 0x32 Created")
	PA := &PlayerAbilities{0x01, 0.05, 0.1}
	writer.WriteByte(PA.Flags)
	writer.WriteFloat(PA.FlyingSpeed)
	writer.WriteFloat(PA.FOVModifier)
	// if DEBUG {
	// 	elapse := time.Since(start)
	// 	log.Debug("Time before wait block:", elapse)
	// }
	// wait := <-C //Blocks goroutine until value is recieved
	// log.Debug("PA:", wait)
	// if DEBUG {
	// 	elapse := time.Since(start)
	// 	log.Warning("Time after wait block:", elapse)
	// }
	SendData(Conn, writer)
	//log.Debug("PlayerAbilities sent")
}
