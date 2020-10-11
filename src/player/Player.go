package player

import (
	config "config"
	"runtime"
	"sync"
	"time"

	logging "github.com/op/go-logging"
)

var log = logging.MustGetLogger("HoneyGO")

var SConfig = config.GetConfig()

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
	//PlayerObjectMutex - needed in the event of concurrent access
	PlayerObjectMutex = &sync.RWMutex{}
	//PlayerEntityMap - Name/EID
	PlayerEntityMap = make(map[string]uint32)
	//PlayerEntityMutex - needed in the event of concurrent access
	PlayerEntityMutex = &sync.RWMutex{}
	//OnlinePlayerMap - Name/bool
	OnlinePlayerMap = make(map[string]bool)
	//OnlinePlayerMutex - needed in the event of concurrent access
	OnlinePlayerMutex = &sync.RWMutex{}
	PlayerCount       uint64
	GCInterval        time.Duration
)

func Init() {
	SConfig = config.GetConfig()
}

//InitPlayer - Create Player Object
func InitPlayer(Name string, UUID string, GameMode uint8) (*PlayerObject, error) {
	if val, tmp := GetPEMSafe(Name); tmp { //If PlayerEntityMap returns a value
		P, _ := GetPOMSafe(val) //PlayerObjectMap[val] //Set P to pre-existing value - Saves time and reuses previous EntityID
		P.Online = true
		SetOPMSafe(P.Name, P.Online) //OnlinePlayerMap[P.Name] = P.Online
		PlayerCount++
		return P, nil
	} else { //Create Player
		EntityID := AssignEID(Name)
		P := &PlayerObject{Name, UUID, EntityID, GameMode, true}
		SetPOMSafe(P.EntityID, P)    //PlayerObjectMap[P.EntityID] = P //Add EID/OBJ
		SetPEMSafe(P.Name, EntityID) //PlayerEntityMap[P.Name] = EntityID //Add Name/EID
		log.Warning("PlayerOBJ:", P)
		PlayerCount++
		return P, nil
	}
}

func PlayerJoin() {
	PlayerCount++
}

//GetPlayerByID - Gets PlayerObject from map by ID
func GetPlayerByID(EID uint32) *PlayerObject {
	P, _ := GetPOMSafe(EID) //P := PlayerObjectMap[EID]
	log.Info("Playermap returned: ", P)
	return P
}

//GetPlayerByName - Gets PlayerObject from map by Name
func GetPlayerByName(Name string) *PlayerObject {
	P, _ := GetPEMSafe(Name) //PlayerEntityMap[Name]
	log.Info("PlayerEntityMap returned: ", P)
	PO, _ := GetPOMSafe(P) //PlayerObjectMap[P]
	return PO
}

//GCPlayer - Garbage Collect offline and expired players
func GCPlayer(GCP chan bool) {
	if SConfig.Performance.GCPlayer == 0 { //Nothing Set
		GCInterval = 15 * time.Minute //Default to 15 minutes
	} else {
		GCInterval = time.Duration(SConfig.Performance.GCPlayer) * time.Minute
	}
	//CreateTicker
	ticker := time.NewTicker(GCInterval)
	go func() {
		for {
			select {
			case <-ticker.C:
				PlayerObjectMutex.RLock() //Lock map before loop reads
				for i, val := range PlayerObjectMap {
					PlayerObjectMutex.RUnlock() //Unlock map to not tie up others
					if val.Online != true {
						PlayerEntityMutex.Lock() //relock to delete from map
						delete(PlayerEntityMap, val.Name)
						PlayerEntityMutex.Unlock()
						PlayerObjectMutex.Lock()
						delete(PlayerObjectMap, i)
						PlayerObjectMutex.Unlock()
						log.Debug("Player:", val.Name, "deleted from map")
					} else {
						log.Debug("No player to GC found in map")
					}
					PlayerObjectMutex.RLock() //Relock for when loop reads
				}
				PlayerObjectMutex.RUnlock() //make sure that map is unlocked
				runtime.GC()                //run a GC
			case <-GCP:
				if <-GCP {
					Log := logging.MustGetLogger("HoneyGO")
					Log.Warning("Stopping GCPlayer")
					ticker.Stop()
					//Cleanup
					ticker = nil
					Log = nil
					return
				}
			}
		}
	}()
}

//Disconnect - Handles player disconnecting
func Disconnect(Name string) {
	P := GetPlayerByName(Name)
	P.Online = false
	OnlinePlayerMutex.Lock()
	delete(OnlinePlayerMap, P.Name)
	OnlinePlayerMutex.Unlock()
	PlayerCount--
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
	for i = 2; i <= 4294967294; i++ { //starts at 2 atm, I've heard of a client bug with entity ID 0 though I'm not sure
		PlayerObjectMutex.RLock()
		_, tmp := PlayerObjectMap[i] //check current objects
		PlayerObjectMutex.RUnlock()
		if tmp {
			log.Debug("ID:", i, " being used, skip")
		} else { //if ID isn't being used
			C <- i
			return
		}
	}
}

//GetPOMSafe - get value from PlayerObjectMap safely
func GetPOMSafe(key uint32) (*PlayerObject, bool) {
	PlayerObjectMutex.RLock()
	P, B := PlayerObjectMap[key]
	PlayerObjectMutex.RUnlock()
	return P, B
}

//SetPOMSafe - set value in PlayerObjectMap safely
func SetPOMSafe(EID uint32, playerobj *PlayerObject) {
	PlayerObjectMutex.Lock()
	PlayerObjectMap[EID] = playerobj
	PlayerObjectMutex.Unlock()
}

//GetOPMSafe - get value from OnlinePlayerMap safely
func GetOPMSafe(key string) (bool, bool) {
	OnlinePlayerMutex.RLock()
	P, B := OnlinePlayerMap[key]
	OnlinePlayerMutex.RUnlock()
	return P, B
}

//SetOPMSafe - Set value in OnlinePlayerMap safely
func SetOPMSafe(player string, status bool) {
	OnlinePlayerMutex.Lock()
	OnlinePlayerMap[player] = status
	OnlinePlayerMutex.Unlock()
}

func GetPEMSafe(key string) (uint32, bool) {
	PlayerEntityMutex.RLock()
	P, B := PlayerEntityMap[key]
	PlayerEntityMutex.RUnlock()
	return P, B
}

func SetPEMSafe(key string, value uint32) {
	PlayerEntityMutex.Lock()
	PlayerEntityMap[key] = value
	PlayerEntityMutex.Unlock()
}
