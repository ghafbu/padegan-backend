package controller

import (
	"github.com/gofiber/fiber/v3"
)

func (receiverController *Controller) Welcome(paramCtx fiber.Ctx) error {
	return paramCtx.Render("welcome", fiber.Map{
		"Title": "Welcome",
	}, "layout/main")
}
