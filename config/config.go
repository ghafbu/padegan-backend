package config

var Cfg = Config{}

type Config struct {
	Application    Application
	OTP            OTP
	HTTP           HTTP
	Postgres       Postgres
	Redis          Redis
	Logger         Logger
	SecretKey      ApplicationSecretKey
	ViewInfo       ViewInfo
	StaticFile     StaticFile
	Other          Other
	ExpireDuration ExpireDuration
}
