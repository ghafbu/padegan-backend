package fiberpkgrequest

import (
	"mime/multipart"
	"reflect"

	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"

	"github.com/ghafbu/padegan-backend/pkg/fiberpkg"
	"github.com/ghafbu/padegan-backend/pkg/headerpkg"
	"github.com/ghafbu/padegan-backend/pkg/validatepkg"
	"github.com/gofiber/fiber/v3"
	"github.com/mitchellh/mapstructure"
)

func Middleware(schema interface{}, BindType string) func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		// get header from ctx Locals
		var _header = headerpkg.New(ctx.GetReqHeaders())
		var err error

		//fmt.Println("inside middleware")

		// Parse incoming data based on the configured source
		RequestStruct := reflect.New(reflect.TypeOf(schema).Elem()).Interface()

		// just bind and sanitize , no validation
		switch BindType {
		case JSON:
			err = ctx.Bind().JSON(RequestStruct)
			if err != nil {
				// ghafpkg.GHPrintFWithColor(ghafpkg.GHPrintColorPaletteWarning, "Fiber-Request-Middleware BindType==JSON > ERR: %v", err)
				return validatepkg.Build(ctx, err, _header)
			}
		case Query:
			err = ctx.Bind().Query(RequestStruct)
			if err != nil {
				return validatepkg.Build(ctx, err, _header)
			}
		case Form:
			err = ctx.Bind().Form(RequestStruct)
			if err != nil {
				// check errors of [ bind/sanitize/ozzovalidation/go-validation/anonymous ]
				return validatepkg.Build(ctx, err, _header)
			}
		case MultipartForm:
			var multipartForm *multipart.Form

			multipartForm, err = ctx.MultipartForm()
			if err != nil {
				return validatepkg.Build(ctx, err, _header)
			}

			err = mapstructure.Decode(multipartForm.File, RequestStruct)
			if err != nil {
				return validatepkg.Build(ctx, err, _header)
			}

		default:
			return fiberpkg.ReturnInternalServerError(ctx, "BindType error")
		}

		ghafpkg.GHPrintJsonLnWithColor(ghafpkg.GHPrintColorPaletteInfo, "Fiber-Request-Middleware end | Request:", RequestStruct)
		ctx.Locals("request", RequestStruct)

		return ctx.Next()
	}
}
