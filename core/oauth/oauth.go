package oauth

import (
	"context"

	"github.com/Zzocker/blab/core/ports"
	"github.com/Zzocker/blab/model"
	"github.com/Zzocker/blab/pkg/errors"
	"github.com/Zzocker/blab/pkg/util"
)

type oauthCore struct {
	tStore ports.OAuthStore
	uStore ports.UserStorePort
}

func (o *oauthCore) Login(ctx context.Context, username, password string) (map[string]model.Token, errors.E) {
	usr, err := o.uStore.Get(ctx, username)
	if err != nil {
		return nil, err
	}
	if usr.Password != util.Hash(password) {
		return nil, errors.New(errors.CodeUnauthorized, "incorrect password")
	}
	refresh := model.NewRefreshToken(username)
	access := model.NewAccessToken(username)
	err = o.tStore.Store(ctx, refresh)
	if err != nil {
		return nil, err
	}
	err = o.tStore.Store(ctx, access)
	if err != nil {
		o.tStore.Delete(ctx, refresh.ID)
		return nil, err
	}
	return map[string]model.Token{
		"refresh": refresh,
		"access":  access,
	}, nil
}
