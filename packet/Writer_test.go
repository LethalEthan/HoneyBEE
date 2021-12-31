package packet

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestCreatePacketWriter(T *testing.T) {
	PW := CreatePacketWriter(0x00)
	if cap(PW.data) != 2048 {
		T.Error(WriterLength)
	}
}

func TestCreatePacketWriterWithCapacity(T *testing.T) {
	PW := CreatePacketWriterWithCapacity(0x00, 256)
	if cap(PW.data) != 256 {
		T.Error(WriterLength)
	}
}

func TestCreatePacketWriterWithCapacityBelow0(T *testing.T) {
	PW := CreatePacketWriterWithCapacity(0x00, -45345)
	if cap(PW.data) != 2048 {
		T.Error(WriterLength)
	}
}

func TestWriteByte(T *testing.T) {
	PW := CreatePacketWriter(0x00)
	PW.WriteByte(8)
	if len(PW.data) != 2 {
		T.Error(WriterLength)
		T.Error(PW.data)
	} else {
		if PW.data[1] != 8 {
			T.Error(WriterValue)
		}
	}
}
func TestWriteBoolean(T *testing.T) {
	PW := CreatePacketWriter(0x00)
	PW.WriteBoolean(true)
	PW.WriteBoolean(false)
	if PW.data[1] != 1 {
		T.Error(WriterValue)
	}
	if PW.data[2] != 0 {
		T.Error(WriterValue)
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
		T.Error(WriterValue)
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
		T.Error(WriterValue)
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
	var B byte
	_, err = PR.ReadUByte()
	if err != nil {
		T.Error(err)
	}
	_, err = PR.ReadUByte()
	if err == nil {
		T.Error(WriterExpectedError)
	}
	B, err = PR.ReadUByte()
	if err == nil {
		T.Error(WriterExpectedError)
	}
	fmt.Print(B)
}

func WriteShort(T *testing.T) {
	PW := CreatePacketWriter(0x00)
	PW.WriteShort(32767)
	PR := CreatePacketReader(PW.GetPacket())
	_, _, err := PR.ReadVarInt()
	if err != nil {
		T.Errorf("err: %s and %d", WriterValue, err)
	}
	_, _, err = PR.ReadVarInt()
	if err != nil {
		T.Errorf("err: %s and %d", WriterValue, err)
	}
	NUM, err := PR.ReadDouble()
	if err != nil {
		T.Errorf("err: %s and %d", WriterValue, err)
	}
	if NUM != 32767 {
		T.Errorf("err: %s and %d", WriterValue, err)
	}
}

func TestAllRWValues(T *testing.T) {
	PW := CreateWriterWithCapacity(1024)
	_ = CreateWriterWithCapacity(-129)
	PW.WriteVarLong(69)
	PW.WriteVarInt(420)
	PW.WriteUByte(1)
	PW.WriteByte(127)
	PW.WriteBoolean(true)
	PW.WriteUShort(65535)
	PW.WriteShort(32767)
	PW.WriteInt(400000000)
	PW.WriteUInt(34)
	PW.WriteFloat(69.420)
	PW.WriteLong(9223372036854775807)
	PW.WriteULong(18446744073709551615)
	PW.WritePosition(1000, 1000, 1000)
	PW.WriteDouble(42069.42069)
	PW.WriteArray([]byte{0, 255, 100, 100, 100, 50, 40, 10, 0, 255, 100, 100, 100, 50, 40, 10, 0, 255, 100, 100, 100, 50, 40, 10, 0, 255, 100, 100, 100, 50, 40, 10})
	PW.WriteLongArray([]int64{0, 1, 2, 4, 7, 7, 5, 346436635, 3, 6, 8, 9, 5, 3, 45, 5})
	PW.WriteString("HoneyBEE!")
	PW.WriteIdentifier("HoneyBEE:pre-alpha")
	PW.WriteArrayIdentifier([]Identifier{"HoneyBEE:blocks", "HoneyBEE:items", "HoneyBEE:helpus"})
	PW.WriteUUID(uuid.MustParse("7935dab6-de0e-3b13-8716-2b4ab8e5f806"))
	// fmt.Println("ALL DATA: ", PW.GetData())
	PR := CreatePacketReader(PW.GetData())
	if VL, err := PR.ReadVarLong(); err != nil || VL != 69 {
		T.Error(ReaderValue)
	}
	if VI, _, err := PR.ReadVarInt(); err != nil || VI != 420 {
		T.Error(ReaderValue)
	}
	if UB, err := PR.ReadUByte(); err != nil || UB != 1 {
		T.Error(ReaderValue)
	}
	if B, err := PR.ReadByte(); err != nil || B != 127 {
		T.Error(ReaderValue)
	}
	if TF, err := PR.ReadBoolean(); err != nil || !TF {
		T.Error(ReaderValue)
	}
	if US, err := PR.ReadUShort(); err != nil || US != 65535 {
		T.Error(ReaderValue)
	}
	if S, err := PR.ReadShort(); err != nil || S != 32767 {
		T.Error(ReaderValue)
	}
	if I, err := PR.ReadInt(); err != nil || I != 400000000 {
		T.Error(ReaderValue)
	}
	if UI, err := PR.ReadUInt(); err != nil || UI != 34 {
		T.Error(ReaderValue)
	}
	if F, err := PR.ReadFloat(); err != nil || F != 69.420 {
		T.Error(ReaderValue)
	}
	if L, err := PR.ReadLong(); err != nil || L != 9223372036854775807 {
		T.Error(ReaderValue)
	}
	if UL, err := PR.ReadULong(); err != nil || UL != 18446744073709551615 {
		T.Error(ReaderValue)
	}
	if X, Y, Z, err := PR.ReadPosition(); err != nil || X != 1000 || Y != 1000 || Z != 1000 {
		T.Error(ReaderValue)
	}
	if D, err := PR.ReadDouble(); err != nil || D != 42069.42069 {
		T.Error(ReaderValue)
	}
	TMP := []byte{0, 255, 100, 100, 100, 50, 40, 10, 0, 255, 100, 100, 100, 50, 40, 10, 0, 255, 100, 100, 100, 50, 40, 10, 0, 255, 100, 100, 100, 50, 40, 10}
	A, err := PR.ReadByteArray(len(TMP))
	if err != nil {
		T.Error()
	} else {
		for i := range A {
			if A[i] != TMP[i] {
				T.Error(ReaderValue)
			}
		}
	}
	// TMP2 := []int64{0, 1, 2, 4, 7, 7, 5, 346436635, 3, 6, 8, 9, 5, 3, 45, 5}
	// LA, err := PR.ReadLongArray(len(TMP2))
	// if err != nil {
	// 	T.Error(ReaderValue)
	// } else {
	// 	for i := range A {
	// 		if LA[i] != TMP2[i] {
	// 			T.Error(ReaderValue)
	// 		}
	// 	}
	// }
	//
	PW.GetPacketID()
	PW.GetPacketSize()
	PW.ClearData()
	PR.SetData([]byte{0})
}

func BenchmarkBlockPositionEncode(B *testing.B) {
	B.ReportAllocs()
	B.ResetTimer()
	for i := 0; i < B.N; i++ {
		Block := i<<12 | (i<<8 | i<<4 | i)
		_ = Block
	}
}
