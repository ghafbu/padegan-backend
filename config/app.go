package config

type Application struct {
	Name    string `env:"APPLICATION_NAME"`
	Version string `env:"APPLICATION_VERSION"`
	Debug   bool   `env:"APPLICATION_DEBUG"`
}
