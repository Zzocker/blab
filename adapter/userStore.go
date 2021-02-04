package adapter

import (
	"context"
	"fmt"

	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/internal/logger"
	"github.com/Zzocker/blab/model"
	"github.com/Zzocker/blab/pkg/code"
	"github.com/Zzocker/blab/pkg/datastore"
	"github.com/Zzocker/blab/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	userStoreLoggerPrefix = "[adater-userStore] %v"
	usernameKey           = "username"

	queryPerPage int64 = 10
)

type userStore struct {
	ds datastore.SmartDS
}

func (u *userStore) Store(ctx context.Context, user model.User) errors.E {
	logger.L.Info(userStoreLoggerPrefix, fmt.Sprintf("storing username=%s", user.Username))
	err := u.ds.Store(ctx, usernameKey, user)
	if err != nil {
		return err
	}
	logger.L.Info(userStoreLoggerPrefix, fmt.Sprintf("successfully stored username=%s", user.Username))
	return nil
}

func (u *userStore) Get(ctx context.Context, username string) (*model.User, errors.E) {
	logger.L.Info(userStoreLoggerPrefix, fmt.Sprintf("retreving username=%s", username))
	raw, err := u.ds.Get(ctx, usernameKey, username)
	if err != nil {
		return nil, err
	}
	var user model.User
	bErr := bson.Unmarshal(raw, &user)
	if bErr != nil {
		return nil, errors.InitErr(err, code.CodeInternal)
	}
	logger.L.Info(userStoreLoggerPrefix, fmt.Sprintf("retrieved username=%s", username))
	return &user, nil
}

func (u *userStore) Update(ctx context.Context, user model.User) errors.E {
	logger.L.Info(userStoreLoggerPrefix, fmt.Sprintf("updating username=%s", user.Username))
	err := u.ds.Update(ctx, usernameKey, user.Username, user)
	if err != nil {
		return err
	}
	logger.L.Info(userStoreLoggerPrefix, fmt.Sprintf("successfully updated username=%s", user.Username))
	return nil
}

func (u *userStore) Delete(ctx context.Context, username string) errors.E {
	logger.L.Info(userStoreLoggerPrefix, fmt.Sprintf("deleting username=%s", username))
	err := u.ds.Delete(ctx, usernameKey, username)
	if err != nil {
		return err
	}
	logger.L.Info(userStoreLoggerPrefix, fmt.Sprintf("successfully deleted username=%s", username))
	return nil
}

func (u *userStore) Query(ctx context.Context, sortKey string, query map[string]interface{}, pageNumber int64) ([]model.User, errors.E) {
	logger.L.Info(userStoreLoggerPrefix, fmt.Sprintf("quering with query=%v", query))
	raws, err := u.ds.Query(ctx, sortKey, query, pageNumber, queryPerPage)
	if err != nil {
		return nil, err
	}
	users := make([]model.User, len(raws))
	for i := range raws {
		bson.Unmarshal(raws[i], &users[i])
	}
	logger.L.Info(userStoreLoggerPrefix, fmt.Sprintf("query success matchCount=%d", len(users)))
	return users, nil
}

func (u *userStore) Close(ctx context.Context) {
	u.ds.Close(ctx)
}

// CreateUserstore creates new user data store
// create ds -> create index
func CreateUserstore(conf config.DatastoreConf) (*userStore, errors.E) {
	logger.L.Info(userStoreLoggerPrefix, "connecting datastore")
	ds, err := datastore.NewSmartDS(conf)
	if err != nil {
		return nil, err
	}
	logger.L.Info(userStoreLoggerPrefix, "creating username[unique] index")
	err = ds.CreateIndex(context.TODO(), usernameKey, true)
	if err != nil {
		return nil, err
	}
	logger.L.Info(userStoreLoggerPrefix, "connected datastore")
	return &userStore{
		ds: ds,
	}, nil
}
