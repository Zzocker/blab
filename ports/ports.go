package ports

import (
	"context"

	"github.com/Zzocker/blab/model"
	"github.com/Zzocker/blab/pkg/errors"
)

// UserDatastorePort provide abstraction over user profile database
// UserDatastorePort will be used by
// UserCore
type UserDatastorePort interface {
	Store(ctx context.Context, user model.User) errors.E
	Get(ctx context.Context, username string) (*model.User, errors.E)
	Update(ctx context.Context, user model.User) errors.E
	Delete(ctx context.Context, username string) errors.E
	Query(ctx context.Context, sortKey string, query map[string]interface{}, pageNumber int64) ([]model.User, errors.E)
	Close(ctx context.Context)
}
