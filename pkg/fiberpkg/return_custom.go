package fiberpkg

import (
	"net/http"

	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"
	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
	"github.com/gofiber/fiber/v3"
)

func ReturnRequestTransformInController(ctx fiber.Ctx) error {
	var mes = "request transform error"
	var data = map[string]any{
		"message": mes,
		"code":    ghafpkg.AppGlobalErrorRequestTransformInController,
	}

	var res = responsepkg.NewV1().
		SetStatus(false).
		SetStatusCode(http.StatusInternalServerError).
		SetMessage(mes).
		SetData([]any{data}).
		SetPath(ctx.OriginalURL()).
		SetExecutionDuration(GetExecuteDuration(ctx)).
		Build()

	return ctx.Status(http.StatusInternalServerError).JSON(res)
}

func ReturnFailedQueryUnescapeInController(ctx fiber.Ctx) error {
	var mes = "query unescape error"
	var data = map[string]any{
		"message": mes,
		"code":    ghafpkg.AppGlobalErrorQueryUnescapeInController,
	}

	var res = responsepkg.NewV1().
		SetStatus(false).
		SetStatusCode(http.StatusInternalServerError).
		SetMessage(mes).
		SetData([]any{data}).
		SetPath(ctx.OriginalURL()).
		SetExecutionDuration(GetExecuteDuration(ctx)).
		Build()

	return ctx.Status(http.StatusInternalServerError).JSON(res)
}
