package core

import (
	"context"
	"io"

	"github.com/Zzocker/blab/model"
	"github.com/Zzocker/blab/pkg/errors"
)

// UserCore represents core business log for user
type UserCore interface {
	Register(ctx context.Context, in RegisterUserInput, password string) errors.E
	GetUser(ctx context.Context, username string) (*model.User, errors.E)
	Update(ctx context.Context, reader io.Reader) (*model.User, errors.E)
	Delete(ctx context.Context) errors.E
}
