package myredis

import "time"

type Config struct {
	Host               string        `env:"REDIS_HOST"`
	Port               string        `env:"REDIS_PORT"`
	Password           string        `env:"REDIS_PASSWORD"`
	DB                 int           `env:"REDIS_DB"`
	DialTimeout        time.Duration `env:"REDIS_DIAL_TIMEOUT"`
	ReadTimeout        time.Duration `env:"REDIS_READ_TIMEOUT"`
	WriteTimeout       time.Duration `env:"REDIS_WRITE_TIMEOUT"`
	PoolSize           int           `env:"REDIS_POOL_SIZE"`
	PoolTimeout        time.Duration `env:"REDIS_POOL_TIMEOUT"`
	IdleCheckFrequency time.Duration `env:"REDIS_IDLE_CHECK_FREQUENCY"`
}
