package datastore

import (
	"context"

	"github.com/Zzocker/blab/pkg/errors"
	"github.com/gomodule/redigo/redis"
)

type redisStore struct {
	p *redis.Pool
}

func (r *redisStore) Store(ctx context.Context, key string, value []byte, expireIn int64) errors.E {
	conn, err := r.p.GetContext(ctx)
	if err != nil {
		return errors.New(errors.CodeInternalErr, "failed to get connection from pool") // Log internal server error with Warn
	}
	defer conn.Close()
	_, err = conn.Do("SET", key, value, "EX", expireIn)
	if err != nil {
		return errors.New(errors.CodeInternalErr, "failed to store")
	}
	return nil
}
func (r *redisStore) Get(ctx context.Context, key string) ([]byte, errors.E) {
	conn, err := r.p.GetContext(ctx)
	if err != nil {
		return nil, errors.New(errors.CodeInternalErr, "failed to get connection from pool")
	}
	defer conn.Close()
	raw, err := redis.Bytes(conn.Do("GET", key))
	if err == redis.ErrNil {
		return nil, errors.New(errors.CodeNotFound, "item not found")
	} else if err != nil {
		return nil, errors.New(errors.CodeInternalErr, "failed to get item")
	}
	return raw, nil
}
func (r *redisStore) Delete(ctx context.Context, key string) errors.E {
	conn, err := r.p.GetContext(ctx)
	if err != nil {
		return errors.New(errors.CodeInternalErr, "failed to get connection from pool")
	}
	defer conn.Close()
	reply, err := redis.Int(conn.Do("DEL", key))
	if err != nil {
		return errors.New(errors.CodeInternalErr, "failed to get connection from pool")
	}
	if reply != 1 {
		return errors.New(errors.CodeNotFound, "item not found")
	}
	return nil
}
