package fiberpkg

import (
	"net/http"

	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
	"github.com/gofiber/fiber/v3"
)

func ReturnBadRequest(ctx fiber.Ctx, _mess any, _data ...any) error {
	var _res = responsepkg.NewV1().
		SetStatus(false).
		SetStatusCode(http.StatusBadRequest).
		SetMessage(_mess).
		SetData(_data).
		SetPath(ctx.OriginalURL()).
		SetExecutionDuration(GetExecuteDuration(ctx)).
		Build()

	return ctx.Status(http.StatusBadRequest).JSON(_res)
}
