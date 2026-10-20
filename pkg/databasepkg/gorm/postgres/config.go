package postgres

import "time"

type Config struct {
	Host            string        `env:"POSTGRES_HOST"`
	Port            string        `env:"POSTGRES_PORT"`
	User            string        `env:"POSTGRES_USER"`
	Password        string        `env:"POSTGRES_PASSWORD"`
	DB              string        `env:"POSTGRES_DB"`
	SSLMode         string        `env:"POSTGRES_SSL_MODE"`
	MaxIdleConns    int           `env:"POSTGRES_MAX_IDLE_CONNS"`
	MaxOpenConns    int           `env:"POSTGRES_MAX_OPEN_CONNS"`
	ConnMaxLifetime time.Duration `env:"POSTGRES_CONN_MAX_LIFETIME"`
	LogLevel        int           `env:"POSTGRES_LOG_LEVEL"`
}
