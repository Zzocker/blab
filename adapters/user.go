package adapters

import (
	"context"

	"github.com/Zzocker/blab/model"
	"github.com/Zzocker/blab/pkg/datastore"
	"github.com/Zzocker/blab/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	userPrimaryKeyName = "username"
)

type userStore struct {
	db datastore.SmartDS
}

func (b *userStore) Store(ctx context.Context, user model.User) errors.E {
	return b.db.Store(ctx, user)
}
func (b *userStore) Get(ctx context.Context, username string) (*model.User, errors.E) {
	raw, err := b.db.Get(ctx, userPrimaryKeyName, username)
	if err != nil {
		return nil, err
	}
	var user model.User
	bErr := bson.Unmarshal(raw, &user)
	if bErr != nil {
		return nil, errors.New(errors.CodeInternalErr, "failed to unmarshal raw data")
	}
	return &user, nil
}
func (b *userStore) Update(ctx context.Context, username string, user model.Book) errors.E {
	return b.db.Update(ctx, userPrimaryKeyName, username, user)
}
func (b *userStore) Delete(ctx context.Context, username string) errors.E {
	return b.db.Delete(ctx, userPrimaryKeyName, username)
}
func (b *userStore) Query(ctx context.Context, sortKey string, query map[string]interface{}, pageNumber, perPage int64) ([]model.User, errors.E) {
	raws, err := b.db.Query(ctx, sortKey, query, pageNumber, perPage)
	if err != nil {
		return nil, err
	}
	users := make([]model.User, len(raws))
	for i := range raws {
		bson.Unmarshal(raws[i], &users[i])
	}
	return users, nil
}
