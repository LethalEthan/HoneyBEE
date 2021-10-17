package server

import (
	"testing"
)

func BenchmarkServer(b *testing.B) {
	// S, err := NewServer("127.0.0.1", "25565", true, false, true, true, 16, 16, 16)
	// if err != nil {
	// 	panic(err)
	// }
	//fmt.Print(S)
	//Conn := new(gnet.Conn)
	// for i := 0; i < b.N; i++ {
	//
	// }
	//S.React(Frame []byte, Conn)
}

func TestServer(t *testing.T) {
	// S, err := NewServer("127.0.0.1", ":25565", true, false, true, true, 16, 16, 16)
	// if err != nil {
	// 	t.Error("Could not verify if server could be started: ", err)
	// } else {
	// 	S.Shutdown()
	// 	return
	// }
	// S.Shutdown()
	t.SkipNow()
}

//
// func Clienttest() {
// 	conn, err := net.Dial("tcp", "127.0.0.1:25565")
// 	if err != nil {
// 		panic("Could not create connection to BenchmarkServer")
// 	}
// 	for {
// 		conn.Write([]byte{16, 0, 242, 5, 9, 49, 50, 55, 46, 48, 46, 48, 46, 49, 99, 226, 1})
// 	}
// }
var TestData = []byte{10, 10, 0, 24, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 100, 105, 109, 101, 110, 115, 105, 111, 110, 95, 116, 121, 112, 101, 8, 0, 4, 116, 121, 112, 101, 0, 24, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 100, 105, 109, 101, 110, 115, 105, 111, 110, 95, 116, 121, 112, 101, 9, 0, 5, 118, 97, 108, 117, 101, 10, 0, 0, 0, 1, 8, 0, 4, 110, 97, 109, 101, 0, 19, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 111, 118, 101, 114, 119, 111, 114, 108, 100, 3, 0, 2, 105, 100, 0, 0, 0, 0, 10, 0, 7, 101, 108, 101, 109, 101, 110, 116, 1, 0, 11, 112, 105, 103, 108, 105, 110, 95, 115, 97, 102, 101, 1, 1, 0, 7, 110, 97, 116, 117, 114, 97, 108, 1, 5, 0, 13, 97, 109, 98, 105, 101, 110, 116, 95, 108, 105, 103, 104, 116, 63, 128, 0, 0, 4, 0, 10, 102, 105, 120, 101, 100, 95, 116, 105, 109, 101, 0, 0, 0, 0, 0, 0, 46, 224, 8, 0, 10, 105, 110, 102, 105, 110, 105, 98, 117, 114, 110, 0, 0, 1, 0, 20, 114, 101, 115, 112, 97, 119, 110, 95, 97, 110, 99, 104, 111, 114, 95, 119, 111, 114, 107, 115, 0, 1, 0, 12, 104, 97, 115, 95, 115, 107, 121, 108, 105, 103, 104, 116, 1, 1, 0, 9, 98, 101, 100, 95, 119, 111, 114, 107, 115, 1, 8, 0, 7, 101, 102, 102, 101, 99, 116, 115, 0, 19, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 111, 118, 101, 114, 119, 111, 114, 108, 100, 1, 0, 9, 104, 97, 115, 95, 114, 97, 105, 100, 115, 1, 3, 0, 5, 109, 105, 110, 95, 121, 0, 0, 0, 0, 3, 0, 6, 104, 101, 105, 103, 104, 116, 0, 0, 1, 0, 3, 0, 14, 108, 111, 103, 105, 99, 97, 108, 95, 104, 101, 105, 103, 104, 116, 0, 0, 0, 64, 5, 0, 16, 99, 111, 111, 114, 100, 105, 110, 97, 116, 101, 95, 115, 99, 97, 108, 101, 63, 128, 0, 0, 1, 0, 9, 117, 108, 116, 114, 97, 119, 97, 114, 109, 0, 1, 0, 11, 104, 97, 115, 95, 99, 101, 105, 108, 105, 110, 103, 0, 0, 0, 0, 10, 0, 24, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 119, 111, 114, 108, 100, 103, 101, 110, 47, 98, 105, 111, 109, 101, 8, 0, 4, 116, 121, 112, 101, 0, 24, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 119, 111, 114, 108, 100, 103, 101, 110, 47, 98, 105, 111, 109, 101, 9, 0, 5, 118, 97, 108, 117, 101, 10, 0, 0, 0, 1, 8, 0, 4, 110, 97, 109, 101, 0, 16, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 112, 108, 97, 105, 110, 115, 3, 0, 2, 105, 100, 0, 0, 0, 0, 10, 0, 7, 101, 108, 101, 109, 101, 110, 116, 8, 0, 13, 112, 114, 101, 99, 105, 112, 105, 116, 97, 116, 105, 111, 110, 0, 4, 110, 111, 110, 101, 5, 0, 5, 100, 101, 112, 116, 104, 63, 128, 0, 0, 5, 0, 11, 116, 101, 109, 112, 101, 114, 97, 116, 117, 114, 101, 63, 128, 0, 0, 5, 0, 5, 115, 99, 97, 108, 101, 63, 128, 0, 0, 5, 0, 8, 100, 111, 119, 110, 102, 97, 108, 108, 63, 128, 0, 0, 8, 0, 8, 99, 97, 116, 101, 103, 111, 114, 121, 0, 6, 112, 108, 97, 105, 110, 115, 10, 0, 7, 101, 102, 102, 101, 99, 116, 115, 3, 0, 9, 115, 107, 121, 95, 99, 111, 108, 111, 114, 0, 127, 161, 255, 3, 0, 15, 119, 97, 116, 101, 114, 95, 102, 111, 103, 95, 99, 111, 108, 111, 114, 0, 127, 161, 255, 3, 0, 9, 102, 111, 103, 95, 99, 111, 108, 111, 114, 0, 127, 161, 255, 3, 0, 11, 119, 97, 116, 101, 114, 95, 99, 111, 108, 111, 114, 0, 127, 161, 255, 3, 0, 13, 102, 111, 108, 105, 97, 103, 101, 95, 99, 111, 108, 111, 114, 0, 127, 161, 255, 3, 0, 11, 103, 114, 97, 115, 115, 95, 99, 111, 108, 111, 114, 0, 127, 161, 255, 8, 0, 20, 103, 114, 97, 115, 115, 95, 99, 111, 108, 111, 114, 95, 109, 111, 100, 105, 102, 105, 101, 114, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type RX struct{}

var R = RX{}

func BenchmarkStructChannel(B *testing.B) {
	channel := make(chan struct{}, 5)
	var RCV = make([]byte, len(TestData))
	go SendRX(channel)
	B.ReportAllocs()
	B.ResetTimer()
	for i := 0; i < B.N; i++ {
		channel <- R
		<-channel
		copy(RCV[0:], TestData)
	}
}

func BenchmarkFrameChannel(B *testing.B) {
	channel := make(chan []byte, 5)
	var RCV = make([]byte, len(TestData))
	go SendFrame(channel)
	B.ReportAllocs()
	B.ResetTimer()
	for i := 0; i < B.N; i++ {
		channel <- TestData
		<-channel
		copy(RCV[0:], TestData)
	}
}

func SendRX(channel chan struct{}) {
	for {
		//channel <- R
	}
}

func SendFrame(channel chan []byte) {
	for {
		//channel <- TestData
	}
}