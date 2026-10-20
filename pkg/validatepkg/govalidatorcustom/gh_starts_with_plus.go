package govalidatorcustom

import (
	"github.com/go-playground/validator/v10"
)

func ghStartsWithPlus(paramValidator *validator.Validate) {
	err := paramValidator.RegisterValidation("gh_starts_with_plus", func(fl validator.FieldLevel) bool {
		value := fl.Field().Interface().(string)
		return len(value) > 0 && value[0] == '+'
	})
	if err != nil {
		return
	}
}
