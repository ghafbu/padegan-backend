package validatepkg

import (
	"encoding/json"
	"errors"

	"github.com/ghafbu/padegan-backend/pkg/fiberpkg"
	"github.com/ghafbu/padegan-backend/pkg/headerpkg"
	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
	"github.com/ghafbu/padegan-backend/pkg/responsepkg"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func New() *CustomValidate {
	return &CustomValidate{}
}

func Build(ctx fiber.Ctx, errs error, header headerpkg.Header) error {
	var lang = headerpkg.GetIndexAsString(header.SLanguage)

	var thisMessage []messagepkg.Response
	var thisData []CustomValidate

	// bind error refactor
	var ErrorUnmarshalTypeError *json.UnmarshalTypeError
	if errors.As(errs, &ErrorUnmarshalTypeError) {
		thisMessage, thisData = TranslateUnmarshalTypeError(lang, ErrorUnmarshalTypeError)
		return fiberpkg.ReturnAuto(ctx,
			responsepkg.Code400(
				thisMessage,
				thisData,
			),
		)
		// return fiberpkg.ReturnBadRequest(ctx, thisMessage, thisData)
	}

	// sanitize
	if errs.Error() == "sanitize-error" {
		thisMessage = append(thisMessage, messagepkg.StatusCode500(&lang, "sanitize-error"))
		return fiberpkg.ReturnAuto(ctx,
			responsepkg.Code500(
				thisMessage,
				thisData,
			),
		)
		// return fiberpkg.ReturnInternalServerError(ctx, messagepkg.StatusCode500(&lang, "sanitize-error"))
	}

	// ozzoValidate
	var ErrorListOzzo validation.Errors
	if errors.As(errs, &ErrorListOzzo) {
		thisMessage, thisData = TranslateOzzoValidation(lang, ErrorListOzzo)
		return fiberpkg.ReturnAuto(ctx,
			responsepkg.Code400(
				thisMessage,
				thisData,
			),
		)
		// return fiberpkg.ReturnBadRequest(ctx, thisMessage, thisData)
	}

	// go-validation
	var ErrorListGoValidator validator.ValidationErrors
	if errors.As(errs, &ErrorListGoValidator) {
		// ghafpkg.GHPrintLnWithColor(ghafpkg.GHPrintColorPaletteFatal, "go-validation error | validatepkg/usecase.go")
		thisMessage, thisData = TranslateGoValidator(lang, ErrorListGoValidator)
		return fiberpkg.ReturnAuto(ctx,
			responsepkg.Code400(
				thisMessage,
				thisData,
				responsepkg.EntityV1UseCaseFeature{
					Api: responsepkg.EntityV1UseCaseApi{
						WithMessage: true,
						// WithData:    true,
					},
				},
			),
		)
		// return fiberpkg.ReturnBadRequest(ctx, thisMessage, thisData)
	}

	// anonymous error
	thisMessage = append(thisMessage, messagepkg.Response{
		Id:   "middleware_validation_build_Bind_anonymous_error",
		Text: errs.Error(),
	})
	return fiberpkg.ReturnAuto(ctx,
		responsepkg.Code400(
			thisMessage,
			errs.Error(),
			// responsepkg.EntityV1UseCaseFeature{
			//	Api: responsepkg.EntityV1UseCaseApi{WithData: true},
			// },
		),
	)
	// return fiberpkg.ReturnBadRequest(ctx, messagepkg.Response{
	//	Id:   "validation-build-Bind-anonymous-error",
	//	Text: "validation-build-Bind-anonymous-error",
	// }, errs.Error())
}
