package npacket

import (
	"fmt"
	"testing"
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
