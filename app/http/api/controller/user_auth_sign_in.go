package controller

import (
	"github.com/ghafbu/padegan-backend/api/dto/request"
	"github.com/ghafbu/padegan-backend/pkg/fiberpkg"
	"github.com/gofiber/fiber/v3"
)

func (receiverController *UsersController) UsersAuthSignIn(ctx fiber.Ctx) error {
	thisHeader := ctx.GetReqHeaders()

	thisRequest, thisTransform := ctx.Locals("request").(*request.UsersAuthSignIn)

	if !thisTransform {
		return fiberpkg.ReturnRequestTransformInController(ctx)
	}

	thisOutput := receiverController.useCase.UsersSignIn(
		ctx.Context(),
		thisRequest,
		thisHeader,
	)

	return fiberpkg.ReturnAuto(ctx, thisOutput)
}
