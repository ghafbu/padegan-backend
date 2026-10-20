package myredis

import (
	redispkg "github.com/ghafbu/padegan-backend/pkg/cachepkg/redis"
)

type MyRedis struct {
	conn *redispkg.Redis
}

func Init(client *redispkg.Redis) (*MyRedis, error) {
	return &MyRedis{conn: client}, nil
}
