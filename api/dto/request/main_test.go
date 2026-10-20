package request

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/ghafbu/padegan-backend/pkg/fiberpkg/fiberpkgrequest"
	"github.com/ghafbu/padegan-backend/pkg/validatepkg/govalidatorcustom"
	"github.com/gofiber/fiber/v3"
)

func fiberServerSetup(t *testing.T) *fiber.App {
	app := fiber.New(
		fiber.Config{
			StructValidator: govalidatorcustom.NewBindStructTags(80).Setup(),
		},
	)

	// Router
	app.Post(
		"/sign-in",
		signInController,
		fiberpkgrequest.Middleware(&UsersAuthSignIn{}, fiberpkgrequest.JSON),
	)
	app.Post(
		"/verify",
		verifyController,
		fiberpkgrequest.Middleware(&UsersAuthVerify{}, fiberpkgrequest.JSON),
	)
	app.Post(
		"/",
		createController,
		fiberpkgrequest.Middleware(&UsersCreate{}, fiberpkgrequest.JSON),
	)
	app.Get(
		"/:id",
		getController,
		fiberpkgrequest.Middleware(&UsersGet{}, fiberpkgrequest.JSON),
	)
	app.Post(
		"/batch/get",
		batchGetController,
		fiberpkgrequest.Middleware(&UsersBatchGet{}, fiberpkgrequest.JSON),
	)
	app.Patch(
		"/batch/update",
		batchUpdateController,
		fiberpkgrequest.Middleware(&UsersBatchUpdate{}, fiberpkgrequest.JSON),
	)
	app.Delete(
		"/batch/delete",
		batchDeleteController,
		fiberpkgrequest.Middleware(&UsersBatchDelete{}, fiberpkgrequest.JSON),
	)

	return app
}

func signInController(ctx fiber.Ctx) error {
	thisRequest, thisTransform := ctx.Locals("request").(*UsersAuthSignIn)
	if !thisTransform {
		return errors.New("failed to retrieve request data: invalid or nil value")
	}

	thisResponse := UsersAuthSignIn{Mobile: thisRequest.Mobile}

	return ctx.JSON(thisResponse)
}

func verifyController(ctx fiber.Ctx) error {
	thisRequest, thisTransform := ctx.Locals("request").(*UsersAuthVerify)
	if !thisTransform {
		return errors.New("failed to retrieve request data: invalid or nil value")
	}

	thisResponse := UsersAuthVerify{
		Mobile: thisRequest.Mobile,
		Code:   thisRequest.Code,
	}

	return ctx.JSON(thisResponse)
}

func createController(ctx fiber.Ctx) error {
	thisRequest, thisTransform := ctx.Locals("request").(*UsersCreate)
	if !thisTransform {
		return errors.New("failed to retrieve request data: invalid or nil value")
	}

	return ctx.JSON(thisRequest)
}

func getController(ctx fiber.Ctx) error {
	thisRequest, thisTransform := ctx.Locals("request").(*UsersGet)
	if !thisTransform {
		return errors.New("failed to retrieve request data: invalid or nil value")
	}

	return ctx.JSON(thisRequest)
}

func batchGetController(ctx fiber.Ctx) error {
	thisRequest, thisTransform := ctx.Locals("request").(*UsersBatchGet)
	if !thisTransform {
		return errors.New("failed to retrieve request data: invalid or nil value")
	}

	return ctx.JSON(thisRequest)
}

func batchUpdateController(ctx fiber.Ctx) error {
	thisRequest, thisTransform := ctx.Locals("request").(*UsersBatchUpdate)
	if !thisTransform {
		return errors.New("failed to retrieve request data: invalid or nil value")
	}

	return ctx.JSON(thisRequest)
}

func batchDeleteController(ctx fiber.Ctx) error {
	thisRequest, thisTransform := ctx.Locals("request").(*UsersBatchDelete)
	if !thisTransform {
		return errors.New("failed to retrieve request data: invalid or nil value")
	}

	return ctx.JSON(thisRequest)
}

func getRandomItem(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", fmt.Errorf("slice is empty")
	}
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(slice))

	return slice[randomIndex], nil
}
