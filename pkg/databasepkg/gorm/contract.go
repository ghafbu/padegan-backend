package gorm

import (
	"gorm.io/gorm"
)

type DB interface {
	Conn() *gorm.DB
	Ping() error
	Close() error
}

type Migrator interface {
	LoadEntities(entities ...interface{}) Migrator
	AutoMigrate() Migrator
	Up() Migrator
	Down() Migrator
}

type Faker interface {
	Create(entity interface{}) (Faker, error)
	Truncate(tableName string) error
	TruncateWithModel(modelName interface{}) error
	TruncateWithModels(modelNames ...interface{}) error
	FirstOrCreate(entity interface{}, conds ...interface{}) (Faker, error)
}
