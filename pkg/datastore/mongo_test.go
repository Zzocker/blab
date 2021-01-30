package datastore

import (
	"context"
	"log"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type testStruct struct {
	F1 string  `bson:"f1"`
	F2 int     `bson:"f2"`
	F3 float64 `bson:"f3"`
	F4 int64   `bson:"f4"`
}

func TestMongoDSType(t *testing.T) {
	var l interface{} = &mongoDS{}
	_, ok := l.(SmartDS)
	assert.True(t, ok)
}

func TestMongoStore(t *testing.T) {
	col := getTestMongoCollection()
	defer cleanUP(col)
	store := mongoDS{
		db: col,
	}
	is := assert.New(t)

	testData := testStruct{
		F1: "f1",
		F2: 23,
		F3: 65.32,
		F4: time.Now().Unix(),
	}

	// happy flow
	err := store.Store(context.Background(), testData)
	is.NoError(err)
	raw, _ := col.FindOne(context.TODO(), bson.M{"f1": "f1"}).DecodeBytes()
	var replyStruct testStruct
	bson.Unmarshal(raw, &replyStruct)
	is.Equal(testData, replyStruct)

	// is.Equal(testData, replyStruct)

	// add unique index for F1
	unique := true
	col.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{{"f1", 1}},
		Options: &options.IndexOptions{
			Unique: &unique,
		},
	})

	// inserting same doc again
	err = store.Store(context.Background(), testData)
	is.Error(err)
	is.Equal(errors.CodeAlreadyExists, err.GetStatus())
}

func TestMongoGet(t *testing.T) {
	col := getTestMongoCollection()
	defer cleanUP(col)
	store := &mongoDS{
		db: col,
	}
	is := assert.New(t)

	// get non-existing item
	_, err := store.Get(context.Background(), "f1", "f1")
	is.Equal(errors.CodeNotFound, err.GetStatus())

	// store item
	testData := testStruct{
		F1: "f1",
		F2: 23,
		F3: 65.32,
		F4: time.Now().Unix(),
	}
	store.db.InsertOne(context.Background(), testData)
	raw, err := store.Get(context.Background(), "f1", "f1")
	is.NoError(err)
	var replyStruct testStruct
	bson.Unmarshal(raw, &replyStruct)
	is.Equal(testData, replyStruct)
}

func TestMongoUpdate(t *testing.T) {
	col := getTestMongoCollection()
	defer cleanUP(col)
	store := &mongoDS{
		db: col,
	}
	is := assert.New(t)
	testData := testStruct{
		F1: "f1",
		F2: 23,
		F3: 65.32,
		F4: time.Now().Unix(),
	}
	// update non-existing item
	err := store.Update(context.Background(), "f1", testData.F1, testData)
	is.Equal(errors.CodeNotFound, err.GetStatus())

	// store testData
	store.db.InsertOne(context.Background(), testData)
	// update testData
	testData.F2 = 24
	testData.F3 = 66
	err = store.Update(context.Background(), "f1", testData.F1, testData)
	var replyData testStruct
	col.FindOne(context.Background(), bson.M{"f1": testData.F1}).Decode(&replyData)
	is.Equal(testData, replyData)
}

func TestMongoDelete(t *testing.T) {
	col := getTestMongoCollection()
	defer cleanUP(col)
	store := &mongoDS{
		db: col,
	}
	is := assert.New(t)

	// delete non existing item
	err := store.Delete(context.Background(), "f1", "f1")
	is.Equal(errors.CodeNotFound, err.GetStatus())

	// load some item
	testData := testStruct{
		F1: "f1",
		F2: 23,
		F3: 65.32,
		F4: time.Now().Unix(),
	}
	store.db.InsertOne(context.Background(), testData)
	err = store.Delete(context.Background(), "f1", testData.F1)
	is.NoError(err)
	reply := col.FindOne(context.Background(), bson.M{"f1": testData.F1})
	is.Equal(mongo.ErrNoDocuments, reply.Err())
}

func TestMongoQuery(t *testing.T) {
	col := getTestMongoCollection()
	defer cleanUP(col)
	store := &mongoDS{
		db: col,
	}
	is := assert.New(t)
	ctx := context.TODO()

	// load some items
	number := time.Now().Unix()
	size := 10
	docs := make([]interface{}, size)
	for i := 0; i < size; i++ {
		docs[i] = testStruct{
			F1: strconv.Itoa(i),
			F2: i % size,
			F3: 2.6,
			F4: number + int64(i),
		}
	}
	col.InsertMany(ctx, docs)
	query := map[string]interface{}{
		"f4": map[string]interface{}{
			"$gte": number,
		},
	}
	pageNumber := 1
	perPage := size - 3
	sortingKey := "f1"
	raws, err := store.Query(ctx, sortingKey, query, int64(pageNumber), int64(perPage))
	is.NoError(err)
	is.Equal(perPage, len(raws))
	for _, raw := range raws {
		var data testStruct
		bson.Unmarshal(raw, &data)
		num, _ := strconv.Atoi(data.F1)
		is.Equal(docs[num], data)
	}
	pageNumber++
	raws, err = store.Query(ctx, sortingKey, query, int64(pageNumber), int64(perPage))
	is.NoError(err)
	is.Equal(size-perPage, len(raws))
	for _, raw := range raws {
		var data testStruct
		bson.Unmarshal(raw, &data)
		num, _ := strconv.Atoi(data.F1)
		is.Equal(docs[num], data)
	}
}

func TestMongoNew(t *testing.T) {
	conf := config.DatastoreConf{
		URL:        "localhost:27018",
		Username:   "root",
		Password:   "password",
		Database:   testMongoDatabase,
		Collection: testMongoCollection,
	}
	ds, err := NewSmartDS(conf)
	assert.NoError(t, err)
	assert.NotNil(t, ds)
	cleanUP(getTestMongoCollection())
}

////

const (
	testMongoDatabase   = "testdb"
	testMongoCollection = "testcollection"
)

func getTestMongoCollection() *mongo.Collection {
	addrs := "mongodb://root:password@localhost:27018"
	client, err := mongo.NewClient(options.Client().ApplyURI(addrs))
	if err != nil {
		log.Println("failed to create client", err)
		os.Exit(1)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Println("failed to connect mongo", err)
		os.Exit(1)
	}
	return client.Database(testMongoDatabase).Collection(testMongoCollection)
}

func cleanUP(collection *mongo.Collection) {
	collection.Database().Drop(context.TODO())
	collection.Database().Client().Disconnect(context.Background())
}
