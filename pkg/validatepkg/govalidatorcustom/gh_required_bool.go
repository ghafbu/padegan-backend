package govalidatorcustom

import (
	"github.com/go-playground/validator/v10"
)

func ghRequiredBool(paramValidator *validator.Validate) {
	err := paramValidator.RegisterValidation("gh_required_bool", func(fl validator.FieldLevel) bool {
		val, ok := fl.Field().Interface().(bool)
		return ok && (val || !val)
	})
	if err != nil {
		return
	}
}
