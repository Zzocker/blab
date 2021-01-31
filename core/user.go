package core

import (
	"github.com/Zzocker/blab/core/ports"
	"github.com/Zzocker/blab/model"
	"github.com/Zzocker/blab/pkg/errors"
)

type userCore struct {
	uStore ports.UserStorePort
}

func (u *userCore) Register(user model.User) errors.E {
	return nil
}
