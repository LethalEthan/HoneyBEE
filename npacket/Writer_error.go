package npacket

import "errors"

var (
	WriterEmpty  = errors.New("Writer Empty Data!")
	WriterLength = errors.New("Writer Length is incorrect!")
	WriterType   = errors.New("Writer Type is incorrect!")
	WriterName   = errors.New("Writer Name incorrect!")
	WriterValue  = errors.New("Writer Value is incorrect!")
	//
	WriterReadOverflow = errors.New("Writer reading overflow")
	WriterRead404      = errors.New("Writer not found!")
	//
	WriterUnknownType   = errors.New("Value type is unknown, Write :(")
	Writernil           = errors.New("You tried passing nil, what are you thinking?")
	WriterExpectedError = errors.New("Expected error, recieved nil")
)
