package config

import "time"

type Other struct {
	PasswordEntropy            float64       `env:"SECURITY_PASSWORD_ENTROPY"`
	CMDGracefulShutDownTimeOut time.Duration `env:"CMD_GRACEFUL_SHUTDOWN_TIMEOUT"`
}
