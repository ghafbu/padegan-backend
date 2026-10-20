package middleware

import (
	"mime/multipart"
	"reflect"

	"github.com/ghafbu/padegan-backend/pkg/fiberpkg"
	"github.com/ghafbu/padegan-backend/pkg/headerpkg"
	"github.com/ghafbu/padegan-backend/pkg/validatepkg"
	"github.com/gofiber/fiber/v3"
	"github.com/mitchellh/mapstructure"
)

const (
	// and the content-type is `multipart/form-data` or `application/x-www-form-urlencoded`.
	MultipartForm = "multipartForm"
	// Form indicates that the data to validate is in the request body,
	Form = "form"
	// JSON indicates that the data to validate is in the request body,
	// and the content-type is `application/json`.
	JSON = "json"
	// XML indicates that the data to validate is in the request body,
	// and the content-type is `application/xml` or `text/xml`.
	XML = "xml"
	// Query indicates that the data to validate is in the query string.
	Query = "query"
	// Params indicates that the data to validate is in the route parameters.
	Params = "params"
	// ExtractIDFromToken indicates that the data to validate is in the token
	// ExtractIDFromToken = "extractIDFromToken"
)

func FiberRequest(schema interface{}, BindType string) func(ctx fiber.Ctx) error {
	return func(ctx fiber.Ctx) error {
		// get header from ctx Locals
		var _header = headerpkg.New(ctx.GetReqHeaders())
		var err error

		// fmt.Println("inside middleware")

		// Parse incoming data based on the configured source
		RequestStruct := reflect.New(reflect.TypeOf(schema).Elem()).Interface()

		// just bind and sanitize , no validation
		switch BindType {
		case JSON:
			err = ctx.Bind().JSON(RequestStruct)
			if err != nil {
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
		// case ExtractIDFromToken:
		// 	var thisIdFromToken uint64
		//
		// 	// Token check is not nil
		// 	if ctx.Locals(consts.ContextKeyUserId) != nil {
		// 		// In middleware_auth: ghafpkg.UserIdContextKey always uint64, therefore not err check
		// 		thisIdFromToken = ctx.Locals(consts.ContextKeyUserId).(uint64)
		// 	}
		//
		// 	err = json.Unmarshal(
		// 		[]byte(fmt.Sprintf("{\"%s\": %d}", consts.ContextKeyUserId, thisIdFromToken)),
		// 		RequestStruct,
		// 	)
		// 	if err != nil {
		// 		// log.Println("fiber_pkg_request - middleware - err: ", err.Error())
		// 		return fiberpkg.ReturnInternalServerError(ctx, "unmarshal error")
		// 	}
		default:
			return fiberpkg.ReturnInternalServerError(ctx, "BindType error")
		}

		ctx.Locals("request", RequestStruct)

		return ctx.Next()
	}
}
