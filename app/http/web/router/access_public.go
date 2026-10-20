package router

import (
	"github.com/ghafbu/padegan-backend/lib/consts"
	"github.com/gofiber/fiber/v3"
)

func (receiverRouter *Router) accessPublic(paramRouter fiber.Router) {
	paramRouter.Get("/", receiverRouter.controller.Welcome).
		Name(consts.Welcome)

	_ = paramRouter.Group("/auth")
	{
		// Login

		// Verify
	}

	_ = paramRouter.Group("/reset-password")
	{
		// Reset Password

		// Reset Password Verify

		// Reset Password Update
	}
}
