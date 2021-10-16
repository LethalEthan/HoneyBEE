package npacket

import "errors"

var (
	writerEmpty  = errors.New("Writer Empty Data!")
	writerLength = errors.New("Writer Length is incorrect!")
	writerType   = errors.New("Writer Type is incorrect!")
	writerName   = errors.New("Writer Name incorrect!")
	writerValue  = errors.New("Writer Value is incorrect!")
	//
	writerReadOverflow = errors.New("Writer reading overflow")
	writerRead404      = errors.New("Writer not found!")
	//
	writerUnknownType   = errors.New("Value type is unknown, Write :(")
	writernil           = errors.New("You tried passing nil, what are you thinking?")
	writerExpectedError = errors.New("Expected error, recieved nil")
)
