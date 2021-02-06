package core

import (
	"github.com/Zzocker/blab/adapter"
	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/internal/logger"
	"github.com/Zzocker/blab/pkg/errors"
)

const (
	userloggerPrefix = "[core-user] %v"
)

var (
	uCore UserCore
)

// CoreBuilder implements core builderinterface
type userBuilder struct{}

// Build : for building core
// 1. connect user store
func (c *userBuilder) build(conf config.ApplicationConf) errors.E {
	logger.L.Info(userloggerPrefix, "building core")
	uStore, err := adapter.CreateUserstore(conf.Core.User.UserDatastore)
	if err != nil {
		return err
	}
	logger.L.Info(userloggerPrefix, "successfully built core")
	uCore = &userCore{
		uStore: uStore,
	}
	return nil
}

// GetUserCore returns global level user core variable
func GetUserCore() UserCore {
	return uCore
}
