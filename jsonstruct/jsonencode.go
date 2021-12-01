package jsonstruct

func (C ChatComponent) MarshalChatComponent() []byte {
	B, err := C.MarshalJSON()
	if err != nil {
		panic(err)
	}
	return B
}
