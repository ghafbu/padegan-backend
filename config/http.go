package config

type HTTP struct {
	Url       string `env:"HTTP_URL"`
	Host      string `env:"HTTP_HOST"`
	PortHttp  string `env:"HTTP_PORT_HTTP"`
	PortHttps string `env:"HTTP_PORT_HTTPS"`
	Protocol  string `env:"HTTP_PROTOCOL"`
}
