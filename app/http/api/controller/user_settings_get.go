package controller

import (
	"github.com/ghafbu/padegan-backend/pkg/fiberpkg"
	"github.com/gofiber/fiber/v3"
)

func (receiverController *UserSettingsController) UserSettingsGet(ctx fiber.Ctx) error {
	thisHeader := ctx.GetReqHeaders()

	var thisOutput = receiverController.useCase.UserSettingsGet(ctx.Context(), thisHeader)

	return fiberpkg.ReturnAuto(ctx, thisOutput)
}
