package govalidatorcustom

import (
	"github.com/go-playground/validator/v10"
	"github.com/microcosm-cc/bluemonday"
)

// TODO اگر پوینتر ارسال بشه نمیتونه بررسی بکنه

func ghSanitize(paramValidator *validator.Validate) {
	err := paramValidator.RegisterValidation("gh_sanitize", func(fl validator.FieldLevel) bool {
		//ghafpkg.GHPrintLn("IsSanitize-------------------start")

		var thisBluMondayPolicy *bluemonday.Policy

		// get param
		thisTagParam := fl.Param()

		// Get Current Field Request
		// *string || string
		thisCurrentFieldRequest := fl.Field().String()

		//fmt.Println("IsSanitize-before", thisCurrentFieldRequest)
		// New Sanitize
		switch thisTagParam {
		case "high":
			thisBluMondayPolicy = bluemonday.StrictPolicy()
		default:
			thisBluMondayPolicy = bluemonday.StrictPolicy()
		}

		thisSanitizeString := thisBluMondayPolicy.Sanitize(thisCurrentFieldRequest)
		//ghafpkg.GHPrintLn("IsSanitize-------------------thisSanitizeString:", thisSanitizeString)
		if thisSanitizeString != thisCurrentFieldRequest {
			return false
		}

		return true
	})
	if err != nil {
		return
	}
}
