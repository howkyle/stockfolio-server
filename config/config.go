package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	DB   string `yaml:"db"`
	Port string `yaml:"port"`
}

//reads config from yaml file
func Config() *Configuration {

	file, err := os.ReadFile("./config.yml")
	if err != nil {
		panic("unable to load config data:" + err.Error())
	}
	c := Configuration{}
	err = yaml.Unmarshal(file, &c)
	if err != nil {
		panic("unable to unmarshall config:" + err.Error())
	}

	return &c
}
