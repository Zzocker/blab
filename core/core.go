package core

import (
	"context"
	"io"

	"github.com/Zzocker/blab/core/user"
	"github.com/Zzocker/blab/model"
	"github.com/Zzocker/blab/pkg/errors"
)

type UserCore interface {
	Register(ctx context.Context, in user.Register) (*model.User, errors.E)
	Get(ctx context.Context, username string) (*model.User, errors.E)
	Update(ctx context.Context, username string, reader io.Reader) (*model.User, errors.E)
	Delete(ctx context.Context, username string) errors.E
}

type OAuthCore interface {
	Login(ctx context.Context, username, password string) (map[string]model.Token, errors.E)
}

type BookCore interface {
}
