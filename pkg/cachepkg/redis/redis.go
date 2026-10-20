package redis

import (
	"context"
	"fmt"

	"github.com/ghafbu/padegan-backend/config"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	cfg config.Redis
	db  *redis.Client
}

func Init(cfg config.Redis) (*Redis, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	// Ping Redis to ensure connection is successful
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis at %s:%s - %v", cfg.Host, cfg.Port, err)
	}

	return &Redis{db: rdb}, nil
}

func (r *Redis) Client() *redis.Client {
	return r.db
}
