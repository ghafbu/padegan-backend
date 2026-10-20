package controller

import (
	"github.com/ghafbu/padegan-backend/api/dto/request"
	"github.com/ghafbu/padegan-backend/lib/consts"
	"github.com/ghafbu/padegan-backend/pkg/fiberpkg"
	"github.com/gofiber/fiber/v3"
)

func (receiverController *UsersController) UsersLogout(ctx fiber.Ctx) error {
	var (
		thisHeader        = ctx.GetReqHeaders()
		thisRequestLogout request.UsersLogout
	)

	// Token check is not nil
	if ctx.Locals(consts.ContextKeyUserId) != nil {
		// In middleware_auth: ghafpkg.UserIdContextKey always uint64, therefore not err check
		thisRequestLogout.Id = ctx.Locals(consts.ContextKeyUserId).(uint64)
	}

	// UseCase
	thisOutput := receiverController.useCase.UsersLogout(
		ctx.Context(), &thisRequestLogout, thisHeader)

	return fiberpkg.ReturnAuto(ctx, thisOutput)
}
