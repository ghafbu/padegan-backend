package giiiratelimiterstorage

import (
	"time"

	"github.com/ghafbu/padegan-backend/config"

	_interface "github.com/ghafbu/padegan-backend/lib/interface"
	"github.com/ghafbu/padegan-backend/lib/interfaceimplement/giiiratelimiterstorage/myredis"
	redispkg "github.com/ghafbu/padegan-backend/pkg/cachepkg/redis"
)

const (
	ImplementNameRedis = "myRedis"
)

func NewIRedis(
	paramImplementName string,
	config config.Redis,
	timeout time.Duration,
) (_interface.IRateLimiterStorage, error) {
	switch paramImplementName {
	case ImplementNameRedis:
		{
			conn, err := redispkg.Init(config)
			if err != nil {
				return nil, err
			}

			return myredis.Init(conn, timeout)
		}
	}

	return nil, nil
}
