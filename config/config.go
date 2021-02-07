package config

import (
	"os"

	"github.com/Zzocker/blab/pkg/log"
	"gopkg.in/yaml.v3"
)

type ApplicationConf struct {
	Port     int       `yaml:"port"`
	LogLevel log.Level `yaml:"logLevel"`
	Core     CoreConf  `yaml:"core"`
}

type CoreConf struct {
	User UserCoreConf `yaml:"user"`
}

type UserCoreConf struct {
	UserStore DatastoreConf `yaml:"userDatastore"`
}
type DatastoreConf struct {
	URL        string `yaml:"URL"`
	Username   string `yaml:"Username"`
	Password   string `yaml:"Password"`
	Database   string `yaml:"Database"`
	Collection string `yaml:"Collection"`
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
