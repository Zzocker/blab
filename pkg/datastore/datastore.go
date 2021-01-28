package datastore

import (
	"context"
	"fmt"
	"os"

	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/pkg/errors"
	"github.com/Zzocker/blab/pkg/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DS is datastore layer for this application
type DS interface {
	Store(ctx context.Context, obj interface{}) errors.E
	Get(ctx context.Context, key string, value string) ([]byte, errors.E)
	Update(ctx context.Context, key, value string, obj interface{}) errors.E
	Query(ctx context.Context, query map[string]interface{}) ([][]byte, errors.E)
	Delete(ctx context.Context, key, value string) errors.E
	Close(ctx context.Context)
}

func NewMongo(conf config.DatastoreConf, l log.Logger) DS {
	address := fmt.Sprintf("mongodb://%s:%s@%s", conf.Username, conf.Password, conf.URL)
	client, err := mongo.NewClient(options.Client().ApplyURI(address))
	if err != nil {
		l.Error(err)
		os.Exit(1)
	}
	if err = client.Connect(context.TODO()); err != nil {
		l.Error(err)
		os.Exit(1)
	}
	if err = client.Ping(context.TODO(), nil); err != nil {
		l.Error(err)
		os.Exit(1)
	}
	l.Infof("connected to [%s][%s]", conf.URL, conf.Collection)
	return &mongoDS{
		db: client.Database(conf.Database).Collection(conf.Collection),
	}
}
