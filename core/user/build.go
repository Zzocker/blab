package user

import (
	"github.com/Zzocker/blab/adapter"
	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/internal/logger"
	"github.com/Zzocker/blab/pkg/errors"
)

const (
	loggerMsgPrefix = "[core-user] %v"
)

// Build : for building core
// 1. connect user store
func Build(conf config.ApplicationConf) (*userCore, errors.E) {
	logger.L.Info(loggerMsgPrefix, "building")
	uStore, err := adapter.CreateUserstore(conf.Core.User.UserDatastore)
	if err != nil {
		return nil, err
	}
	logger.L.Info(loggerMsgPrefix, "successfully built user core")
	return &userCore{
		uStore: uStore,
	}, nil
}
