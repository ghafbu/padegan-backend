package govalidatorcustom

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func ghGEndsWith(paramValidator *validator.Validate) {
	err := paramValidator.RegisterValidation("gh_g_ends_with", func(fl validator.FieldLevel) bool {
		// Get Current Field Request
		// *string || string
		thisCurrentFieldRequest := fl.Field().String()

		// Check Logic
		// Get the parameters from the tag (".jpg .jpeg .png ...")
		// Return []string
		thisTagParam := strings.Fields(fl.Param())

		for _, ext := range thisTagParam {
			if strings.HasSuffix(thisCurrentFieldRequest, ext) {
				return true
			}
		}

		return false
	})
	if err != nil {
		return
	}
}
