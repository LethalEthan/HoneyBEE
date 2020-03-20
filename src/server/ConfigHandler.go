package server

//var Log *logging.Logger
import (
	"encoding/json"
	"io/ioutil"
)

var defaultserverport = ":25565"

type config struct {
	debug      bool
	serverport string `json:text`
}

func main() {
	data := config{
		debug:      false,
		serverport: ":25565",
	}
	file, _ := json.MarshalIndent(data, "", "")
	_ = ioutil.WriteFile("config.json", file, 0775)
}
func check(e error) {
	if e != nil {
		Log.Info("UhOh")
		panic(e)
	}
}
func testconf() string {
	//debug := "true"
	var serverport string
	serverport = ":25565"
	// c := new(config)
	// c.serverport = config{serverport: ":25565", debug: true}
	//c.serverport = ":25565"
	//return serverport
	return serverport
}
