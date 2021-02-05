package config

import (
	"os"

	"github.com/Zzocker/blab/pkg/code"
	"github.com/Zzocker/blab/pkg/errors"
	"gopkg.in/yaml.v3"
)

// ApplicationConf represents all configs for this project
// Port : is port on which this server will run
type ApplicationConf struct {
	Port int  `yaml:"Port"`
	Core Core `yaml:"Core"`
}

// ReadConf will read read project configuration file
// and decode to ApplicationConf
func ReadConf(file string) (*ApplicationConf, errors.E) {
	f, err := os.Open(file)
	if err != nil {
		return nil, errors.InitErr(err, code.CodeInternal)
	}
	defer f.Close()
	var conf ApplicationConf
	err = yaml.NewDecoder(f).Decode(&conf)
	if err != nil {
		return nil, errors.InitErr(err, code.CodeInternal)
	}
	return &conf, nil
}

// Core are configurations of core business logics
type Core struct {
	User UserCore `yaml:"User"`
}

// UserCore is configuration for user core
type UserCore struct {
	UserDatastore DatastoreConf `yaml:"UserDatastore"`
}

// DatastoreConf represents configuration of a datastore (like mongo/redis)
type DatastoreConf struct {
	URL        string `yaml:"URL"`
	Username   string `yaml:"Username"`
	Password   string `yaml:"Password"`
	Database   string `yaml:"Database"`
	Collection string `yaml:"Collection"`
}
