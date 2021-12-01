package player

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	logging "github.com/op/go-logging"
)

var log = logging.MustGetLogger("HoneyBEE")

//Information on player
type PlayerObject struct {
	PlayerName   string
	Locale       string
	ViewDistance byte
	ChatMode     byte
	UUID         uuid.UUID
}

var (
	//PlayerObjectMap - EID/PlayerObject
	PlayerObjectMap = make(map[uint32]*PlayerObject)
	//PlayerObjectMutex - needed in the event of concurrent access
	PlayerObjectMutex = sync.RWMutex{}
	//PlayerEntityMap - Name/EID
	PlayerEntityMap = make(map[string]uint32)
	//PlayerEntityMutex - needed in the event of concurrent access
	PlayerEntityMutex = sync.RWMutex{}
	//OnlinePlayerMap - Name/bool
	//OnlinePlayerMap = make(map[string]bool)
	//OnlinePlayerMutex - needed in the event of concurrent access
	//OnlinePlayerMutex = sync.RWMutex{}
	PlayerCount   uint64
	GCInterval    time.Duration
	GlobalIDCount uint32
)

/*InitPlayer - Create Player Object
func InitPlayer(Name string, UUID string, GameMode uint8) (*PlayerObject, error) {
	if val, tmp := GetPEM(Name); tmp { //If PlayerEntityMap returns a value
		P, _ := GetPOM(val) //PlayerObjectMap[val] //Set P to pre-existing value - Saves time and reuses previous EntityID
		P.Online = true
		SetOPM(P.Name, P.Online) //OnlinePlayerMap[P.Name] = P.Online
		PlayerCount++
		return P, nil
	} else { //Create Player
		EntityID := AssignEID(Name)
		P := &PlayerObject{Name, UUID, EntityID, GameMode, true}
		SetPOM(P.EntityID, P)    //PlayerObjectMap[P.EntityID] = P //Add EID/OBJ
		SetPEM(P.Name, EntityID) //PlayerEntityMap[P.Name] = EntityID //Add Name/EID
		log.Warning("PlayerOBJ:", P)
		PlayerCount++
		return P, nil
	}
}

//GetPlayerByID - Gets PlayerObject from map by ID
func GetPlayerByID(EID uint32) *PlayerObject {
	P, _ := GetPOM(EID) //P := PlayerObjectMap[EID]
	log.Info("Playermap returned: ", P)
	return P
}

//GetPlayerByName - Gets PlayerObject from map by Name
func GetPlayerByName(Name string) *PlayerObject {
	P, _ := GetPEM(Name) //PlayerEntityMap[Name]
	log.Info("PlayerEntityMap returned: ", P)
	PO, _ := GetPOM(P) //PlayerObjectMap[P]
	return PO
}

//Disconnect - Handles player disconnecting
func Disconnect(Name string) {
	P := GetPlayerByName(Name)
	P.Online = false
	OnlinePlayerMutex.Lock()
	delete(OnlinePlayerMap, P.Name)
	OnlinePlayerMutex.Unlock()
	PlayerCount--
}*/

func AssignEID(P string) uint32 {
	if val, tmp := PlayerEntityMap[P]; tmp { //Pre-Existing Value
		return val
	} else {
		//Maybe parallelised later by splitting the work load i.e. one go routine searches ID range 1-2000 and another searches 20001-4000
		//This may not be necessary though but is a consideration in case there is a performance impact
		ID, err := FindFreeID()
		if err != nil {
			fmt.Println("GlobalIDCount: ", GlobalIDCount)
			panic(err)
		}
		log.Debug("IIL: ", ID)
		return ID
	}
}

//FindFreeID - Finds a free ID to assign for players
func FindFreeID() (uint32, error) {
	for GlobalIDCount = 2; GlobalIDCount <= 4294967294; GlobalIDCount++ { //starts at 2 atm, I've heard of a client bug with entity ID 0 though I'm not sure
		PlayerObjectMutex.RLock()
		_, tmp := PlayerObjectMap[GlobalIDCount] //check current objects
		PlayerObjectMutex.RUnlock()
		if tmp {
			log.Debug("ID:", GlobalIDCount, " being used, skip")
		} else { //if ID isn't being used
			return GlobalIDCount, nil
		}
	}
	return 4, errors.New("Could not assign an ID!")
}

//GetPOM - get value from PlayerObjectMap
func GetPOM(key uint32) (*PlayerObject, bool) {
	PlayerObjectMutex.RLock()
	P, B := PlayerObjectMap[key]
	PlayerObjectMutex.RUnlock()
	return P, B
}

//SetPOM - set value in PlayerObjectMap
func SetPOM(EID uint32, playerobj *PlayerObject) {
	PlayerObjectMutex.Lock()
	PlayerObjectMap[EID] = playerobj
	PlayerObjectMutex.Unlock()
}

/*GetOPM - get value from OnlinePlayerMap
func GetOPM(key string) (bool, bool) {
	OnlinePlayerMutex.RLock()
	P, B := OnlinePlayerMap[key]
	OnlinePlayerMutex.RUnlock()
	return P, B
}

//SetOPM - Set value in OnlinePlayerMap
func SetOPM(player string, status bool) {
	OnlinePlayerMutex.Lock()
	OnlinePlayerMap[player] = status
	OnlinePlayerMutex.Unlock()
}*/

func GetPEM(key string) (uint32, bool) {
	PlayerEntityMutex.RLock()
	P, B := PlayerEntityMap[key]
	PlayerEntityMutex.RUnlock()
	return P, B
}

func SetPEM(key string, value uint32) {
	PlayerEntityMutex.Lock()
	PlayerEntityMap[key] = value
	PlayerEntityMutex.Unlock()
}
