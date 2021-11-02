package packet

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestLoginSuccess(T *testing.T) {
	LS := new(Login_0x02_CB)
	var PW *PacketWriter
	var err error
	LS.UUID, err = uuid.Parse("4ed0c55d-caa9-4669-8bea-f3a0052e6f1f")
	if err != nil {
		T.Error(err)
	}
	LS.Username = "LethalEthan8"
	PW = LS.Encode()
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
		T.Error(errWriterValue)
	}
	Test3, err := PR.ReadString()
	if err != nil {
		T.Error(err)
	}
	if Test3 != "LethalEthan8" {
		T.Error(errWriterValue)
	}
}

func TestFaultyLoginSuccess(T *testing.T) {
	LS := new(Login_0x02_CB)
	var PW *PacketWriter
	var err error
	LS.UUID, err = uuid.Parse("4e342333534d0c55d-caa9345435563523-4669-8bea-f3a0052e6f1f")
	if LS.UUID != uuid.Nil {
		T.Error(errWriterExpectedError)
	}
	LS.Username = "LethalEthan8"
	PW = LS.Encode()
	PR := CreatePacketReader(PW.GetPacket())
	_, _, _ = PR.ReadVarInt()
	_, _, _ = PR.ReadVarInt()
	Test, err := PR.ReadUUID()
	if err == nil && Test != uuid.Nil {
		T.Error(errWriterExpectedError)
		T.Error(errWriterValue)
	}
	Test2, err := Test.MarshalText()
	if err == nil && string(Test2) != "00000000-0000-0000-0000-000000000000" {
		T.Error(errWriterExpectedError)
	}
}
