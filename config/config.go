package config

import (
	"os"

	"github.com/Zzocker/blab/pkg/log"
	"gopkg.in/yaml.v3"
)

type ApplicationConf struct {
	Port     int       `yaml:"port"`
	LogLevel log.Level `yaml:"logLevel"`
}

// LoadConfig will read config file form specified path
func LoadConfig(filename string) (*ApplicationConf, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var conf ApplicationConf
	if err = yaml.NewDecoder(f).Decode(&conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
