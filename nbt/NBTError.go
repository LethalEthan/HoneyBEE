package nbt

import "errors"

var (
	NBTEmpty  = errors.New("NBT Empty Data!")
	NBTLength = errors.New("NBT Length is incorrect!")
	NBTType   = errors.New("NBT Type is incorrect!")
	NBTName   = errors.New("NBT Name incorrect!")
	NBTValue  = errors.New("NBT Value is incorrect!")
	//
	NBTReadOverflow = errors.New("NBT reading overflow")
	NBTRead404      = errors.New("NBT not found!")
	//
	NBTUnknownType   = errors.New("Value type is unknown, cannot create an NBT tag :(")
	NBTEmptyName     = errors.New("NBT Name is empty!")
	NBTnil           = errors.New("You tried passing nil, what are you thinkng?")
	NBTExpectedError = errors.New("Expected error, recieved nil")
)
