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

func (receiverRouter Router) accessUsersPublic(paramRouter fiber.Router) {
	authRoutes := paramRouter.Group("/auth")
	{
		authRoutes.Post(
			"/sign-in",
			limiter.New(middleware.RateLimiter(config.RateLimiterConfig{
				Max:        consts.RateLimiterSensitiveLimit,
				Expiration: consts.RateLimiterExpiration,
			}, receiverRouter.app.RateLimiterStorage)),
			fiberpkgrequest.Middleware(&request.UsersAuthSignIn{}, fiberpkgrequest.JSON),
			receiverRouter.usersController.UsersAuthSignIn,
		)

		authRoutes.Post(
			"/verify",
			limiter.New(middleware.RateLimiter(config.RateLimiterConfig{
				Max:        consts.RateLimiterSensitiveLimit,
				Expiration: consts.RateLimiterExpiration,
			}, receiverRouter.app.RateLimiterStorage)),
			fiberpkgrequest.Middleware(&request.UsersAuthVerify{}, fiberpkgrequest.JSON),
			receiverRouter.usersController.UsersAuthVerify,
		)

	}

	resetPasswordRoutes := paramRouter.Group("/reset-password")
	{
		resetPasswordRoutes.Post(
			"/",
			limiter.New(middleware.RateLimiter(config.RateLimiterConfig{
				Max:        consts.RateLimiterSensitiveLimit,
				Expiration: consts.RateLimiterExpiration,
			}, receiverRouter.app.RateLimiterStorage)),
			fiberpkgrequest.Middleware(&request.UsersResetPassword{}, fiberpkgrequest.JSON),
			receiverRouter.usersController.UsersResetPassword,
		)

		resetPasswordRoutes.Post(
			"/verify",
			limiter.New(middleware.RateLimiter(config.RateLimiterConfig{
				Max:        consts.RateLimiterSensitiveLimit,
				Expiration: consts.RateLimiterExpiration,
			}, receiverRouter.app.RateLimiterStorage)),
			fiberpkgrequest.Middleware(&request.UsersResetPasswordVerify{}, fiberpkgrequest.JSON),
			receiverRouter.usersController.UsersResetPasswordVerify,
		)

		resetPasswordRoutes.Post(
			"/update",
			limiter.New(middleware.RateLimiter(config.RateLimiterConfig{
				Max:        consts.RateLimiterSensitiveLimit,
				Expiration: consts.RateLimiterExpiration,
			}, receiverRouter.app.RateLimiterStorage)),
			fiberpkgrequest.Middleware(&request.UsersResetPasswordUpdate{}, fiberpkgrequest.JSON),
			receiverRouter.usersController.UsersResetPasswordUpdate,
		)
	}
}
