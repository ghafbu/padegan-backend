package controller

import (
	"github.com/ghafbu/padegan-backend/api/dto/request"
	"github.com/ghafbu/padegan-backend/pkg/fiberpkg"
	"github.com/gofiber/fiber/v3"
)

func (receiverController *UsersController) UsersResetPassword(ctx fiber.Ctx) error {
	var thisHeader = ctx.GetReqHeaders()

	thisRequest, thisTransform := ctx.Locals("request").(*request.UsersResetPassword)
	if !thisTransform {
		return fiberpkg.ReturnRequestTransformInController(ctx)
	}

	var thisOutput = receiverController.useCase.UsersResetPassword(
		ctx.Context(),
		thisRequest,
		thisHeader,
	)

	return fiberpkg.ReturnAuto(ctx, thisOutput)
}
