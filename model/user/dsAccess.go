package user

import (
	"context"

	"github.com/Zzocker/blab/pkg/datastore"
	"github.com/Zzocker/blab/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	usernameKey = "username"
)

type DSAccess struct {
	db datastore.DS
}

// NewAccess creates a new user dataaccess layer
func NewAccess(db datastore.DS) *DSAccess {
	// create username and email index
	// TODO
	return &DSAccess{
		db: db,
	}
}

func (d *DSAccess) Create(ctx context.Context, user User) errors.E {
	return d.db.Store(ctx, user)
}

func (d *DSAccess) Get(ctx context.Context, username string) (*User, errors.E) {
	raw, err := d.db.Get(ctx, usernameKey, username)
	if err != nil {
		return nil, err
	}
	var usr User
	bson.Unmarshal(raw, &usr)
	return &usr, nil
}

func (d *DSAccess) GetRaw(ctx context.Context, username string) ([]byte, errors.E) {
	return d.db.Get(ctx, usernameKey, username)
}

func (d *DSAccess) Delete(ctx context.Context, username string) errors.E {
	return d.db.Delete(ctx, usernameKey, username)
}

func (d *DSAccess) Update(ctx context.Context, user *User) errors.E {
	return d.db.Update(ctx, usernameKey, user.Username, user)
}
