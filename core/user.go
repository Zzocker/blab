package core

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/Zzocker/blab/internal/logger"
	"github.com/Zzocker/blab/model"
	"github.com/Zzocker/blab/pkg/code"
	"github.com/Zzocker/blab/pkg/errors"
	"github.com/Zzocker/blab/pkg/util"
	"github.com/Zzocker/blab/ports"
)

type userCore struct {
	uStore ports.UserDatastorePort
}

func (u *userCore) Register(ctx context.Context, in RegisterUserInput, password string) errors.E {
	logger.L.Info(userloggerPrefix, fmt.Sprintf("Registering %s", in.Username))
	err := in.validate(password)
	if err != nil {
		return err
	}
	hashPass := hash(password)
	user := in.toUser(hashPass)
	err = u.uStore.Store(ctx, user)
	if err != nil {
		return err
	}
	logger.L.Info(userloggerPrefix, fmt.Sprintf("successfully Registered %s", in.Username))
	return nil
}

func (u *userCore) GetUser(ctx context.Context, username string) (*model.User, errors.E) {
	return u.uStore.Get(ctx, username)
}

func (u *userCore) Update(ctx context.Context, reader io.Reader) (*model.User, errors.E) {
	logger.L.Info(userloggerPrefix, "updating userprofile")
	owner := util.UnWrapCtx(ctx, util.UsernameCtxKey).(string)
	logger.L.Info(userloggerPrefix, "geting old userprofile")
	usr, err := u.uStore.Get(ctx, owner)
	if err != nil {
		return nil, err
	}
	logger.L.Info(userloggerPrefix, "updating old userprofile with request")
	jErr := json.NewDecoder(reader).Decode(&usr)
	if jErr != nil {
		return nil, errors.InitErr(fmt.Errorf("invalid json body"), code.CodeInvalidArgument)
	}
	logger.L.Info(userloggerPrefix, "storing updated userprofile")
	usr.Username = owner
	err = u.uStore.Update(ctx, *usr)
	if err != nil {
		return nil, err
	}
	logger.L.Info(userloggerPrefix, "successfully updated")
	return usr, nil
}

func (u *userCore) Delete(ctx context.Context) errors.E {
	owner := util.UnWrapCtx(ctx, util.UsernameCtxKey).(string)
	return u.uStore.Delete(ctx, owner)
}
