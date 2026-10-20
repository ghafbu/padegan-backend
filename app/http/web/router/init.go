package router

import (
	"github.com/ghafbu/padegan-backend/app/http/web/controller"
	internalmiddleware "github.com/ghafbu/padegan-backend/app/http/web/middleware"
	"github.com/ghafbu/padegan-backend/lib/bootstrap"
	"github.com/gofiber/fiber/v3"
)

type Router struct {
	app        *bootstrap.Application
	middleware *internalmiddleware.Middleware
	controller *controller.Controller
	fiberApp   *fiber.App
}

func Init(
	paramApp *bootstrap.Application,
	//paramRateLimiter globalinterface.IRateLimiterStorage,
	//paramToken globalinterface.IToken,
	paramMiddleware *internalmiddleware.Middleware,
	paramController *controller.Controller,
	paramFiberApp *fiber.App,
) *Router {
	return &Router{
		app:        paramApp,
		middleware: paramMiddleware,
		controller: paramController,
		fiberApp:   paramFiberApp,
	}
}

func (receiverRouter *Router) Register() {
	// Routers Prefix
	thisRouterPrefix := receiverRouter.fiberApp.Group("")

	// Global Middleware
	receiverRouter.middleware.GlobalRegister(thisRouterPrefix)

	// Initialize Routers
	receiverRouter.routers(thisRouterPrefix)
}
