package config

import (
	"os"
	"gopkg.in/yaml.v3"
	"mp1-server/logger"
)

/*
	Port: port on which server process runs
	AddressPath: filename and path of the membership list file
	TimeOut: time the server thread waits for the client thread to send a request
	LogPath: path of the log files
*/
type Config struct{
	Port			string		`yaml:"port"`
	AddressPath		string		`yaml:"addresspath"`
	TimeOut			int			`yaml:"timeout"`
	LogPath			string		`yaml:"logpath"`
}

/*
	unmarshalls the config.yaml file to initialize the given variables
	returns the Config struct
*/
func NewConfig(logger *logger.CustomLogger) *Config{

	yamlBytes, err := os.ReadFile("config/config.yaml")
	if err!= nil {
		logger.Error("Some error while reading config file",err)
		return nil
	}

	config := &Config{}
	err = yaml.Unmarshal(yamlBytes,&config)
	if err!= nil {
		logger.Error("Some error while reading config file",err)
		return nil
	}

	return config
}