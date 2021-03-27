package utils

import (
	"runtime"
	"sync"
	"time"

	logging "github.com/op/go-logging"
)

type Semaphore struct {
	Limit   uint32 //Limit of concurrent access
	Data    interface{}
	Channel chan interface{}
	Running bool
	//	Pause   chan bool
	Stopped bool
	Stop    chan bool
	Mutex   sync.Mutex
	Delay   time.Ticker //Add a delay to how the semaphore fills - temporary for now (will probably remove later)
	DataSet chan interface{}
	//	Comm    chan bool
}

var Log = logging.MustGetLogger("HoneyGO")

func CreateSemaphore(Limit uint32) *Semaphore {
	S := new(Semaphore)
	S.Limit = Limit
	S.Channel = make(chan interface{}, Limit)
	//S.Pause = make(chan bool, 1)
	S.Stop = make(chan bool, 1)
	S.Running = true
	//S.Comm = make(chan bool)
	return S
}

func (S *Semaphore) SetData(i interface{}) {
	S.SetStat(false)
	_ = <-S.Channel
	S.Mutex.Lock()
	S.DataSet <- i
	S.Mutex.Unlock()
	S.SetStat(true)
	return
}

func (S *Semaphore) Start() {
	//S.Data = i
	//S.SetRun(true)
	go func() {
		for {
			select {
			case <-S.Stop:
				return
			// case p := <-S.Pause:
			// 	//if config.GConfig.Server.DEBUG {
			// 	if p {
			// 		if config.GConfig.Server.DEBUG {
			// 			Log.Debug("Semaphore: Pause recieved")
			// 		}
			// 	} else {
			// 		if config.GConfig.Server.DEBUG {
			// 			Log.Debug("Semaphore: Unpause recieved")
			// 		}
			// 		//S.Mutex.Lock()
			// 		S.Running = true
			// 		//S.Mutex.Unlock()
			// 		S.Comm <- true
			// 	}
			case d := <-S.DataSet:
				//S.Mutex.Lock()
				S.Data = d
				//S.Mutex.Unlock()
			default:
				if S.GetStat() {
					S.Channel <- S.Data //Will block automatically when the channel is full and re-fill the buffer when needed
				}
			}
		}
	}()
	return
}

func (S *Semaphore) GetStat() bool {
	S.Mutex.Lock()
	P := S.Running
	S.Mutex.Unlock()
	return P
}

func (S *Semaphore) SetStat(b bool) {
	S.Mutex.Lock()
	S.Running = b
	S.Mutex.Unlock()
	return
}

func (S *Semaphore) StopSemaphore() {
	S.Stop <- true
	S.Data = nil
	S.Stopped = true
	S.Running = false
	S = nil
	runtime.GC()
	return
}

//FlushSemaphore - Flush the semaphore to propagate changes to the data interface immediately
func (S *Semaphore) FlushAndSetSemaphore(i interface{}) {
	S.SetStat(false)
	for j := 0; j < 50; j++ { //limit iterations to prevent deadlock
		select {
		case <-S.Channel:
			_ = <-S.Channel
		default:
			S.Data = i
			S.SetStat(true)
			break
		}
		if j > 50 {
			Log.Critical("Semaphore could not flush channel after 50 iterations!")
		}
		return
	}
}

func (S *Semaphore) GetData() interface{} {
	if S.Running == true {
		Log.Debug("Semaphore accessed!")
		CR := <-S.Channel
		Log.Debug("Result:", CR)
		return CR
	}
	return nil
}
