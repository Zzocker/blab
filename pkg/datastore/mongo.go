package datastore

import (
	"context"
	"fmt"
	"time"

	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/pkg/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	smartDSPkgName = "datastore-smartDS"
)

type mongoDS struct {
	db *mongo.Collection
}

func newMongoDS(reqID int64, conf config.DatastoreConf, logger log.Logger) (*mongoDS, error) {
	logger.Infof(reqID, smartDSPkgName, "connecting url=mongodb://%s database=%s collection=%s", conf.URL, conf.Database, conf.Collection)
	URI := fmt.Sprintf("mongodb://%s:%s@%s", conf.Username, conf.Password, conf.URL)
	logger.Debugf(reqID, smartDSPkgName, "creating new smartDS client")
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
	if err != nil {
		return nil, err
	}
	logger.Debugf(reqID, smartDSPkgName, "connecting with created client")
	if err = client.Connect(context.Background()); err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	logger.Debugf(reqID, smartDSPkgName, "pinging smart datastore")
	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	logger.Infof(reqID, smartDSPkgName, "connected")
	return &mongoDS{
		db: client.Database(conf.Database).Collection(conf.Collection),
	}, nil
}
