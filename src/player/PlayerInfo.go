package Player

type VarInt int32

type playerinfo struct {
	Action       VarInt
	PlayerNumber VarInt
	//Player []{UUID,}
}

// type Action struct {
//   AddPlayer
// }
