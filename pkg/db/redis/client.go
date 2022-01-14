package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func NewRedisConnection(host, port, password string) (*redis.Client, error) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
	})
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return rdb, nil
}
