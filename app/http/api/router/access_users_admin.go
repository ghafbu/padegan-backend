package router

import (
	"github.com/ghafbu/padegan-backend/api/dto/request"
	"github.com/ghafbu/padegan-backend/config"
	"github.com/ghafbu/padegan-backend/lib/consts"
	"github.com/ghafbu/padegan-backend/lib/middleware"
	"github.com/ghafbu/padegan-backend/pkg/fiberpkg"
	"github.com/ghafbu/padegan-backend/pkg/fiberpkg/fiberpkgrequest"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
)

func (receiverRouter Router) accessUsersAdmin(paramRouter fiber.Router) {
	// Set middleware for protected routes
	//paramRouter.Use(
	//	middleware.FiberAdminAuth(
	//		receiverRouter.app.Config.SecretKey.JwtAuthAccess,
	//		consts.DefaultKeyAdminJWTAccess,
	//		receiverRouter.app.Token,
	//		receiverRouter.app.Redis,
	//	))
	{
		// Create
		paramRouter.Post(
			"/",
			limiter.New(middleware.RateLimiter(config.RateLimiterConfig{
				Max:        consts.RateLimiterGeneralLimit,
				Expiration: consts.RateLimiterExpiration,
			}, receiverRouter.app.RateLimiterStorage)),
			fiberpkgrequest.Middleware(&request.UsersCreate{}, fiberpkgrequest.JSON),
			receiverRouter.usersController.UsersCreate,
		)

		// Get
		paramRouter.Get(
			"/:id",
			limiter.New(middleware.RateLimiter(config.RateLimiterConfig{
				Max:        consts.RateLimiterGeneralLimit,
				Expiration: consts.RateLimiterExpiration,
			}, receiverRouter.app.RateLimiterStorage)),
			fiberpkgrequest.Middleware(&request.UsersGet{}, fiberpkgrequest.Query),
			//fiberpkg.MiddlewareCheckAccess(),
			receiverRouter.usersController.UsersGet,
		)

		// Update
		paramRouter.Patch(
			"/:id",
			limiter.New(middleware.RateLimiter(config.RateLimiterConfig{
				Max:        consts.RateLimiterGeneralLimit,
				Expiration: consts.RateLimiterExpiration,
			}, receiverRouter.app.RateLimiterStorage)),
			fiberpkgrequest.Middleware(&request.UsersUpdate{}, fiberpkgrequest.JSON),
			fiberpkg.MiddlewareCheckAccess(),
			receiverRouter.usersController.UsersUpdate,
		)

		// Review
		paramRouter.Get(
			"/",
			limiter.New(middleware.RateLimiter(config.RateLimiterConfig{
				Max:        consts.RateLimiterGeneralLimit,
				Expiration: consts.RateLimiterExpiration,
			}, receiverRouter.app.RateLimiterStorage)),
			fiberpkgrequest.Middleware(&request.UsersReview{}, fiberpkgrequest.Query),
			receiverRouter.usersController.UsersReview,
		)

		// Batch-get
		paramRouter.Post(
			"/batch/get",
			limiter.New(middleware.RateLimiter(config.RateLimiterConfig{
				Max:        consts.RateLimiterGeneralLimit,
				Expiration: consts.RateLimiterExpiration,
			}, receiverRouter.app.RateLimiterStorage)),
			fiberpkgrequest.Middleware(&request.UsersBatchGet{}, fiberpkgrequest.JSON),
			receiverRouter.usersController.UsersBatchGet,
		)

		// Batch-update
		paramRouter.Patch(
			"/batch/update",
			limiter.New(middleware.RateLimiter(config.RateLimiterConfig{
				Max:        consts.RateLimiterGeneralLimit,
				Expiration: consts.RateLimiterExpiration,
			}, receiverRouter.app.RateLimiterStorage)),
			fiberpkgrequest.Middleware(&request.UsersBatchUpdate{}, fiberpkgrequest.JSON),
			receiverRouter.usersController.UsersBatchUpdate,
		)

		// Batch-delete
		paramRouter.Delete(
			"/batch/delete",
			limiter.New(middleware.RateLimiter(config.RateLimiterConfig{
				Max:        consts.RateLimiterGeneralLimit,
				Expiration: consts.RateLimiterExpiration,
			}, receiverRouter.app.RateLimiterStorage)),
			fiberpkgrequest.Middleware(&request.UsersBatchDelete{}, fiberpkgrequest.JSON),
			receiverRouter.usersController.UsersBatchDelete,
		)

		// Logout
		paramRouter.Post(
			"/logout",
			limiter.New(middleware.RateLimiter(config.RateLimiterConfig{
				Max:        consts.RateLimiterGeneralLimit,
				Expiration: consts.RateLimiterExpiration,
			}, receiverRouter.app.RateLimiterStorage)),
			receiverRouter.usersController.UsersLogout)
	}
}
