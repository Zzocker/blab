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
	Port int `yaml:"port"`
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
