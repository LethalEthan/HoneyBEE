package config

import (
	"flag"
	"fmt"
	"os"

	logging "github.com/op/go-logging"
	yaml "gopkg.in/yaml.v3"
)

var (
	log        = logging.MustGetLogger("HoneyBEE")
	GConfig    Config
	Memprofile string
	Cpuprofile string
	configPath string
)

// Config struct for HoneyBEE config
type Config struct {
	Server struct {
		Host          string `yaml:"host"`
		Port          string `yaml:"port"`
		DEBUG         bool   `yaml:"debug"`
		Timeout       int    `yaml:"timeout"`
		MultiCore     bool   `yaml:"multicore"`
		LockOSThread  bool   `yaml:"lock-os-thread"`
		Reuse         bool   `yaml:"SO_REUSE"`
		SendBuf       int    `yaml:"send-buffer"`
		RecieveBuf    int    `yaml:"recv-buffer"`
		ReadBufferCap int    `yaml:"read-buffer-cap"`
		NumEventLoop  int    `yaml:"net-event-loop"`
		Protocol      struct {
			AvailableProtocols  []int32 `yaml:"available-protocols"`
			BlockPlayersOnLogin bool    `yaml:"block-players-on-login"`
		} `yaml:"protocol"`
	} `yaml:"server"`
	Performance struct {
		CPU                int `yaml:"cpu"`
		GCPercent          int `yaml:"gc-percent"`
		ViewDistance       int `yaml:"view-distance"`
		SimulationDistance int `yaml:"simulation-distance"`
	} `yaml:"performance"`
	DEBUGOPTS struct {
		Maintenance bool `yaml:"maintenance"`
		//ServerStatusMessage string `yaml:"server-status-message"`
	} `yaml:"debug-opts"`
}

// NewConfig returns a new decoded Config struct
func NewConfig(configPath string) error {
	//Create config structure
	GConfig = *new(Config)
	//Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	d := yaml.NewDecoder(file) //Create new YAML decode
	//Start YAML decoding from file
	if err := d.Decode(&GConfig); err != nil {
		return err
	}
	file.Close()
	return nil
}

//ValidateConfigPath - makes sure that the path provided is a file that can be read
func ValidateConfigPath(path string) error {
	stat, err := os.Stat(path)
	if err != nil {
		return err
	}
	if stat.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

//ParseFlags - will create and parse the CLI flags and return the path to be used
func ParseFlags() (string, error) {
	//var configPath string
	//Set up a CLI flag "-config" to allow users to supply the configuration file - defaults to config.yml
	flag.StringVar(&configPath, "config", "./config.yml", "path to config file")
	flag.StringVar(&Memprofile, "memprofile", "", "write memory profile to this file")
	flag.StringVar(&Cpuprofile, "cpuprofile", "", "write cpu profile to file")
	//Parse the flags
	flag.Parse()
	//Validate the path
	if err := ValidateConfigPath(configPath); err != nil {
		return "", err
	}
	//Return the configuration path
	return configPath, nil
}

//ConfigStart - Handles the config struct creation
func ConfigStart() error {
	//Create config struct
	cfgPath, err := ParseFlags()
	if err != nil {
		return err
	}
	err = NewConfig(cfgPath)
	if err != nil {
		return err
	}
	if GConfig.Server.DEBUG {
		fmt.Println("cfg: ", GConfig)
	}
	return nil
}

func ConfigReload() {
	err := NewConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}
	if GConfig.Server.DEBUG {
		log.Debug("cfg: ", GConfig)
	}
}
