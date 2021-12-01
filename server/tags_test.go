package server

import (
	"HoneyBEE/packet"
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"
)

func TestTags(T *testing.T) {
	PR := packet.CreatePacketReader(TagsPacket)
	_, _, _ = PR.ReadVarInt()
	_, _, _ = PR.ReadVarInt()
	TA := PR.ReadTagArray()
	Log.Info(TA)
	Log.Info("TagName: ", TA[0].TagArray[0].TagName)
	Log.Info("TagCount: ", TA[0].TagArray[0].Count)
	Log.Info("Entries: ", TA[0].TagArray[0].Entries)
	TN := time.Now()
	file, _ := json.MarshalIndent(TA, "", " ")
	_ = ioutil.WriteFile("test.json", file, 0644)
	Log.Info("Time took to marshal: ", time.Since(TN))
}
