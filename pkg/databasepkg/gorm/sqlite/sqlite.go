package sqlite

import (
	"gorm.io/gorm"
)

type SqliteDB struct {
	config Config
	db     *gorm.DB
}

// func Init(config Config) gorm2.DB {
// db, err := orm.Open(sqliteorm.Open("orm.db"), &orm.App{})
// "orm.io/driver/sqliteorm" // Sqlite driver based on CGO & Conn: db

// return &SqliteDB{config: config}
// }

// func (s *SqliteDB) Conn() *gorm.DB {
// 	return s.db
// }
