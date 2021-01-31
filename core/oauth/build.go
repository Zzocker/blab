package oauth

import (
	"github.com/Zzocker/blab/adapters"
	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/pkg/errors"
)

func NewOAuthCore(conf config.C) (*oauthCore, errors.E) {
	uStore, err := adapters.CreateUserStore(conf.Core.User.UserStoreConf)
	if err != nil {
		return nil, err
	}
	tStore, err := adapters.CreateOAuthStore(conf.Core.OAuth.TokenStoreConf)
	if err != nil {
		return nil, err
	}
	return &oauthCore{
		uStore: uStore,
		tStore: tStore,
	}, nil
}
