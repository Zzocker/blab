package datastore

import (
	"context"

	"github.com/Zzocker/blab/pkg/errors"
)

// DumbDS : represents a datastore which doesn't support query feature
// this type datastore don't care about value of the key ,only matters
// eg redis, etcd
type DumbDS interface {
	Store(ctx context.Context, key string, value []byte, expireIn int64) errors.E
	Get(ctx context.Context, key string) ([]byte, errors.E)
	Delete(ctx context.Context, key string) errors.E
}

// SmartDS : represnets a datastore which support query feature
// eg : mongo
type SmartDS interface {
	Store(ctx context.Context, in interface{}) errors.E
	Get(ctx context.Context, isbn string) ([]byte, errors.E)
	Update(ctx context.Context, key, value string, in interface{}) errors.E
	Delete(ctx context.Context, key, value string) errors.E
	Query(ctx context.Context, query map[string]interface{}, pageNumber, perPage int) ([][]byte, errors.E)
}