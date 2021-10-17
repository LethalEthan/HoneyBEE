package npacket

import "errors"

var (
	errWriterEmpty  = errors.New("Writer Empty Data!")
	errWriterLength = errors.New("Writer Length is incorrect!")
	errWriterType   = errors.New("Writer Type is incorrect!")
	errWriterName   = errors.New("Writer Name incorrect!")
	errWriterValue  = errors.New("Writer Value is incorrect!")
	//
	errWriterReadOverflow = errors.New("Writer reading overflow")
	errWriterRead404      = errors.New("Writer not found!")
	//
	errWriterUnknownType   = errors.New("Value type is unknown, Write :(")
	errWriternil           = errors.New("You tried passing nil, what are you thinking?")
	errWriterExpectedError = errors.New("Expected error, recieved nil")
)
