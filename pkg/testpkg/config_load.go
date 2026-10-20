package testpkg

import (
	"fmt"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

func GetConfig[ConfigType interface{}](paths ...string) (ConfigType, error) {
	var Cfg ConfigType

	for _, v := range paths {
		if err := godotenv.Load(v); err != nil {
			return Cfg, fmt.Errorf("error loading global file '%s': %w", v, err)
		}

		if err := env.Parse(&Cfg); err != nil {
			return Cfg, fmt.Errorf("error parsing global environment variables into struct: %w", err)
		}
	}

	return Cfg, nil
}
