package core

import (
	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/internal/logger"
	"github.com/Zzocker/blab/pkg/errors"
)

const (
	coreLoggerPrefix = "[core] %v"
)

var (
	coreFactory = []coreBuilder{
		&userBuilder{},
	}
)

type coreBuilder interface {
	build(conf config.ApplicationConf) errors.E
}

// BuildAll will build all core
func BuildAll(conf config.ApplicationConf) errors.E {
	logger.L.Info(coreLoggerPrefix, "building all core")
	for coreName := range coreFactory {
		if err := coreFactory[coreName].build(conf); err != nil {
			return err
		}
	}
	logger.L.Info(coreLoggerPrefix, "all core build successful")
	return nil
}
