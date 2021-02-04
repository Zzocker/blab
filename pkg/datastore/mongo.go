package datastore

import (
	"context"
	"fmt"

	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/internal/logger"
	"github.com/Zzocker/blab/pkg/code"
	"github.com/Zzocker/blab/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	smartDSLoggerPrefix = "[datastore-smartDS] %s"
)

type mongoDS struct {
	db *mongo.Collection
}

func newMongoDS(conf config.DatastoreConf) (SmartDS, errors.E) {
	logger.L.Info(smartDSLoggerPrefix, fmt.Sprintf("connecting url=mongodb://%s database=%s collection=%s", conf.URL, conf.Database, conf.Collection))
	URI := fmt.Sprintf("mongodb://%s:%s@%s", conf.Username, conf.Password, conf.URL)
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
	if err != nil {
		return nil, errors.InitErr(err, code.CodeInternal)
	}
	if err = client.Connect(context.TODO()); err != nil {
		return nil, errors.InitErr(err, code.CodeInternal)
	}
	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, errors.InitErr(err, code.CodeInternal)
	}
	logger.L.Info(smartDSLoggerPrefix, "connected")
	return &mongoDS{
		db: client.Database(conf.Database).Collection(conf.Collection),
	}, nil
}

func (m *mongoDS) Close(ctx context.Context) {
	logger.L.Info(smartDSLoggerPrefix, "closed")
	m.db.Database().Client().Disconnect(ctx)
}

// Store :
func (m *mongoDS) Store(ctx context.Context, key string, in interface{}) errors.E {
	logger.L.Info(smartDSLoggerPrefix, fmt.Sprintf("storing doc with id=%s", key))
	_, err := m.db.InsertOne(ctx, in)
	if isDuplicate(err) {
		return errors.InitErr(err, code.CodeAlreadyExists)
	} else if err != nil {
		return errors.InitErr(err, code.CodeInternal)
	}
	logger.L.Info(smartDSLoggerPrefix, fmt.Sprintf("doc with id=%s successfully stored", key))
	return nil
}

func (m *mongoDS) Get(ctx context.Context, key, value string) ([]byte, errors.E) {
	logger.L.Info(smartDSLoggerPrefix, fmt.Sprintf("retreving doc with %s=%s", key, value))
	reply := m.db.FindOne(ctx, bson.M{key: value})
	if reply.Err() == mongo.ErrNoDocuments {
		return nil, errors.InitErr(reply.Err(), code.CodeNotFound)
	} else if reply.Err() != nil {
		return nil, errors.InitErr(reply.Err(), code.CodeInternal)
	}
	raw, err := reply.DecodeBytes()
	if err != nil {
		return nil, errors.InitErr(reply.Err(), code.CodeInternal)
	}
	logger.L.Info(smartDSLoggerPrefix, fmt.Sprintf("retrived doc with %s=%s", key, value))
	return raw, nil
}

func (m *mongoDS) Update(ctx context.Context, key, value string, in interface{}) errors.E {
	logger.L.Info(smartDSLoggerPrefix, fmt.Sprintf("updating doc with %s=%s", key, value))
	reply, err := m.db.UpdateOne(ctx, bson.M{key: value}, bson.M{
		"$set": in,
	})
	if err != nil {
		return errors.InitErr(err, code.CodeInternal)
	}
	if reply.MatchedCount != 1 {
		return errors.InitErr(fmt.Errorf("update fail : doc with %s=%s not found", key, value), code.CodeNotFound)
	}
	logger.L.Info(smartDSLoggerPrefix, fmt.Sprintf("updated doc with %s=%s", key, value))
	return nil
}

func (m *mongoDS) Delete(ctx context.Context, key, value string) errors.E {
	logger.L.Info(smartDSLoggerPrefix, fmt.Sprintf("deleting doc with %s=%s", key, value))
	reply, err := m.db.DeleteOne(ctx, bson.M{key: value})
	if err != nil {
		return errors.InitErr(err, code.CodeInternal)
	}
	if reply.DeletedCount != 1 {
		return errors.InitErr(fmt.Errorf("delete fail : doc with %s=%s not found", key, value), code.CodeNotFound)
	}
	logger.L.Info(smartDSLoggerPrefix, fmt.Sprintf("deleted doc with %s=%s", key, value))
	return nil
}

func (m *mongoDS) Query(ctx context.Context, sortingKey string, query map[string]interface{}, pageNumber, perPage int64) ([][]byte, errors.E) {
	logger.L.Info(smartDSLoggerPrefix, fmt.Sprintf("querying with query=%v pageNumber=%d perPage=%d sortKey=%s", query, pageNumber, perPage, sortingKey))
	skip := perPage * (pageNumber - 1)
	if skip < 0 {
		skip = 0
	}
	if sortingKey == "" {
		sortingKey = "_id"
	}
	opts := options.FindOptions{
		Limit: &perPage,
		Skip:  &skip,
		Sort:  bson.D{{sortingKey, 1}},
	}
	cur, err := m.db.Find(ctx, query, &opts)
	if err != nil {
		return nil, errors.InitErr(err, code.CodeInternal)
	}
	defer cur.Close(ctx)
	raws := make([][]byte, 0, 0)
	for cur.Next(ctx) {
		raws = append(raws, cur.Current)
	}
	logger.L.Info(smartDSLoggerPrefix, fmt.Sprintf("query success matchCount=%d", len(raws)))
	return raws, nil
}

func (m *mongoDS) CreateIndex(ctx context.Context, key string, unique bool) errors.E {
	logger.L.Info(smartDSLoggerPrefix, fmt.Sprintf("creating index with key=%s and isUnique=%v", key, unique))
	indexname, err := m.db.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{key, 1}},
		Options: &options.IndexOptions{
			Unique: &unique,
		},
	})
	if err != nil {
		return errors.InitErr(err, code.CodeInternal)
	}
	logger.L.Info(smartDSLoggerPrefix, fmt.Sprintf("created index name=%s with key=%s and isUnique=%v", indexname, key, unique))
	return nil
}

// helper
func isDuplicate(err error) bool {
	if mErr, ok := err.(mongo.WriteException); ok {
		for _, e := range mErr.WriteErrors {
			if e.Code == 11000 {
				return true
			}
		}
	}
	return false
}
