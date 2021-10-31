package packet

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestCreatePacketWriter(T *testing.T) {
	PW := CreatePacketWriter(0x00)
	if cap(PW.data) != 128 {
		T.Error(errWriterLength)
	}
}

func TestCreatePacketWriterWithCapacity(T *testing.T) {
	PW := CreatePacketWriterWithCapacity(0x00, 256)
	if cap(PW.data) != 256 {
		T.Error(errWriterLength)
	}
}

func TestCreatePacketWriterWithCapacityBelow0(T *testing.T) {
	PW := CreatePacketWriterWithCapacity(0x00, -45345)
	if cap(PW.data) != 128 {
		T.Error(errWriterLength)
	}
}

func TestWriteByte(T *testing.T) {
	PW := CreatePacketWriter(0x00)
	PW.WriteByte(8)
	if len(PW.data) != 2 {
		T.Error(errWriterLength)
		T.Error(PW.data)
	} else {
		if PW.data[1] != 8 {
			T.Error(errWriterValue)
		}
	}
}
func TestWriteBoolean(T *testing.T) {
	PW := CreatePacketWriter(0x00)
	PW.WriteBoolean(true)
	PW.WriteBoolean(false)
	if PW.data[1] != 1 {
		T.Error(errWriterValue)
	}
	if PW.data[2] != 0 {
		T.Error(errWriterValue)
	}
}
func TestVarInt(T *testing.T) {
	PW := CreatePacketWriter(0x00)
	PW.WriteVarInt(2000000)
	P := PW.GetPacket()
	fmt.Println(P)
	PR := CreatePacketReader(P)
	PS, NR, err := PR.ReadVarInt()
	if err != nil {
		panic(err)
	}
	fmt.Println("PS: ", PS, "NR: ", NR)
	PID, NR, _ := PR.ReadVarInt()
	fmt.Println("PID: ", PID, "NR: ", NR)
	NUM, NR2, _ := PR.ReadVarInt()
	fmt.Println("NUM: ", NUM, "NR: ", NR2)
}

func TestString(T *testing.T) {
	PW := CreatePacketWriter(0x00)
	PW.WriteString("Hello World!")
	fmt.Println(PW.GetPacket())
	PR := CreatePacketReader(PW.GetPacket())
	PS, _, _ := PR.ReadVarInt()
	fmt.Println("Size:", PS)
	PID, _, _ := PR.ReadVarInt()
	fmt.Println("ID: ", PID)
	Test, err := PR.ReadString()
	if err != nil {
		panic(err)
	}
	fmt.Println(Test)
}

func TestThis(T *testing.T) {
	bruh := []byte{2, 36, 52, 101, 100, 48, 99, 53, 53, 100, 45, 99, 97, 97, 57, 45, 52, 54, 54, 57, 45, 56, 98, 101, 97, 45, 102, 51, 97, 48, 48, 53, 50, 101, 54, 102, 49, 102, 12, 76, 101, 116, 104, 97, 108, 69, 116, 104, 97, 110, 56}
	PR := CreatePacketReader(bruh)
	fmt.Println(len(bruh))
	PS, _, _ := PR.ReadVarInt()
	fmt.Println("Size:", PS)
	PID, _, _ := PR.ReadVarInt()
	fmt.Println("ID: ", PID)
	Test, err := PR.ReadString()
	if err != nil {
		panic(err)
	}
	fmt.Println(Test)
	Test2, err := PR.ReadString()
	fmt.Println(Test2)
}

func TestLoginSuccess(T *testing.T) {
	LS := new(Login_0x02_CB)
	var PW *PacketWriter
	var err error
	LS.UUID, err = uuid.Parse("4ed0c55d-caa9-4669-8bea-f3a0052e6f1f")
	if err != nil {
		T.Error(err)
		panic(err)
	}
	LS.Username = "LethalEthan8"
	PW = LS.Encode()
	PR := CreatePacketReader(PW.GetPacket())
	PS, _, _ := PR.ReadVarInt()
	fmt.Println("Size:", PS)
	PID, _, _ := PR.ReadVarInt()
	fmt.Println("ID: ", PID)
	Test, err := PR.ReadString()
	if err != nil {
		panic(err)
	}
	fmt.Println(Test)
	Test2, err := PR.ReadString()
	if err != nil {
		panic(err)
	}
	fmt.Println(Test2)
}
