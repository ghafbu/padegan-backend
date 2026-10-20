package controller

import (
	"github.com/ghafbu/padegan-backend/app/model"
	"github.com/ghafbu/padegan-backend/app/usecase"
	"github.com/ghafbu/padegan-backend/lib/bootstrap"
)

type UsersController struct {
	app     *bootstrap.Application
	model   *model.UserModel
	useCase *usecase.UsersUseCase
}

func InitUsers(
	paramApp *bootstrap.Application,
	paramModel *model.UserModel,
	paramUseCase *usecase.UsersUseCase,
) *UsersController {
	return &UsersController{
		app:     paramApp,
		model:   paramModel,
		useCase: paramUseCase,
	}
}
