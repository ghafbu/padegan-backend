package govalidatorcustom

import "github.com/go-playground/validator/v10"

func ghCheckStatus(paramValidator *validator.Validate) {
	err := paramValidator.RegisterValidation("gh_status", func(fl validator.FieldLevel) bool {
		// Check status is 1 | 2 | 3
		if fl.Field().Uint() > 3 {
			return false
		}

		return true
	})
	if err != nil {
		return
	}
}
