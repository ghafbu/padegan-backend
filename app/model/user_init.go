package model

import (
	"errors"

	"github.com/ghafbu/padegan-backend/config"
	"github.com/ghafbu/padegan-backend/lib/consts"
	_interface "github.com/ghafbu/padegan-backend/lib/interface"
	"github.com/ghafbu/padegan-backend/pkg/databasepkg/gorm"
)

type UserModel struct {
	config   config.Config
	logger   _interface.ILogger
	postgres gorm.DB
}

func InitUsers(
	paramConfig config.Config,
	paramLogger _interface.ILogger,
	paramPostgres gorm.DB,
) *UserModel {
	return &UserModel{
		config:   paramConfig,
		logger:   paramLogger,
		postgres: paramPostgres,
	}
}

func (myModel *UserModel) checkConnection() error {
	if myModel == nil || myModel.postgres == nil {
		return errors.New(consts.ErrorKeyModelDatabaseConnectionInit)
	}

	return nil
}
