package postgres

import (
	"fmt"

	"github.com/ghafbu/padegan-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	gorm2 "github.com/ghafbu/padegan-backend/pkg/databasepkg/gorm"
)

type Postgres struct {
	//config Config
	config config.Postgres
	conn   *gorm.DB
	// db     *sql.DB
}

// func Init(config Config) (gorm2.DB, error) {
func Init(config config.Postgres) (gorm2.DB, error) {
	cnn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
		config.Host, config.Port, config.User, config.Password,
		config.DB, config.SSLMode)

	conn, err := gorm.Open(postgres.Open(cnn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(config.LogLevel)),
	})
	if err != nil {
		return nil, fmt.Errorf("can`t open psql conn: %w", err)
	}

	db, _ := conn.DB()
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to obtain the underlying SQL database instance: %w", err)
	// }

	// See "Important config" section
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetConnMaxLifetime(config.ConnMaxLifetime)

	return &Postgres{config: config, conn: conn}, nil
}

func (p *Postgres) Ping() error {
	db, err := p.conn.DB()
	if err != nil {
		return fmt.Errorf("failed to obtain the underlying SQL database instance: %w", err)
	}

	return db.Ping()
}

func (p *Postgres) Conn() *gorm.DB {
	return p.conn
}

func (p *Postgres) Close() error {
	db, err := p.conn.DB()
	if err != nil {
		return fmt.Errorf("failed to close postgres SQL database instance: %w", err)
	}

	return db.Close()
}
