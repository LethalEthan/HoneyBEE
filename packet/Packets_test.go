package packet

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestLoginSuccess(T *testing.T) {
	LS := new(Login_0x02_CB)
	PW := CreatePacketWriter(0x02)
	var err error
	LS.UUID, err = uuid.Parse("4ed0c55d-caa9-4669-8bea-f3a0052e6f1f")
	if err != nil {
		T.Error(err)
	}
	LS.Username = "LethalEthan8"
	LS.Encode(&PW)
	PR := CreatePacketReader(PW.GetPacket())
	PS, _, _ := PR.ReadVarInt()
	fmt.Println("Size:", PS)
	PID, _, _ := PR.ReadVarInt()
	fmt.Println("ID: ", PID)
	Test, err := PR.ReadUUID()
	if err != nil {
		T.Error(err)
	}
	Test2, err := Test.MarshalText()
	if err != nil {
		T.Error(err)
	}
	if string(Test2) != "4ed0c55d-caa9-4669-8bea-f3a0052e6f1f" {
		T.Error(WriterValue)
	}
	Test3, err := PR.ReadString()
	if err != nil {
		T.Error(err)
	}
	if Test3 != "LethalEthan8" {
		T.Error(WriterValue)
	}
}

func TestFaultyLoginSuccess(T *testing.T) {
	LS := new(Login_0x02_CB)
	var PW PacketWriter
	var err error
	LS.UUID, _ = uuid.Parse("4e342333534d0c55d-caa9345435563523-4669-8bea-f3a0052e6f1f")
	if LS.UUID != uuid.Nil {
		T.Error(WriterExpectedError)
	}
	LS.Username = "LethalEthan8"
	if err := LS.Encode(&PW); err != nil {
		panic(err)
	}
	PR := CreatePacketReader(PW.GetPacket())
	_, _, _ = PR.ReadVarInt()
	_, _, _ = PR.ReadVarInt()
	Test, err := PR.ReadUUID()
	if err == nil && Test != uuid.Nil {
		T.Error(WriterExpectedError)
		T.Error(WriterValue)
	}
	Test2, err := Test.MarshalText()
	if err == nil && string(Test2) != "00000000-0000-0000-0000-000000000000" {
		T.Error(WriterExpectedError)
	}
}
