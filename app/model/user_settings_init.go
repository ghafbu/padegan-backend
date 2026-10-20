package model

import (
	"errors"

	"github.com/ghafbu/padegan-backend/config"
	"github.com/ghafbu/padegan-backend/lib/consts"
	_interface "github.com/ghafbu/padegan-backend/lib/interface"
	"github.com/ghafbu/padegan-backend/pkg/databasepkg/gorm"
)

type UserSettingModel struct {
	config   config.Config
	logger   _interface.ILogger
	postgres gorm.DB
}

func InitUserSettings(
	paramConfig config.Config,
	paramLogger _interface.ILogger,
	paramPostgres gorm.DB,
) *UserSettingModel {
	return &UserSettingModel{
		config:   paramConfig,
		logger:   paramLogger,
		postgres: paramPostgres,
	}
}

func (myModel *UserSettingModel) checkConnection() error {
	if myModel == nil || myModel.postgres == nil {
		return errors.New(consts.ErrorKeyModelDatabaseConnectionInit)
	}

	return nil
}
