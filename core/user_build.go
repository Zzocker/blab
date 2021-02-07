package core

import (
	"github.com/Zzocker/blab/adapters"
	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/internal/logger"
)

const (
	userCorePkg = "core-user"
)

type userBuilder struct{}

func (u *userBuilder) build(conf *config.ApplicationConf) error {
	logger.L.Infof(-1, userCorePkg, "building core")
	logger.L.Debugf(-1, userCorePkg, "createing user datastore")
	uStore, err := adapters.CreateUserstore(&conf.Core.User.UserStore)
	if err != nil {
		return err
	}
	varUserCore = &userCore{
		uStore: uStore,
	}
	logger.L.Infof(-1, userCorePkg, "core build done")
	return nil
}
