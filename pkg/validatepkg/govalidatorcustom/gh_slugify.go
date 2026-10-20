package govalidatorcustom

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/gosimple/slug"
)

func ghSlugify(paramValidator *validator.Validate) {
	err := paramValidator.RegisterValidation("gh_slugify", func(fl validator.FieldLevel) bool {
		if fl.Field().Kind() == reflect.String {
			value := fl.Field().String()
			slugifyFieldValue := slug.Make(value)
			fl.Field().SetString(slugifyFieldValue)
			return true
		}
		return false
	})
	if err != nil {
		return
	}
}
