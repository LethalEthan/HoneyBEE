package worldtime

import (
	"HoneyGO/Packet"
	"HoneyGO/player"
	"fmt"
	"net"
	"time"

	logging "github.com/op/go-logging"
)

type world_time struct {
	World_Age   int64
	Time_Of_Day int64
}

var Tick bool
var shutdown chan bool
var Log = logging.MustGetLogger("HoneyGO")

func WorldTime(S chan bool) {
	WT := new(world_time)
	// time.NewTimer(50000000)
	ticker := time.NewTicker(1000 * time.Millisecond)
	shutdown = make(chan bool)
	go func() {
		for {
			select {
			case <-S:
				ticker.Stop()
				println("Ticker stopped")
				return
			case t := <-ticker.C:
				//fmt.Println("Tick at", t)
				_ = t
				WT.Time_Of_Day = WT.Time_Of_Day + 20
				go SendTime(WT)
			}
		}
	}()
}

func Shutdown() {
	shutdown <- true
	fmt.Println("Stopping Ticker")
}

func SendTime(WT *world_time) {
	Log.Debug("Sending time")
	TS := time.Now()
	writer := Packet.CreatePacketWriter(0x4F)
	writer.WriteLong(WT.World_Age)
	writer.WriteLong(WT.Time_Of_Day)
	for _, v := range player.OnlinePlayerMap {
		if v {
			print("player online, sending time")
			print(v)
		} else {
			continue
		}
		Elapse := time.Since(TS)
		print("Time send Took: ", Elapse, "µs")
	}
	Elapse := time.Since(TS)
	print("Time send Took: ", Elapse, "µs")
}

func SendData(Connection net.Conn, writer *Packet.PacketWriter) {
	Connection.Write(writer.GetPacket())
}
