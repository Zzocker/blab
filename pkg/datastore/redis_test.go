package datastore

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/pkg/errors"
	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
)

func TestDumbType(t *testing.T) {
	var l interface{} = &redisStore{}
	_, ok := l.(DumbDS)
	assert.True(t, ok)
}

func TestNewDumbDS(t *testing.T) {
	conf := config.DatastoreConf{
		URL: "localhost:6378",
	}
	_, err := NewDumbDS(conf)
	assert.NoError(t, err)
}

func TestRedisStore(t *testing.T) {
	is := assert.New(t)
	store := redisStore{p: testRedisPool()}
	conn := store.p.Get()
	defer conn.Close()

	testKey := "store"
	testValue := "value"
	// happy flow
	err := store.Store(context.Background(), testKey, []byte(testValue), int64(time.Second.Seconds()))
	is.NoError(err)
	value, _ := redis.Bytes(conn.Do("Get", testKey))
	is.Equal(testValue, string(value))
	time.Sleep(time.Second)
	value, _ = redis.Bytes(conn.Do("Get", testKey))
	is.Zero(value)
}

func TestRedisGet(t *testing.T) {
	is := assert.New(t)
	store := redisStore{p: testRedisPool()}
	conn := store.p.Get()
	defer conn.Close()

	getKey := "get"
	value := "value"
	// store in redis
	conn.Do("SET", getKey, []byte(value))

	// run test on get
	raw, err := store.Get(context.Background(), getKey)
	is.NoError(err)
	is.Equal(value, string(raw))

	// get non-existing item
	raw, err = store.Get(context.Background(), "notFound")
	is.Error(err)
	is.Zero(raw)
	is.Equal(errors.CodeNotFound, err.GetStatus())
}

func TestRedisDelete(t *testing.T) {
	is := assert.New(t)
	store := redisStore{p: testRedisPool()}
	conn := store.p.Get()
	defer conn.Close()

	delKey := "del"
	value := "value"
	// store in redis
	conn.Do("SET", delKey, []byte(value))

	err := store.Delete(context.Background(), "noFound")
	is.Error(err)
	is.Equal(errors.CodeNotFound, err.GetStatus())

	err = store.Delete(context.Background(), delKey)
	is.NoError(err)
	_, rErr := redis.Bytes(conn.Do("GET", delKey))
	is.Error(rErr)
	is.Equal(redis.ErrNil, rErr)
}

//
func testRedisPool() *redis.Pool {
	pool := redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial("tcp", ":6378")
	}, 0)
	conn := pool.Get()
	defer conn.Close()
	if _, err := conn.Do("Ping"); err != nil {
		os.Exit(1)
	}
	return pool
}
