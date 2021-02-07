package core

import (
	"os"

	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/internal/logger"
)

const (
	corePkgName = "core"
)

type coreBuilder interface {
	build(conf *config.ApplicationConf) error
}

var (
	factory = []coreBuilder{
		&userBuilder{},
	}
)

func BuildAll(conf *config.ApplicationConf) {
	logger.L.Infof(-1, corePkgName, "building all core")
	var err error
	for _, c := range factory {
		if err = c.build(conf); err != nil {
			logger.L.Errorf(-1, corePkgName, "failed to build core : %v", err)
			os.Exit(1)
		}
	}
	logger.L.Infof(-1, corePkgName, "core build successful")
}
