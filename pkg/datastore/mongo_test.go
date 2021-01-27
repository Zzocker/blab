package datastore

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const testMongoAddres = "mongodb://root:password@localhost:27017"
const testMongoDatabase = "tblab"
const testMongoCollection = "tcol"

var mds *mongoDS

func TestMongoType(t *testing.T) {
	var l interface{} = &mongoDS{}
	_, ok := l.(DS)
	assert.True(t, ok)
}

type testStruct struct {
	Username string `bson:"username"`
	Name     string `bson:"name"`
	Age      int    `bson:"age"`
}

func TestMongoMain(t *testing.T) {
	client, err := mongo.NewClient(options.Client().ApplyURI(testMongoAddres))
	if err != nil {
		t.Log(err)
		os.Exit(1)
	}
	if err = client.Connect(context.TODO()); err != nil {
		t.Log(err)
		os.Exit(1)
	}
	if err = client.Ping(context.TODO(), nil); err != nil {
		t.Log(err)
		os.Exit(1)
	}
	collection := client.Database(testMongoDatabase).Collection(testMongoCollection)
	var unique bool = true
	_, err = collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{{"username", 1}},
		Options: &options.IndexOptions{
			Unique: &unique,
		},
	})
	if err != nil {
		t.Log(err)
		os.Exit(1)
	}
	mds = &mongoDS{
		db: collection,
	}
}

var testUser = testStruct{
	Username: "Zzocker",
	Name:     "Pritam Singh",
	Age:      18,
}

func TestMongoStore(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	err := mds.Store(ctx, testUser)
	assert.NoError(err)

	err = mds.Store(ctx, testUser)
	assert.Error(err)
	// username already exists
	assert.Equal(http.StatusConflict, err.GetStatus())
}

func TestMongoGet(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	// happy flow
	raw, err := mds.Get(ctx, "username", testUser.Username)
	assert.NoError(err)
	var user testStruct
	bson.Unmarshal(raw, &user)
	assert.Equal(testUser, user)

	// key is incorrect
	_, err = mds.Get(ctx, "usern", testUser.Username)
	assert.Error(err)
	assert.Equal(http.StatusNotFound, err.GetStatus())

	// value not found
	_, err = mds.Get(ctx, "username", "notfound")
	assert.Error(err)
	assert.Equal(http.StatusNotFound, err.GetStatus())
}

func TestMongoUpdate(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	user := testStruct{
		Username: testUser.Username,
		Name:     "New Name",
		Age:      testUser.Age,
	}
	// happy flow
	err := mds.Update(ctx, "username", testUser.Username, user)
	assert.NoError(err)
	raw, _ := mds.Get(ctx, "username", user.Username)
	var newUser testStruct
	bson.Unmarshal(raw, &newUser)
	assert.Equal(newUser.Name, user.Name)
}

func TestMongoDelete(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	// happy flow
	err := mds.Delete(ctx, "username", testUser.Username)
	assert.NoError(err)

	// entry not found
	err = mds.Delete(ctx, "username", "not_found")
	assert.Error(err)
	assert.Equal(http.StatusNotFound, err.GetStatus())
}

func TestMongoQuery(t *testing.T) {
	assert := assert.New(t)
	ctx := context.TODO()
	// load data onto database
	loadSize := 10
	for i := 0; i < loadSize; i++ {
		testUser.Username = fmt.Sprintf("query_username_%d", i)
		mds.Store(ctx, testUser)
	}
	entries, err := mds.Query(ctx, map[string]interface{}{
		"name": "Pritam Singh",
	})
	assert.NoError(err)
	users := make([]testStruct, len(entries))
	for i, entry := range entries {
		bson.Unmarshal(entry, &users[i])
	}
	assert.Equal(loadSize, len(users))
}

func TestMongoDone(t *testing.T) {
	mds.db.Database().Drop(context.Background())
	mds.db.Database().Client().Disconnect(context.Background())
}