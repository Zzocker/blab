package datastore

import (
	"context"

	"github.com/Zzocker/blab/pkg/errors"
)

// DS is datastore layer for this application
type DS interface {
	Store(ctx context.Context, obj interface{}) errors.E
	Get(ctx context.Context, key string, value string) ([]byte, errors.E)
	Update(ctx context.Context, key, value string, obj interface{}) errors.E
	Query(ctx context.Context, query map[string]interface{}) ([][]byte, errors.E)
	Delete(ctx context.Context, key, value string) errors.E
}
