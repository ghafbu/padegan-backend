package giiiredis

import (
	"github.com/ghafbu/padegan-backend/config"
	_interface "github.com/ghafbu/padegan-backend/lib/interface"
	"github.com/ghafbu/padegan-backend/lib/interfaceimplement/giiiredis/myredis"
	redispkg "github.com/ghafbu/padegan-backend/pkg/cachepkg/redis"
)

const (
	ImplementNameRedis = "myRedis"
)

func NewIRedis(paramImplementName string, config config.Redis) (_interface.IRedis, error) {
	switch paramImplementName {
	case ImplementNameRedis:
		{
			conn, err := redispkg.Init(config)
			if err != nil {
				return nil, err
			}

			return myredis.Init(conn)
		}
	}

	return nil, nil
}
