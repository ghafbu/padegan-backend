package middleware

import (
	"github.com/gofiber/fiber/v3"
)

func FiberRenderNotFoundPage(paramPathView string) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		return ctx.Status(fiber.StatusNotFound).Render(paramPathView, nil, "layout/main")
	}
}
