package utils

import (
	"config"

	logging "github.com/op/go-logging"
)

type Semaphore struct {
	Limit   int32 //Limit of concurrent access
	Data    interface{}
	Channel chan interface{}
	Running bool
	Stop    chan bool
	Stopped bool
}

var Log = logging.MustGetLogger("HoneyGO")

func CreateSemaphore(Limit int32) *Semaphore {
	S := new(Semaphore)
	S.Limit = Limit
	S.Channel = make(chan interface{}, Limit)
	S.Stop = make(chan bool, 1)
	//S.Channel = Chan
	//S.SetRun(false)
	return S
}

func (S *Semaphore) SetData(i interface{}) {
	S.Data = i
	//S.SetRun(true)
	go func() {
		for {
			select {
			case <-S.Stop:
				if config.GConfig.Server.DEBUG {
					Log.Debug("Semaphore: Stop recieved")
				}
				S.Stopped = true
				return
			default:
				S.Channel <- i //Will block automatically when the channel is full and re-fill the buffer when needed
			}
		}
	}()
	return
}

func (S *Semaphore) FlushSemaphore() bool {
	//go func() { S.Stop <- true }()
	S.Stop <- true
	for {
		select {
		case <-S.Channel:
		default:
			S.Stopped = false
			break
		}
		break
	}
	return true
}

func (S *Semaphore) GetData() interface{} {
	Log.Debug("Semaphore accessed!")
	CR := <-S.Channel
	Log.Debug("Result:", CR)
	return CR
}
