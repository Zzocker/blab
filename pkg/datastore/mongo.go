package datastore

import (
	"context"
	"net/http"

	"github.com/Zzocker/blab/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoDS struct {
	db *mongo.Collection
}

// Store : obj should have bson field tags
func (m *mongoDS) Store(ctx context.Context, obj interface{}) errors.E {
	_, err := m.db.InsertOne(ctx, obj)
	if isDuplicate(err) {
		return errors.New(http.StatusConflict, "duplicate entry")
	} else if err != nil {
		return errors.New(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (m *mongoDS) Get(ctx context.Context, key string, value string) ([]byte, errors.E) {
	result := m.db.FindOne(ctx, bson.M{key: value})
	if result.Err() == mongo.ErrNoDocuments {
		return nil, errors.New(http.StatusNotFound, "entry not found")
	} else if result.Err() != nil {
		return nil, errors.New(http.StatusInternalServerError, result.Err().Error())
	}
	raw, err := result.DecodeBytes()
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, err.Error())
	}
	return raw, nil
}

func (m *mongoDS) Update(ctx context.Context, key, value string, obj interface{}) errors.E {
	_, err := m.db.UpdateOne(ctx, bson.M{key: value}, bson.M{"$set": obj})
	if err == mongo.ErrNoDocuments {
		return errors.New(http.StatusNotFound, "entry not found")
	} else if err != nil {
		return errors.New(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (m *mongoDS) Query(ctx context.Context, query map[string]interface{}) ([][]byte, errors.E) {
	cur, err := m.db.Find(ctx, query)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, err.Error())
	}
	defer cur.Close(ctx)
	raws := make([][]byte, 0, 10)
	for cur.Next(ctx) {
		raws = append(raws, cur.Current)
	}
	return raws, nil
}

func (m *mongoDS) Delete(ctx context.Context, key, value string) errors.E {
	result, err := m.db.DeleteOne(ctx, bson.M{key: value})
	if err != nil {
		return errors.New(http.StatusInternalServerError, err.Error())
	}
	if result.DeletedCount != 1 {
		return errors.New(http.StatusNotFound, "entry not found")
	}
	return nil
}

func (m *mongoDS) Close(ctx context.Context) {
	m.db.Database().Client().Disconnect(ctx)
}

// helper
func isDuplicate(err error) bool {
	if merr, ok := err.(mongo.WriteException); ok {
		for _, e := range merr.WriteErrors {
			if e.Code == 11000 {
				return true
			}
		}
	}
	return false
}
