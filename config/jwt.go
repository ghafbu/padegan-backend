package config

type ApplicationSecretKey struct {
	Aes              string `env:"SECRET_KEY_AES"`
	JwtPublic        string `env:"SECRET_KEY_JWT_PUBLIC"`
	JwtAuthAccess    string `env:"SECRET_KEY_JWT_AUTH_ACCESS"`
	JwtAuthRefresh   string `env:"SECRET_KEY_JWT_AUTH_REFRESH"`
	JwtResetPassword string `env:"SECRET_KEY_JWT_RESET_PASSWORD"`
}
