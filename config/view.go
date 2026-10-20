package config

type ViewInfo struct {
	Path      string `env:"VIEW_PATH"`
	Extension string `env:"VIEW_EXTENSION"`
}
