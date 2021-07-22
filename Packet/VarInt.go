package Packet

import (
	"fmt"
	"net"
)

func ParseVarIntFromConnection(conn net.Conn) (int32, error) {
	var result int32
	var numRead uint32
	buff := make([]byte, 1)
	for {
		_, err := conn.Read(buff)
		if err != nil {
			return 0, err
		}
		val := int32((buff[0] & 0x7F))
		result |= (val << (7 * numRead))

		numRead++

		if numRead > 5 {
			return 0, fmt.Errorf("varint was over five bytes without termination")
		}

		if buff[0]&0x80 == 0 {
			break
		}
	}
	return result, nil
}

func ParseVarLongFromConnection(conn net.Conn) (int64, error) {
	var result int64
	var numRead uint64
	buff := make([]byte, 1)
	for {
		_, err := conn.Read(buff)
		if err != nil {
			return 0, err
		}
		val := int64((buff[0] & 0x7F))
		result |= (val << (7 * numRead))

		numRead++

		if numRead > 10 {
			return 0, fmt.Errorf("varlong was over ten bytes without termination")
		}

		if buff[0]&0x80 == 0 {
			break
		}
	}
	return result, nil
}

//ParseBoolFromConnection - Parses true/false data from connection
func ParseBoolFromConnection(conn net.Conn) (bool, error) {
	//var result bool
	var numRead bool
	numRead = false
	//error := 0
	return numRead, nil
}
