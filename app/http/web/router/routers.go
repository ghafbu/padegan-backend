package router

import (
	"github.com/ghafbu/padegan-backend/lib/middleware"
	"github.com/gofiber/fiber/v3"
)

func (receiverRouter *Router) routers(paramRouter fiber.Router) {
	// All Public APIs
	//thisPublicRouter := paramRouter.Use()
	receiverRouter.accessPublic(paramRouter)

	// All Private APIs For User
	//thisProtectedUserRouter := paramRouter.Use()
	//receiverRouter.accessUser(thisProtectedUserRouter)

	// All Private APIs For Admin
	//thisProtectedAdminRouter := paramRouter.Use()
	//receiverRouter.accessAdmin(thisProtectedAdminRouter)

	// ***************** Other ***************

	// For Not Found Page (NoExists) - 404
	paramRouter.Use(middleware.FiberRenderNotFoundPage("errors/404"))
}
