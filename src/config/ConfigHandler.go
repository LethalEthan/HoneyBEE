package config

import (
	"flag"
	"fmt"
	"os"

	logging "github.com/op/go-logging"
	yaml "gopkg.in/yaml.v3"
)

var (
	log     = logging.MustGetLogger("HoneyGO")
	ConfigR *Config
)

// Config struct for HoneyGO config
type Config struct {
	Server struct {
		Host     string `yaml:"host"`    //IP Address to bind the Server to -- TBD
		Port     string `yaml:"port"`    //TCP Port to bind the Server to
		DEBUG    bool   `yaml:"debug"`   //Output DEBUG info -- TO BE LINKED
		Timeout  int    `yaml:"timeout"` // Server timeout to use until a connection is destroyed when unresponsive (in seconds)
		Protocol struct {
			AvailableProtocols  []int32 `yaml:"available-protocols"`
			BlockPlayersOnLogin bool    `yaml:"block-players-on-login"`
		} `yaml:"protocol"`
	} `yaml:"server"`
	Performance struct {
		CPU            int  `yaml:"cpu"`
		EnableGCPlayer bool `yaml:"enable-gc-player"`
		GCPlayer       int  `yaml:"gc-player-interval"`
		GCPercent      int  `yaml:"gc-percent"`
	} `yaml:"performance"`
	DEBUGOPTS struct {
		PacketAnal        bool   `yaml:"packet-anal"`
		PacketAnalAddress string `yaml:"packet-anal-address"`
	} `yaml:"debug-opts"`
}

// NewConfig returns a new decoded Config struct
func NewConfig(configPath string) (*Config, error) {
	//Create config structure
	config := &Config{}

	//Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	d := yaml.NewDecoder(file) //Create new YAML decode
	//Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
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

var configPath string

//ParseFlags - will create and parse the CLI flags and return the path to be used
func ParseFlags() (string, error) {
	//var configPath string
	//Set up a CLI flag "-config" to allow users to supply the configuration file - defaults to config.yml
	flag.StringVar(&configPath, "config", "./config.yml", "path to config file")
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
func ConfigStart() *Config {
	//Create config struct
	cfgPath, err := ParseFlags()
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := NewConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}
	ConfigR = cfg
	if cfg.Server.DEBUG {
		log.Debug("cfg: ", cfg)
	}
	return cfg
}

func GetConfig() *Config {
	return ConfigR
}

func GetSPort() string {
	return ConfigR.Server.Port
}

func ConfigReload() {
	ConfigR, err := NewConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}
	if ConfigR.Server.DEBUG {
		log.Debug("cfg: ", ConfigR)
	}
}
