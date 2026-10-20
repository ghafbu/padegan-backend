package controller

import (
	"github.com/ghafbu/padegan-backend/api/dto/request"
	"github.com/ghafbu/padegan-backend/pkg/fiberpkg"
	"github.com/gofiber/fiber/v3"
)

func (receiverController *UsersController) UsersCreate(ctx fiber.Ctx) error {
	thisHeader := ctx.GetReqHeaders()

	thisRequest, thisTransform := ctx.Locals("request").(*request.UsersCreate)

	if !thisTransform {
		return fiberpkg.ReturnRequestTransformInController(ctx)
	}

	// Params
	// var thisParams = param.UsersUsersCreate{
	//	//UserId: fiber.Params[uint64](ctx, "user_id"),
	//	//OrganizationId: ctx.Params("organization_id"),
	//	Id: fiber.Params[uint64](ctx, "id"),
	// }

	thisOutput := receiverController.useCase.UsersCreate(
		ctx.Context(),
		thisRequest,
		thisHeader,
	)

	return fiberpkg.ReturnAuto(ctx, thisOutput)
}
