package player

import (
	config "config"
	"time"

	logging "github.com/op/go-logging"
)

var log = logging.MustGetLogger("HoneyGO")

var SConfig = config.GetConfig()

//import "time"

//Information on player
type PlayerObject struct {
	Name     string
	UUID     string
	EntityID uint32
	GameMode uint8
	//	TOC      time.Time //Time Of Creation, used for GC
	Online bool
}

var (
	//PlayerObjectMap - EID/PlayerObject
	PlayerObjectMap = make(map[uint32]*PlayerObject)
	//PlayerEntityMap - Name/EID
	PlayerEntityMap = make(map[string]uint32)
	//OnlinePlayerMap - Name/bool
	OnlinePlayerMap = make(map[string]bool)
	PlayerCount     uint64
	GCInterval      time.Duration
)

func Init() {
	SConfig = config.GetConfig()
}

//InitPlayer - Create Player Object
func InitPlayer(Name string, UUID string, GameMode uint8) (*PlayerObject, error) {
	if val, tmp := PlayerEntityMap[Name]; tmp { //If PlayerEntityMap returns a value
		P := PlayerObjectMap[val] //Set P to pre-existing value - Saves time and reuses previous EntityID
		P.Online = true
		OnlinePlayerMap[P.Name] = P.Online
		return P, nil
	} else { //Create Player
		//P := new(PlayerObject)
		EntityID := AssignEID(Name)
		P := &PlayerObject{Name, UUID, EntityID, GameMode, true}
		PlayerObjectMap[P.EntityID] = P    //Add EID/OBJ
		PlayerEntityMap[P.Name] = EntityID //Add Name/EID
		log.Warning("PlayerOBJ:", P)
		return P, nil
	}
}

func PlayerJoin() {
	PlayerCount++
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
	if SConfig.Performance.GCPlayer == 0 { //Nothing Set
		GCInterval = 15 * time.Minute //Default to 15 minutes
	} else {
		GCInterval = time.Duration(SConfig.Performance.GCPlayer) * time.Minute
	}

	ticker := time.NewTicker(GCInterval)
	go func() {
		for {
			select {
			case <-ticker.C:
				if PlayerObjectMap == nil {
					log.Info("PlayerObject not initialised")
				}
				for i, val := range PlayerObjectMap {
					if val.Online != true {
						delete(PlayerEntityMap, val.Name)
						delete(PlayerObjectMap, i)
						log.Debug("Player:", val.Name, "deleted from map")
					} else {
						log.Debug("No player to GC found in map")
						break
					}
				}
			}
		}
	}()
}

func TrackPlayerCount(D bool) {

}

//Disconnect - Handles player disconnecting
func Disconnect(Name string) {
	P := GetPlayerByName(Name)
	P.Online = false
	delete(OnlinePlayerMap, P.Name)
	go GCPlayer()
	PlayerCount--
	//delete(PlayerObjectMap, P.EntityID)
}

func AssignEID(P string) uint32 {
	if val, tmp := PlayerEntityMap[P]; tmp { //Pre-Existing Value
		return val
	} else {
		//Maybe parallelised later by splitting the work load i.e. one go routine searches ID range 1-2000 and another searches 20001-4000
		//This may not be necessary though but is a consideration in case there is a performance impact
		C := make(chan uint32)
		go FindFreeID(C)
		ID := <-C
		log.Debug("IIL: ", ID)
		close(C)
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
			}
		case false: //if ID isn't being used
			{
				C <- i
				return
			}
		}
	}
}
