package datastore

import (
	"context"

	"github.com/Zzocker/blab/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDS struct {
	db *mongo.Collection
}

func (m *mongoDS) Store(ctx context.Context, in interface{}) errors.E {
	_, err := m.db.InsertOne(ctx, in)
	if isDuplicate(err) {
		return errors.New(errors.CodeAlreadyExists, "item already exists")
	} else if err != nil {
		return errors.New(errors.CodeInternalErr, "failed to store")
	}
	return nil
}

func (m *mongoDS) Get(ctx context.Context, key, value string) ([]byte, errors.E) {
	reply := m.db.FindOne(ctx, bson.M{key: value})
	if reply.Err() == mongo.ErrNoDocuments {
		return nil, errors.New(errors.CodeNotFound, "item not found")
	} else if reply.Err() != nil {
		return nil, errors.New(errors.CodeInternalErr, "failed to get")
	}
	raw, err := reply.DecodeBytes()
	if err != nil {
		return nil, errors.New(errors.CodeInternalErr, "failed to decode")
	}
	return raw, nil
}

func (m *mongoDS) Update(ctx context.Context, key, value string, in interface{}) errors.E {
	reply, err := m.db.UpdateOne(ctx, bson.M{key: value}, bson.M{
		"$set": in,
	})
	if err != nil {
		return errors.New(errors.CodeInternalErr, "failed to decode")
	}
	if reply.MatchedCount != 1 {
		return errors.New(errors.CodeNotFound, "item not found")
	}
	return nil
}

func (m *mongoDS) Delete(ctx context.Context, key, value string) errors.E {
	reply, err := m.db.DeleteOne(ctx, bson.M{key: value})
	if err != nil {
		return errors.New(errors.CodeInternalErr, "failed to decode")
	}
	if reply.DeletedCount != 1 {
		return errors.New(errors.CodeNotFound, "item not found")
	}
	return nil
}

func (m *mongoDS) Query(ctx context.Context, sortingKey string, query map[string]interface{}, pageNumber, perPage int64) ([][]byte, errors.E) {
	skip := perPage * (pageNumber - 1)
	if skip < 0 {
		skip = 0
	}
	opts := options.FindOptions{
		Limit: &perPage,
		Skip:  &skip,
		Sort:  bson.D{{sortingKey, 1}},
	}
	cur, err := m.db.Find(ctx, query, &opts)
	if err != nil {
		return nil, errors.New(errors.CodeInternalErr, "get make query")
	}
	defer cur.Close(ctx)
	raws := make([][]byte, 0, 0)
	for cur.Next(ctx) {
		raws = append(raws, cur.Current)
	}
	return raws, nil
}
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
