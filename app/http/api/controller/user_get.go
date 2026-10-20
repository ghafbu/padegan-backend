package controller

import (
	"net/url"

	"github.com/ghafbu/padegan-backend/api/dto/params"
	"github.com/ghafbu/padegan-backend/api/dto/request"
	"github.com/ghafbu/padegan-backend/pkg/fiberpkg"
	"github.com/gofiber/fiber/v3"
)

func (receiverController *UsersController) UsersGet(ctx fiber.Ctx) error {
	thisHeader := ctx.GetReqHeaders()

	thisRequest, thisTransform := ctx.Locals("request").(*request.UsersGet)
	if !thisTransform {
		return fiberpkg.ReturnRequestTransformInController(ctx)
	}

	thisQueryUnescaped, thisErr := url.QueryUnescape(fiber.Params[string](ctx, "id"))
	if thisErr != nil {
		return fiberpkg.ReturnFailedQueryUnescapeInController(ctx)
	}

	var thisParams = params.UsersGet{
		Id: thisQueryUnescaped,
	}

	var thisOutput = receiverController.useCase.UsersGet(
		ctx.Context(),
		thisRequest,
		thisParams,
		thisHeader,
	)

	return fiberpkg.ReturnAuto(ctx, thisOutput)
}
