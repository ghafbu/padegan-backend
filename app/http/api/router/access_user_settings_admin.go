package router

import (
	"github.com/ghafbu/padegan-backend/api/dto/request"
	"github.com/ghafbu/padegan-backend/config"
	"github.com/ghafbu/padegan-backend/lib/consts"
	"github.com/ghafbu/padegan-backend/lib/middleware"
	"github.com/ghafbu/padegan-backend/pkg/fiberpkg/fiberpkgrequest"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
)

func (receiverRouter Router) accessUserSettingsAdmin(paramRouter fiber.Router) {
	// Set middleware for protected routes
	//paramRouter.Use(
	//	middleware.FiberAdminAuth(
	//		receiverRouter.app.Config.SecretKey.JwtAuthAccess,
	//		consts.DefaultKeyAdminJWTAccess,
	//		receiverRouter.app.Token,
	//		receiverRouter.app.Redis,
	//	))
	{
		// Get
		paramRouter.Get(
			"/",
			limiter.New(middleware.RateLimiter(config.RateLimiterConfig{
				Max:        consts.RateLimiterGeneralLimit,
				Expiration: consts.RateLimiterExpiration,
			}, receiverRouter.app.RateLimiterStorage)),
			receiverRouter.userSettingsController.UserSettingsGet,
		)

		// Update
		paramRouter.Patch(
			"/",
			limiter.New(middleware.RateLimiter(config.RateLimiterConfig{
				Max:        consts.RateLimiterGeneralLimit,
				Expiration: consts.RateLimiterExpiration,
			}, receiverRouter.app.RateLimiterStorage)),
			fiberpkgrequest.Middleware(&request.UserSettingsUpdate{}, fiberpkgrequest.JSON),
			receiverRouter.userSettingsController.UserSettingsUpdate,
		)

	}
}
