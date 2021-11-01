package packet

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

func TestVarLong(T *testing.T) {
	PW := CreatePacketWriter(0x00)
	PW.WriteVarLong(9223372036854775807)
	P := PW.GetPacket()
	fmt.Println(P)
	PR := CreatePacketReader(P)
	PS, _, err := PR.ReadVarInt()
	if err != nil {
		T.Error(err)
	}
	fmt.Println("PS: ", PS)
	PID, _, err := PR.ReadVarInt()
	if err != nil {
		T.Error(err)
	}
	fmt.Println("PID: ", PID)
	NUM, err := PR.ReadVarLong()
	if err != nil {
		T.Error(err)
	}
	if NUM != 9223372036854775807 {
		T.Error(errWriterValue)
	}
	fmt.Println("NUM: ", NUM)
}

func TestString(T *testing.T) {
	PW := CreatePacketWriter(0x00)
	PW.WriteString("Hello World!")
	fmt.Println(PW.GetPacket())
	PR := CreatePacketReader(PW.GetPacket())
	PS, _, err := PR.ReadVarInt()
	if err != nil {
		T.Error(err)
	}
	fmt.Println("Size:", PS)
	PID, _, err := PR.ReadVarInt()
	if err != nil {
		T.Error(err)
	}
	fmt.Println("ID: ", PID)
	Test, err := PR.ReadString()
	if err != nil {
		T.Error(err)
	}
	fmt.Println(Test)
}

func TestDouble(T *testing.T) {
	PW := CreatePacketWriter(0x00)
	PW.WriteDouble(20.0)
	P := PW.GetPacket()
	fmt.Println(P)
	PR := CreatePacketReader(P)
	PS, _, err := PR.ReadVarInt()
	if err != nil {
		panic(err)
	}
	fmt.Println("PS: ", PS)
	PID, _, err := PR.ReadVarInt()
	if err != nil {
		T.Error(err)
	}
	fmt.Println("PID: ", PID)
	NUM, err := PR.ReadDouble()
	if err != nil {
		T.Error(err)
	}
	fmt.Println("NUM: ", NUM)
	if NUM != 20 {
		T.Error(errWriterValue)
	}
}

func CauseEOF(T *testing.T) {
	PW := CreatePacketWriter(0x00)
	PW.WriteBoolean(false)
	PR := CreatePacketReader(PW.GetPacket())
	PS, _, err := PR.ReadVarInt()
	if err != nil {
		panic(err)
	}
	fmt.Println("PS: ", PS)
	PID, _, err := PR.ReadVarInt()
	if err != nil {
		T.Error(err)
	}
	fmt.Println("PID: ", PID)
	B, err := PR.ReadUnsignedByte()
	if err != nil {
		T.Error(err)
	}
	B, err = PR.ReadUnsignedByte()
	if err == nil {
		T.Error(errWriterExpectedError)
	}
	B, err = PR.ReadUnsignedByte()
	if err == nil {
		T.Error(errWriterExpectedError)
	}
	fmt.Print(B)

}
