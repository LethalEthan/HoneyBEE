package player

import (
	"time"
)

//import "time"

//Information on player
type PlayerObject struct {
	Name     string
	UUID     string
	EntityID uint32
	GameMode uint8
	TOC      time.Time //Time Of Creation, used for GC
	Online   bool
}

var (
	PlayerObjectMap = make(map[uint32]*PlayerObject) //PlayerObjectMap - EID/PlayerObject
	PlayerEntityMap = make(map[string]uint32)        //PlayerEntityMap - Name/EID
)

//InitPlayer - Create Player Object
func InitPlayer(Name string, UUID string/*, EntityID uint32*/, GameMode uint8) (*PlayerObject, error) {
	//Overflow protect - in the unlikely event that the assigned number of Player EID's is too big
	if val, tmp := PlayerEntityMap[Name]; tmp { //If PlayerEntityMap returns a value
		P := PlayerObjectMap[val] //Set P to pre-existing value - Saves time and reuses previous EntityID
		P.Online = true
		return P, nil
	} else { //Create Player
		P := new(PlayerObject)
		EntityID := AssignEID(Name)
		P = &PlayerObject{Name, UUID, EntityID, GameMode, time.Now(), true}
		PlayerObjectMap[P.EntityID] = P    //Add EID/OBJ
		PlayerEntityMap[P.Name] = EntityID //Add Name/EID
		log.Warning("PlayerOBJ:", P)
		return P, nil
	}
}


//GetPlayerByID - Gets PlayerObject from map by ID
func GetPlayerByID(EID uint32) *PlayerObject {
	P := PlayerObjectMap[EID]
	log.Info("Playermap returned: ", P)
	return P
}

//GetPlayerByName - Gets PlayerObject from map by ID
func GetPlayerByName(Name string) *PlayerObject {
	P := PlayerEntityMap[Name]
	log.Info("PlayerEntityMap returned: ", P)
	PO := PlayerObjectMap[P]
	return PO
}

//GCPlayer - Garbage Collect offline and expired players
func GCPlayer() {
	//--//
	L := uint32(len(PlayerObjectMap))    //size of map
	TT, err := time.ParseDuration("20m") //Duration to check for
	if err != nil {
		log.Error("Time didn't work")
	}
	TTL := int64(TT) //Duration -> int64 - simplifies checking
	var i uint32
	//--//
	//--//
	//iterate through map and check if player object should expire
	//if PlayerObjectMap != nil {
	for i = 2; i <= L; i++ {
		val, tmp := PlayerObjectMap[i]
		switch tmp {
		case true:
			{
				if val.Online != true {
					Elapse := int64(time.Since(val.TOC)) //time.Time -> int64 - simplifies checking
					if Elapse > TTL {
						log.Debug("Longer than 15 mins")
						delete(PlayerEntityMap, val.Name)
						delete(PlayerObjectMap, i)
						log.Debug("Player:", val.Name, "deleted from map")
					}
				}
				break
			}
		case false:
			{
				log.Debug("No player to GC found in map")
				break
			}
		}
	}
}

func TrackPlayerCount(D bool) {

}

//Disconnect - Handles player disconnecting
func Disconnect(Name string) {
	P := GetPlayerByName(Name)
	P.Online = false
	//delete(PlayerObjectMap, P.EntityID)
}

func AssignEID(P string) uint32 {
	if val, tmp := PlayerEntityMap[P]; tmp { //Pre-Existing Value
		return val
	} else {
		C := make(chan uint32)
		go FindFreeID(C)
		ID := <-C
		log.Debug("IIL: ", ID)
//		close(C)
		return ID
	}
}

//FindFreeID - Finds a free ID to assign for players
func FindFreeID(C chan uint32) {
	var i uint32
	for i = 2; i <= 4294967294; i++ {
		_, tmp := PlayerObjectMap[i] //check current objects
		switch tmp {
		case true:
			{
				log.Debug("ID:", i, " being used, skip")
				//break
			}
		case false: //if ID isn't being used
			{
				C <- i
				//close(C)
				return
			}
		}
	}
}
/*
func Findy(P string) uint32 {
	var i uint32
	if val, tmp := PlayerEntityMap[P]; tmp { //Pre-Existing non-exValue
		return val
	} else {
		for i = 2; i <= 4294967294; i++ {
			_, tmp := PlayerObjectMap[i] //check current objects
			switch tmp {
			case true:
				{
				log.Debug("ID:", i, " being used, skip")
				//break
				}
			case false: //if ID isn't being used
				{
				return i
				}
			}
		}
	}
	return 90000
}*/
