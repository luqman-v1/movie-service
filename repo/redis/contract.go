package redis

import (
	"context"
	"time"
)

var (
	GetKey          = "movie-service:imdb-id:%v"
	DefaultDuration = 2 * time.Minute
)

type IRedisDb interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, value interface{}, duration time.Duration) error
	Del(ctx context.Context, key string) error
}

type RedisDB struct {
}

func (r *RedisDB) Get(ctx context.Context, key string) ([]byte, error) {
	return RedisClient.Get(ctx, key).Bytes()
}

func (r *RedisDB) Set(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	return RedisClient.Set(ctx, key, value, duration).Err()
}

func (r *RedisDB) Del(ctx context.Context, key string) error {
	return RedisClient.Del(ctx, key).Err()
}

func NewRedis() IRedisDb {
	return &RedisDB{}
}
