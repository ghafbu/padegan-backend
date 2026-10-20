package controller

import (
	"github.com/ghafbu/padegan-backend/app/model"
	"github.com/ghafbu/padegan-backend/app/usecase"
	"github.com/ghafbu/padegan-backend/lib/bootstrap"
)

type UserSettingsController struct {
	app     *bootstrap.Application
	model   *model.UserSettingModel
	useCase *usecase.UserSettingsUseCase
}

func InitUserSettings(
	paramApp *bootstrap.Application,
	paramModel *model.UserSettingModel,
	paramUseCase *usecase.UserSettingsUseCase,
) *UserSettingsController {
	return &UserSettingsController{
		app:     paramApp,
		model:   paramModel,
		useCase: paramUseCase,
	}
}
