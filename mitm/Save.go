package mitm

import (
	"os"
	"strconv"
)

func SaveClient() {
	var name string
	var file *os.File
	var err error
	for i, v := range ClientPackets {
		name = "./CLP/Packet" + strconv.Itoa(i)
		file, err = os.Create(name)
		if err != nil {
			panic(err)
		}
		file.Write(v)
	}
	file.Close()
}

func SaveServer() {
	var name string
	var file *os.File
	var err error
	for i, v := range ServerPackets {
		name = "./SP/Packet" + strconv.Itoa(i)
		file, err = os.Create(name)
		if err != nil {
			panic(err)
		}
		file.Write(v)
	}
	file.Close()
}
