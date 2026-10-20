package config

type Logger struct {
	FilePath string `env:"LOGGER_FILEPATH"`
	Encoding string `env:"LOGGER_ENCODING"`
	Level    string `env:"LOGGER_LEVEL"`
	Logger   string `env:"LOGGER_LOGGER"`
}
