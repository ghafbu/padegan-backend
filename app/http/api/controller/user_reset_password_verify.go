package controller

import (
	"github.com/ghafbu/padegan-backend/api/dto/request"
	"github.com/ghafbu/padegan-backend/pkg/fiberpkg"
	"github.com/gofiber/fiber/v3"
)

func (receiverController *UsersController) UsersResetPasswordVerify(ctx fiber.Ctx) error {
	var thisHeader = ctx.GetReqHeaders()

	thisReq, thisTransform := ctx.Locals("request").(*request.UsersResetPasswordVerify)
	if !thisTransform {
		return fiberpkg.ReturnRequestTransformInController(ctx)
	}

	var thisOutput = receiverController.useCase.UsersResetPasswordVerify(
		ctx.Context(),
		thisReq,
		thisHeader,
	)

	return fiberpkg.ReturnAuto(ctx, thisOutput)
}
