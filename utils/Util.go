package utils

import (
	"encoding/binary"
	"errors"
)

var Ascii = `
     ___           ___           ___           ___           ___           ___           ___
    /\__\         /\  \         /\__\         /\  \         |\__\         /\  \         /\  \
   /:/  /        /::\  \       /::|  |       /::\  \        |:|  |       /::\  \       /::\  \
  /:/__/        /:/\:\  \     /:|:|  |      /:/\:\  \       |:|  |      /:/\:\  \     /:/\:\  \
 /::\  \ ___   /:/  \:\  \   /:/|:|  |__   /::\~\:\  \      |:|__|__   /:/  \:\  \   /:/  \:\  \
/:/\:\  /\__\ /:/__/ \:\__\ /:/ |:| /\__\ /:/\:\ \:\__\     /::::\__\ /:/__/_\:\__\ /:/__/ \:\__\
\/__\:\/:/  / \:\  \ /:/  / \/__|:|/:/  / \:\~\:\ \/__/    /:/~~/|__/ \:\  /\ \/__/ \:\  \ /:/  /
     \::/  /   \:\  /:/  /      |:/:/  /   \:\ \:\__\     /:/  /       \:\ \:\__\    \:\  /:/  /
     /:/  /     \:\/:/  /       |::/  /     \:\ \/__/    /:/  /         \:\/:/  /     \:\/:/  /
    /:/  /       \::/  /        /:/  /       \:\__\      \/__/           \::/  /       \::/  /
    \/__/         \/__/         \/__/         \/__/                       \/__/         \/__/    `

var Ascii2 = `
   _
  /|\
  \|/ //
-(||)(')
  '''`

var BACONV = errors.New("Error converting byte array to desired type")

func Int16ToByteArray(val int16) []byte {
	Bint16 := []byte{byte(val >> 8), byte(val)}
	return Bint16
}

func Int32ToByteArray(val int32) []byte {
	Bint32 := []byte{byte(val >> 24), byte(val >> 16), byte(val >> 8), byte(val)}
	return Bint32
}

func Int64ToByteArray(val int64) []byte {
	Bint64 := []byte{byte(val >> 56), byte(val >> 48), byte(val >> 40), byte(val >> 32), byte(val >> 24), byte(val >> 16), byte(val >> 8), byte(val)}
	return Bint64
}

func ByteArrayToInt16(val []byte) (int16, error) {
	if len(val) > 2 {
		return 0, BACONV
	} else {
		var conval int16
		conval = int16(binary.BigEndian.Uint16(val))
		return conval, nil
	}
}

func ByteArrayToInt32(val []byte) (int32, error) {
	if len(val) > 4 {
		return 0, BACONV
	}
	var conval int32
	conval = int32(binary.BigEndian.Uint32(val))
	return conval, nil
}

func ByteArrayToInt64(val []byte) (int64, error) {
	if len(val) > 8 {
		return 0, BACONV
	} else {
		var conval int64
		conval = int64(binary.BigEndian.Uint64(val))
		return conval, nil
	}
}
