package usecase

import (
	"github.com/ghafbu/padegan-backend/config"
	_interface "github.com/ghafbu/padegan-backend/lib/interface"
)

type UserSettingsUseCase struct {
	config config.Config
	logger _interface.ILogger
	token  _interface.IToken
	redis  _interface.IRedis
	model  _interface.IUserSettingsModel
}

func InitUserSettings(
	paramConfig config.Config,
	paramLogger _interface.ILogger,
	paramToken _interface.IToken,
	paramRedis _interface.IRedis,
	paramModel _interface.IUserSettingsModel,
) *UserSettingsUseCase {
	return &UserSettingsUseCase{
		config: paramConfig,
		logger: paramLogger,
		token:  paramToken,
		redis:  paramRedis,
		model:  paramModel,
	}
}
