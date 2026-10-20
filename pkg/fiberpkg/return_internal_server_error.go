package fiberpkg

import (
	"net/http"

	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
	"github.com/gofiber/fiber/v3"
)

func ReturnInternalServerError(ctx fiber.Ctx, _mess any, _data ...any) error {
	var _res = responsepkg.NewV1().
		SetStatus(true).
		SetStatusCode(http.StatusInternalServerError).
		SetMessage(_mess).
		SetData(_data).
		SetPath(ctx.OriginalURL()).
		SetExecutionDuration(GetExecuteDuration(ctx)).
		Build()

	return ctx.Status(http.StatusInternalServerError).JSON(_res)
}
