package config

import "time"

type ExpireDuration struct {
	In1m time.Duration `env:"EXPIRE_DURATION_1M"`
	In2m time.Duration `env:"EXPIRE_DURATION_2M"`
	In3m time.Duration `env:"EXPIRE_DURATION_3M"`
	In4m time.Duration `env:"EXPIRE_DURATION_4M"`
	In8m time.Duration `env:"EXPIRE_DURATION_8M"`
	In1w time.Duration `env:"EXPIRE_DURATION_1W"`
	In1y time.Duration `env:"EXPIRE_DURATION_1Y"`
}
