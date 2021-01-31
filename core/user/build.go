package user

import (
	"github.com/Zzocker/blab/adapters"
	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/pkg/errors"
)

func NewUserCore(conf config.UserCoreConf) (*userCore, errors.E) {
	store, err := adapters.CreateUserStore(conf.UserStoreConf)
	if err != nil {
		return nil, err
	}
	return &userCore{
		uStore: store,
	}, nil
}
