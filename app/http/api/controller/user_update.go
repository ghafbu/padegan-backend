package controller

import (
	"github.com/ghafbu/padegan-backend/api/dto/params"
	"github.com/ghafbu/padegan-backend/api/dto/request"
	"github.com/ghafbu/padegan-backend/pkg/fiberpkg"
	"github.com/gofiber/fiber/v3"
)

func (receiverController *UsersController) UsersUpdate(ctx fiber.Ctx) error {
	thisHeader := ctx.GetReqHeaders()

	thisRequest, thisTransform := ctx.Locals("request").(*request.UsersUpdate)

	if !thisTransform {
		return fiberpkg.ReturnRequestTransformInController(ctx)
	}

	var thisParams = params.UsersUpdate{
		Id: fiber.Params[uint64](ctx, "id"),
	}

	thisOutput := receiverController.useCase.UsersUpdate(
		ctx.Context(),
		thisRequest,
		thisHeader,
		thisParams,
	)

	return fiberpkg.ReturnAuto(ctx, thisOutput)
}
