package router

import (
	"github.com/ghafbu/padegan-backend/app/http/api/controller"
	"github.com/ghafbu/padegan-backend/app/http/api/middleware"
	"github.com/ghafbu/padegan-backend/lib/bootstrap"
	"github.com/gofiber/fiber/v3"
)

type Router struct {
	app                    *bootstrap.Application
	middleware             *middleware.Middleware
	usersController        *controller.UsersController
	userSettingsController *controller.UserSettingsController
	fiberApp               *fiber.App
}

func Init(
	paramApp *bootstrap.Application,
	paramMiddleware *middleware.Middleware,
	paramUsersController *controller.UsersController,
	paramUserSettingsController *controller.UserSettingsController,
	paramFiberApp *fiber.App,
) *Router {
	return &Router{
		app:                    paramApp,
		middleware:             paramMiddleware,
		usersController:        paramUsersController,
		userSettingsController: paramUserSettingsController,
		fiberApp:               paramFiberApp,
	}
}

func (receiverRouter Router) Register() {
	// Routers Prefix
	thisRouterPrefix := receiverRouter.fiberApp.Group("/v1/api")

	// Global Middleware
	receiverRouter.middleware.GlobalRegister(thisRouterPrefix)

	// Initialize Routers
	receiverRouter.routers(thisRouterPrefix)
}
