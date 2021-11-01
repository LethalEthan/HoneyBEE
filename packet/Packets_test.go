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
		panic(err)
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
		panic(err)
	}
	Test2, err := Test.MarshalText()
	if err != nil {
		panic(err)
	}
	fmt.Println(Test2)
	Test3, err := PR.ReadString()
	if err != nil {
		panic(err)
	}
	fmt.Println(Test3)

}
