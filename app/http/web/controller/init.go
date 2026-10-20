package controller

import (
	"github.com/ghafbu/padegan-backend/app/model"
	"github.com/ghafbu/padegan-backend/app/usecase"
	"github.com/ghafbu/padegan-backend/lib/bootstrap"
)

type Controller struct {
	app     *bootstrap.Application
	model   *model.UserModel
	useCase *usecase.UsersUseCase
}

func Init(
	paramApp *bootstrap.Application,
	paramModel *model.UserModel,
	paramUseCase *usecase.UsersUseCase,
) *Controller {
	return &Controller{
		app:     paramApp,
		model:   paramModel,
		useCase: paramUseCase,
	}
}
