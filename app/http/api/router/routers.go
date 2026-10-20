package router

import "github.com/gofiber/fiber/v3"

func (receiverRouter Router) routers(paramRouter fiber.Router) {

	// ********************* Admin Members Routers ***********************
	thisUsersRouter := paramRouter.Group("/users")

	// All Public APIs
	//thisPublicRouter := paramRouter.Use()
	receiverRouter.accessUsersPublic(thisUsersRouter)

	// All Private APIs For User
	//thisProtectedUserRouter := paramRouter.Use()
	receiverRouter.accessUsersUser(thisUsersRouter)

	// All Private APIs For Admin
	//thisProtectedAdminRouter := paramRouter.Use()
	receiverRouter.accessUsersAdmin(thisUsersRouter)

	// ********************* Admin Settings Routers ***********************
	thisUserSettingsRouter := paramRouter.Group("/users/settings")
	receiverRouter.accessUserSettingsAdmin(thisUserSettingsRouter)
}
