package middleware

import (
	"github.com/ghafbu/padegan-backend/app/model"
	"github.com/ghafbu/padegan-backend/lib/bootstrap"
	"github.com/ghafbu/padegan-backend/pkg/fiberpkg"
	"github.com/gofiber/fiber/v3"
)

type Middleware struct {
	app   *bootstrap.Application
	model *model.UserModel
}

func Init(
	paramApp *bootstrap.Application,
	paramModel *model.UserModel,
) *Middleware {
	return &Middleware{
		app:   paramApp,
		model: paramModel,
	}
}

func (receiverMiddleware Middleware) GlobalRegister(paramRouter fiber.Router) {
	// Global Middleware For All Routers
	paramRouter.Use(
		fiberpkg.InitSession(),
		//receiverMiddleware.NotFoundPage(),
	)
}
