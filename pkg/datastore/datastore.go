package datastore

import (
	"context"
	"fmt"

	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/pkg/errors"
	"github.com/gomodule/redigo/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	Get(ctx context.Context, key, value string) ([]byte, errors.E)
	Update(ctx context.Context, key, value string, in interface{}) errors.E
	Delete(ctx context.Context, key, value string) errors.E
	Query(ctx context.Context, sortingKey string, query map[string]interface{}, pageNumber, perPage int64) ([][]byte, errors.E)
}

// NewDumbDS creates a new DumbDB; database which doesn't suport query
// currently redis is being used
func NewDumbDS(conf config.DatastoreConf) (DumbDS, errors.E) {
	pool := redis.Pool{
		DialContext: func(ctx context.Context) (redis.Conn, error) {
			return redis.DialContext(
				ctx,
				"tcp",
				conf.URL,
				redis.DialUsername(conf.Username),
				redis.DialPassword(conf.Password),
			)
		},
	}
	conn, err := pool.GetContext(context.TODO())
	if err != nil {
		return nil, errors.New(errors.CodeInternalErr, "failed to dial")
	}
	defer conn.Close()
	if _, err := conn.Do("PING"); err != nil {
		return nil, errors.New(errors.CodeInternalErr, "failed to ping redis")
	}
	return &redisStore{p: &pool}, nil
}

// NewSmartDS create a new smart datastore
// smart datastore support rich query feature
// currently mongodb
func NewSmartDS(conf config.DatastoreConf) (SmartDS, errors.E) {
	addrs := fmt.Sprintf("mongodb://%s:%s@%s", conf.Username, conf.Password, conf.URL)
	client, err := mongo.NewClient(options.Client().ApplyURI(addrs))
	if err != nil {
		return nil, errors.New(errors.CodeInternalErr, fmt.Sprintf("failed to create client %+v", err))
	}
	err = client.Connect(context.Background())
	if err != nil {
		return nil, errors.New(errors.CodeInternalErr, fmt.Sprintf("failed to connect %+v", err))
	}
	return &mongoDS{
		db: client.Database(conf.Database).Collection(conf.Collection),
	}, nil
}
