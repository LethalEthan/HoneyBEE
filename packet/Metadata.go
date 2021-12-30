package packet

type MetadataFormat struct {
	Index byte
	Type  int32  //Opt
	Value []byte //we'll encode the type into bytes for simplicty instead of using empty interface
}
