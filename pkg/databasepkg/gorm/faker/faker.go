package faker

import (
	"fmt"

	"github.com/ghafbu/padegan-backend/pkg/databasepkg/gorm"
)

type Faker struct {
	conn gorm.DB
}

func Init(conn gorm.DB) *Faker {
	return &Faker{
		conn: conn,
	}
}

func (m *Faker) Create(entity interface{}) (gorm.Faker, error) {
	res := m.conn.Conn().Create(entity)

	if res.Error != nil {
		return nil, fmt.Errorf("can`t faker create, %v", res.Error)
	}

	return m, nil
}

func (m *Faker) Truncate(tableName string) error {
	return m.conn.Conn().Exec(fmt.Sprintf("TRUNCATE TABLE %s", tableName)).Error
}

func (m *Faker) TruncateWithModel(modelName interface{}) error {
	return m.conn.Conn().Unscoped().Where("1 = 1").Delete(modelName).Error
}

func (m *Faker) TruncateWithModels(modelNames ...interface{}) error {
	for _, modelName := range modelNames {
		return m.conn.Conn().Unscoped().Where("1 = 1").Delete(modelName).Error
	}

	return nil
}

func (m *Faker) FirstOrCreate(entity interface{}, conds ...interface{}) (gorm.Faker, error) {
	res := m.conn.
		Conn().
		FirstOrCreate(entity, conds)

	if res.Error != nil {
		return nil, fmt.Errorf("can`t faker create, %v", res.Error)
	}

	return m, nil
}
