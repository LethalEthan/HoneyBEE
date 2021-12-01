package utils

import (
	"encoding/binary"
	"errors"
	"reflect"
	"unsafe"
)

var Ascii = `
     ___           ___           ___           ___           ___            ___           ___           ___    
    /\__\         /\  \         /\__\         /\  \         |\__\          /\  \         /\  \         /\  \   
   /:/  /        /::\  \       /::|  |       /::\  \        |:|  |        /::\  \       /::\  \       /::\  \  
  /:/__/        /:/\:\  \     /:|:|  |      /:/\:\  \       |:|  |       /:/\:\  \     /:/\:\  \     /:/\:\  \ 
 /::\  \ ___   /:/  \:\  \   /:/|:|  |__   /::\~\:\  \      |:|__|__    /::\ \:\__\   /::\~\:\  \   /::\~\:\  \ 
/:/\:\  /\__\ /:/__/ \:\__\ /:/ |:| /\__\ /:/\:\ \:\__\     /::::\__\  /:/\:\ \:|__| /:/\:\ \:\__\ /:/\:\ \:\__\
\/__\:\/:/  / \:\  \ /:/  / \/__|:|/:/  / \:\~\:\ \/__/    /:/~~/|__/  \:\ \:\/:/  / \:\~\:\ \/__/ \:\~\:\ \/__/
     \::/  /   \:\  /:/  /      |:/:/  /   \:\ \:\__\     /:/  /        \:\ \::/  /   \:\ \:\__\    \:\ \:\__\  
     /:/  /     \:\/:/  /       |::/  /     \:\ \/__/    /:/  /          \:\/:/  /     \:\ \/__/     \:\ \/__/  
    /:/  /       \::/  /        /:/  /       \:\__\      \/__/            \::/__/       \:\__\        \:\__\   
    \/__/         \/__/         \/__/         \/__/                                      \/__/         \/__/ `

/*
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
    \/__/         \/__/         \/__/         \/__/                       \/__/         \/__/    `*/

var Ascii2 = `
   _
  /|\
  \|/ //
-(||)(')
  '''`

var errbytearrayconversion = errors.New("error converting byte array to desired type")

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

func Uint16ToByteArray(val uint16) []byte {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, val)
	return buf
}

func Uint32ToByteArray(val uint32) []byte {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, val)
	return buf
}

func Uint64ToByteArray(val uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, val)
	return buf
}

func ByteArrayToInt16(val []byte) (int16, error) {
	if len(val) > 2 || len(val) < 2 {
		return 0, errbytearrayconversion
	} else {
		conval := int16(binary.BigEndian.Uint16(val))
		return conval, nil
	}
}

func ByteArrayToInt32(val []byte) (int32, error) {
	if len(val) > 4 || len(val) < 4 {
		return 0, errbytearrayconversion
	}
	conval := int32(binary.BigEndian.Uint32(val))
	return conval, nil
}

func ByteArrayToInt64(val []byte) (int64, error) {
	if len(val) > 8 || len(val) < 8 {
		return 0, errbytearrayconversion
	} else {
		conval := int64(binary.BigEndian.Uint64(val))
		return conval, nil
	}
}

// Unsafe :)

func UnsafeCastInt16ToBytes(val int16) []byte {
	hdr := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(&val)), Len: 2, Cap: 2}
	return *(*[]byte)(unsafe.Pointer(&hdr))
}

func UnsafeCastInt16ArrayToBytes(ints []int16) []byte {
	length := len(ints) * 2
	hdr := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(&ints[0])), Len: length, Cap: length}
	return *(*[]byte)(unsafe.Pointer(&hdr))
}

func UnsafeCastInt32ToBytes(val int32) []byte {
	hdr := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(&val)), Len: 4, Cap: 4}
	return *(*[]byte)(unsafe.Pointer(&hdr))
}

func UnsafeCastInt32ArrayToBytes(ints []int32) []byte {
	length := len(ints) * 4
	hdr := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(&ints[0])), Len: length, Cap: length}
	return *(*[]byte)(unsafe.Pointer(&hdr))
}

func UnsafeCastInt64ToBytes(val int64) []byte {
	hdr := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(&val)), Len: 8, Cap: 8}
	return *(*[]byte)(unsafe.Pointer(&hdr))
}

func UnsafeCastInt64ArrayToBytes(ints []int64) []byte {
	length := len(ints) * 8
	hdr := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(&ints[0])), Len: length, Cap: length}
	return *(*[]byte)(unsafe.Pointer(&hdr))
}

//uints

func UnsafeCastUint16ToBytes(val uint16) []byte {
	hdr := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(&val)), Len: 2, Cap: 2}
	return *(*[]byte)(unsafe.Pointer(&hdr))
}

func UnsafeCastUint16ArrayToBytes(ints []uint16) []byte {
	length := len(ints) * 2
	hdr := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(&ints[0])), Len: length, Cap: length}
	return *(*[]byte)(unsafe.Pointer(&hdr))
}

func UnsafeCastUint32ToBytes(val uint32) []byte {
	hdr := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(&val)), Len: 4, Cap: 4}
	return *(*[]byte)(unsafe.Pointer(&hdr))
}

func UnsafeCastUint32ArrayToBytes(ints []uint32) []byte {
	length := len(ints) * 4
	hdr := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(&ints[0])), Len: length, Cap: length}
	return *(*[]byte)(unsafe.Pointer(&hdr))
}

func UnsafeCastUint64ToBytes(val uint64) []byte {
	hdr := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(&val)), Len: 8, Cap: 8}
	return *(*[]byte)(unsafe.Pointer(&hdr))
}

func UnsafeCastUint64ArrayToBytes(ints []uint64) []byte {
	length := len(ints) * 8
	hdr := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(&ints[0])), Len: length, Cap: length}
	return *(*[]byte)(unsafe.Pointer(&hdr))
}
