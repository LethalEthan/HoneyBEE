package utils

/*
import (
	"runtime"
	"sync"
)

type Semaphore struct {
	Limit   uint32 //Limit of concurrent access
	Data    interface{}
	Channel chan interface{}
	Running bool
	Pause   chan bool
	Stopped bool
	Stop    chan bool
	Mutex   sync.Mutex
	//Delay   time.Duration //Add a delay to how the semaphore fills - temporary for now (will probably remove later)
	DataSet chan interface{}
}

func CreateSemaphore(Limit uint32) *Semaphore {
	S := new(Semaphore)
	S.Limit = Limit
	S.Channel = make(chan interface{}, Limit)
	S.Pause = make(chan bool, 1)
	S.Stop = make(chan bool, 1)
	S.Running = true
	return S
}

//
func (S *Semaphore) SetData(i interface{}) {
	S.SetStat(false)
	S.DataSet <- i
	S.SetStat(true)
	return
}

func (S *Semaphore) Start() error {
	for {
		select {
		case <-S.Stop:
			return nil
		case <-S.Pause:
			_ = <-S.Pause
			print("bruh")
		case d := <-S.DataSet:
			S.Data = d
		default:
			if S.GetStat() {
				S.Channel <- S.Data //Will block automatically when the channel is full and re-fill the buffer when needed
			}
		}
	}
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
	var j uint32
	for j = 0; j < S.Limit; j++ { //limit iterations to prevent deadlock
		select {
		case <-S.Channel:
			_ = <-S.Channel
			print("semaphore flushed: ", j)
		default:
			S.Data = i
			S.SetStat(true)
			break
		}
		return
	}
	if j > S.Limit {
		print("Semaphore could not flush channel after ", S.Limit, " iterations!")
	}
	S.SetStat(true)
}

func (S *Semaphore) GetData() interface{} {
	if S.Running == true {
		//print("Semaphore accessed!")
		CR := <-S.Channel
		//print("Result:", CR)
		return CR
	}
	return nil
}

//FlushSemaphoreNEW - Flush the semaphore to propagate changes to the data interface immediately
func (S *Semaphore) FlushAndSetSemaphoreNEW(i interface{}) {
	S.SetStat(false)
	//S.Stop <- true
	//S.Data = i
	//go S.Start()
	S.Stop <- true
	S.Data = i
	go S.Start()
	// S.Pause <- true
	// S.Data = i
	// S.Pause <- true
	// close(S.Channel)
	// S.Channel = nil
	// S.Channel = make(chan interface{}, S.Limit)
	// S.Pause <- true
	//S.Start()
	S.SetStat(true)
}*/
