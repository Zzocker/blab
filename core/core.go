package core

import (
	"context"

	"github.com/Zzocker/blab/core/user"
	"github.com/Zzocker/blab/model"
	"github.com/Zzocker/blab/pkg/errors"
)

type UserCore interface {
	Register(ctx context.Context, in user.Register) (*model.User, errors.E)
	Get(ctx context.Context, username string) (*model.User, errors.E)
}
