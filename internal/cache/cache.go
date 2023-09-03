package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type CacheHandler struct {
	client *redis.Client
	ctx    context.Context
}

func NewCacheHandler(addr, password string, db int) (*CacheHandler, error) {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &CacheHandler{client: client, ctx: ctx}, nil
}

func (ch *CacheHandler) GetValue(key string) (string, error) {
	return ch.client.Get(ch.ctx, key).Result()
}

func (ch *CacheHandler) SetValue(key, value string, expiration time.Duration) error {
	return ch.client.Set(ch.ctx, key, value, expiration).Err()
}

func (ch *CacheHandler) Close() {
	ch.client.Close()
}
