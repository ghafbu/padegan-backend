package migrator

import (
	"fmt"

	"github.com/ghafbu/padegan-backend/pkg/databasepkg/gorm"
)

type Migrator struct {
	conn     gorm.DB
	entities []interface{}
}

func Init(conn gorm.DB) *Migrator {
	return &Migrator{
		conn:     conn,
		entities: make([]interface{}, 0),
	}
}

func (m *Migrator) LoadEntities(entities ...interface{}) gorm.Migrator {
	m.entities = append(m.entities, entities...)

	return m
}

func (m *Migrator) AutoMigrate() gorm.Migrator {
	if len(m.entities) == 0 {
		return nil
	}

	for _, entity := range m.entities {
		_ = m.conn.Conn().AutoMigrate(entity)
		// log.Println(m.conn.Conn().AutoMigrate(entity))
	}

	fmt.Print("Auto migrate completed\n")

	return m
}

func (m *Migrator) Up() gorm.Migrator {
	for _, entity := range m.entities {
		if !m.conn.Conn().Migrator().HasTable(entity) {
			_ = m.conn.Conn().Migrator().CreateTable(entity)
		}
	}
	fmt.Print("Up tables completed\n")

	return m
}

func (m *Migrator) Down() gorm.Migrator {
	for _, entity := range m.entities {
		_ = m.conn.Conn().Migrator().DropTable(entity)
	}

	fmt.Print("Down tables completed\n")

	return m
}
