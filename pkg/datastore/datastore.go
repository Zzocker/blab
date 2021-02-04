package datastore

import (
	"context"

	"github.com/Zzocker/blab/pkg/errors"
)

// SmartDS : represnets a datastore which support query feature
// eg : mongo
type SmartDS interface {
	Store(ctx context.Context, key string, in interface{}) errors.E
	Get(ctx context.Context, key, value string) ([]byte, errors.E)
	Update(ctx context.Context, key, value string, in interface{}) errors.E
	Delete(ctx context.Context, key, value string) errors.E
	Query(ctx context.Context, sortingKey string, query map[string]interface{}, pageNumber, perPage int64) ([][]byte, errors.E)
	CreateIndex(ctx context.Context, key string, unique bool) errors.E
	Close(ctx context.Context)
}
