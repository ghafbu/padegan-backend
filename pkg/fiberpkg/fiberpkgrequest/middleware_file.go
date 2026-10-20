package fiberpkgrequest

import (
	"github.com/ghafbu/padegan-backend/pkg/headerpkg"
	"github.com/ghafbu/padegan-backend/pkg/validatepkg"
	"github.com/gofiber/fiber/v3"
)

func MiddlewareFile(keyForSaveToFiberContext string) func(ctx fiber.Ctx) error {

	return func(ctx fiber.Ctx) error {
		var _header = headerpkg.New(ctx.GetReqHeaders())
		var err error
		//var files = requests.FilesRequest{}

		multipartForm, err := ctx.MultipartForm()
		if err != nil {
			return validatepkg.Build(ctx, err, _header)
		}

		ctx.Locals(keyForSaveToFiberContext, multipartForm.File)

		return ctx.Next()
	}
}
