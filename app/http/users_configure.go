package http

import (
	"github.com/ghafbu/padegan-backend/app/model"
	"github.com/ghafbu/padegan-backend/app/usecase"

	apicontroller "github.com/ghafbu/padegan-backend/app/http/api/controller"
	apimiddleware "github.com/ghafbu/padegan-backend/app/http/api/middleware"
	apirouter "github.com/ghafbu/padegan-backend/app/http/api/router"
	webcontroller "github.com/ghafbu/padegan-backend/app/http/web/controller"
	webmiddleware "github.com/ghafbu/padegan-backend/app/http/web/middleware"
	webrouter "github.com/ghafbu/padegan-backend/app/http/web/router"
)

func (s Server) UsersConfigure() {
	// *************** Dependencies Section ***************
	thisUsersModel := model.InitUsers(s.app.Config, s.app.Logger, s.app.Postgres)
	thisUserSettingsModel := model.InitUserSettings(s.app.Config, s.app.Logger, s.app.Postgres)

	thisUserSettingsUseCase := usecase.InitUserSettings(
		s.app.Config,
		s.app.Logger,
		s.app.Token,
		s.app.Redis,
		thisUserSettingsModel,
	)

	thisUsersUseCase := usecase.InitUsers(
		s.app.Config,
		s.app.Logger,
		s.app.Token,
		s.app.Redis,
		thisUsersModel,
	)

	// API Section
	thisApiMiddleware := apimiddleware.Init(s.app, thisUsersModel)

	thisApiController := apicontroller.InitUsers(s.app, thisUsersModel, thisUsersUseCase)
	thisApiUserSettingsController := apicontroller.InitUserSettings(s.app, thisUserSettingsModel, thisUserSettingsUseCase)

	apirouter.Init(
		s.app,
		thisApiMiddleware,
		thisApiController,
		thisApiUserSettingsController,
		s.fiber,
	).
		Register()

	// Web Section
	thisWebMiddleware := webmiddleware.Init(s.app, thisUsersModel)

	thisWebController := webcontroller.Init(s.app, thisUsersModel, thisUsersUseCase)

	webrouter.Init(s.app, thisWebMiddleware, thisWebController, s.fiber).
		Register()
}
