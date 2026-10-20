package fiberpkg

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func InitRecover() fiber.Handler {
	return recover.New()
}
