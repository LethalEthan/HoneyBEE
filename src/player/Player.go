package player

//Information on player
type PlayerE struct {
	Name     string
	UUID     string
	EntityID uint32
	GameMode uint8
}

var PlayerEMap = make(map[uint32]*PlayerE) //PlayerEntityMap

//InitPlayer - Create Player Object
func InitPlayer(Name string, UUID string, EntityID uint32, GameMode uint8) *PlayerE {
	//Overflow protect - in the unlikely event that the assigned number of Player EID's is too big
	if EntityID > 4294967294 {
		log.Critical("Number of entity ID's assigned is over the uint32 limit, PANIC")
		panic("Number of entity ID's assigned is over the uint32 limit, PANIC")
	} else { //NOF
		if val, tmp := PlayerEMap[EntityID]; tmp { //If PlayerEMap returns a value
			P := val //Set P to pre-existing value
			return P
		} else { //Create Player
			P := new(PlayerE)
			P = &PlayerE{Name, UUID, EntityID, GameMode}
			PlayerEMap[P.EntityID] = P //Add to map
			return P
		}
	}
}

func GetPlayer(EID uint32) *PlayerE {
	P := PlayerEMap[EID]
	log.Info("Playermap returned: ", P)
	return P
}
