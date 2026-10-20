package mongodbpkg

type Config struct {
	Host     string `env:"MONGO_HOST"`
	Port     string `env:"MONGO_PORT"`
	User     string `env:"MONGO_USER"`
	Password string `env:"MONGO_PASSWORD"`
	DB       string `env:"MONGO_DATABASE"`
	Url      string `env:"MONGODB_URI"`
	//SSLMode         string        `env:"POSTGRES_SSL_MODE"`
	//MaxIdleConns    int           `env:"POSTGRES_MAX_IDLE_CONNS"`
	//MaxOpenConns    int           `env:"POSTGRES_MAX_OPEN_CONNS"`
	//ConnMaxLifetime time.Duration `env:"POSTGRES_CONN_MAX_LIFETIME"`
	//LogLevel        int           `env:"POSTGRES_LOG_LEVEL"`
}
