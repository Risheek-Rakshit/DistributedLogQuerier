package config

import (
	"os"
	"gopkg.in/yaml.v3"
	"mp1-server/logger"
)
type Config struct{
	Port			string		`yaml:"port"`
	AddressPath		string		`yaml:"addresspath"`
	TimeOut			int			`yaml:"timeout"`
	LogPath			string		`yaml:"logpath"`
}

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