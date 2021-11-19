package packet

const (
	WriterEmpty  = "Writer Empty Data!"
	WriterLength = "Writer Length is incorrect!"
	WriterType   = "Writer Type is incorrect!"
	WriterName   = "Writer Name incorrect!"
	WriterValue  = "Writer Value is incorrect!"
	//
	WriterReadOverflow = "Writer reading overflow"
	WriterRead404      = "Writer not found!"
	//
	WriterUnknownType   = "Value type is unknown, Write :("
	Writernil           = "You tried passing nil, what are you thinking?"
	WriterExpectedError = "Expected error, recieved nil"
)

const (
	ReaderEmpty  = "Reader Empty Data!"
	ReaderLength = "Reader Length is incorrect!"
	ReaderType   = "Reader Type is incorrect!"
	ReaderName   = "Reader Name incorrect!"
	ReaderValue  = "Reader Value is incorrect!"
	//
	ReaderReadOverflow = "Reader reading overflow"
	ReaderRead404      = "Reader not found!"
	//
	ReaderUnknownType   = "Value type is unknown, Write :("
	Readernil           = "You tried passing nil, what are you thinking?"
	ReaderExpectedError = "Expected error, recieved nil"
)
