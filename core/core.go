package core

import (
	"github.com/Zzocker/blab/adapters"
	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/model"
	"github.com/Zzocker/blab/pkg/errors"
)

type UserCore interface {
	Register(user model.User) errors.E
}

func NewUserCore(conf config.UserCoreConf) (UserCore, errors.E) {
	store, err := adapters.CreateUserStore(conf.UserStoreConf)
	if err != nil {
		return nil, err
	}
	return &userCore{
		uStore: store,
	}, nil
}
