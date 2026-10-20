package router

import (
	"github.com/gofiber/fiber/v3"
)

func (receiverRouter *Router) accessAdmin(paramRouter fiber.Router) {
	{
		paramRouter.Get("/", func(c fiber.Ctx) error {
			return c.SendString("-- Access Admin --")
		})
	}
}
