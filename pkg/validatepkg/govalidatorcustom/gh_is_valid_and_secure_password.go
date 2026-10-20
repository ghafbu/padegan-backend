package govalidatorcustom

import (
	"github.com/go-playground/validator/v10"

	passwordvalidator "github.com/wagslane/go-password-validator"
)

func ghIsValidAndSecurePassword(passwordEntropy float64, paramValidator *validator.Validate) {
	err := paramValidator.RegisterValidation("gh_is_valid_and_secure_password", func(fl validator.FieldLevel) bool {
		// Get Current Field Request
		// *string || string
		thisCurrentFieldRequest := fl.Field().String()

		if err := passwordvalidator.Validate(thisCurrentFieldRequest, passwordEntropy); err != nil {
			return false
		}

		return true
	})
	if err != nil {
		return
	}
}
