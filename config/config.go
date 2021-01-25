package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// C represents loaded config from yaml file
type C struct {
	Port int `yaml:"port"`
}

// Init load config file from confPath
// os.Exit(1) is called if any error occurred while reading config file
func Init(confPath string) *C {
	f, err := os.Open(confPath)
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()
	var conf C
	if err = yaml.NewDecoder(f).Decode(&conf); err != nil {
		os.Exit(1)
	}
	return &conf
}
