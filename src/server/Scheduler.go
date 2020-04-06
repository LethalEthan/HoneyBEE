package server

import (
	"Packet"
	"time"
)

type ScheduleS struct {
	TNow   time.Duration
	Tsince time.Duration
}
type Comms struct {
	isFinished chan bool
	Channel    chan bool
	Channel2   chan int
	TimeElapse time.Duration
}

func Schedule() {
	go TestSched()
	Log.Debug()
}

func CreatePacketSchedule(ID int8, Connection *ClientConnection) {
	var MaxTime = time.Duration(500) //500µs
	var TimeSince = time.Duration(0)
	var TimeNow = time.Now()
	switch ID {
	case 0:
		{
			Log.Debug("Scheduler 0 created")
			var isFinished bool
			Chann := make(chan bool)
			PE := TranslatePacketStruct(Connection)
			go Packet.CreateEncryptionRequest(PE, Chann)
			for {
				isFinished = <-Chann
				if isFinished {
					TimeSince = time.Since(TimeNow)
					Log.Debug("PE: ", TimeSince)
					break
				}
				if TimeSince > MaxTime {
					Log.Debug("OOPS")
					break
				}
				return
			}
			return
		}
	case 1:
		{
			Log.Debug("Scheduler 1 created")
		}
	}
}
func CreateComms() *Comms {
	Comm := new(Comms)
	Comm.isFinished = make(chan bool)
	Comm.Channel = make(chan bool)
	Comm.Channel2 = make(chan int)
	return Comm
}

func LockupCheck() {

}

func GetNumGOroutines(c chan bool, c2 chan int, c3 chan bool) {
	//r := runtime.NumGoroutine()
	time.Sleep(1000000)
	c <- true
	c2 <- 3
	c3 <- true
}

func TestSched() {
	t := time.Now()
	Comm := CreateComms()
	//var wg sync.WaitGroup
	//wg.Add(1)
	//
	go GetNumGOroutines(Comm.Channel, Comm.Channel2, Comm.isFinished)
	// 	for {
	// 		i := <-Comm.Channel
	// 		i2 := <-Comm.Channel
	// 		wg.Done()
	// 		return
	// 	}
	// }()
	//go func() {
	// for {
	//     select {
	//     case <- Comm.Channel:
	//         return
	//     default:
	//         // Do other stuff
	//     }
	// }
	i := <-Comm.Channel
	i2 := <-Comm.Channel2
	if i == true && i2 == 3 {
		Log.Debug("Test Schdeule Success")
	} else {
		Log.Debug("An error has occcured")
	}
	Log.Debug("Took:", time.Since(t))
}

// func TTT() {
// 	var wg sync.WaitGroup
// 	wg.Add(1)
//
// 	ch := make(chan int)
// 	go func() {
// 		for {
// 			foo, ok := <-ch
// 			if !ok {
// 				println("done")
// 				wg.Done()
// 				return
// 			}
// 			println(foo)
// 		}
// 	}()
// 	ch <- 1
// 	ch <- 2
// 	ch <- 3
// 	close(ch)
//
// 	wg.Wait()
// }
