package nserver

import (
	"fmt"
	"testing"
)

func BenchmarkServer(b *testing.B) {
	S, err := NewServer("127.0.0.1", "25565", true, false, true, true, 16, 16, 16)
	if err != nil {
		panic(err)
	}
	fmt.Print(S)
	//Conn := new(gnet.Conn)
	for i := 0; i < b.N; i++ {

	}
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
