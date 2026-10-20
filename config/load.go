package config

import (
	"fmt"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

// Load loads environment variables from a list of file paths and binds them to a Config struct.
func Load(cfgPath ...string) (Config, error) {
	var cfg Config

	// Iterate over each path in cfgPath to load the environment variables from each file
	if err := godotenv.Overload(cfgPath...); err != nil {
		// Return an error with context if unable to load a file
		return cfg, fmt.Errorf("error loading file '%s': %w", cfgPath, err)
	}

	// Parse loaded environment variables into the App struct
	if err := env.Parse(&cfg); err != nil {
		// Log and return error if binding to the struct fails
		return cfg, fmt.Errorf("error parsing environment variables into struct: %w", err)
	}

	return cfg, nil
}
