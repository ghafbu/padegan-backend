package fiberpkg

import (
	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/requestid"
)

func InitRequestID() fiber.Handler {
	return requestid.New(requestid.Config{
		Generator: func() string {
			return ghafpkg.GenUUID()
		},
	})
}
