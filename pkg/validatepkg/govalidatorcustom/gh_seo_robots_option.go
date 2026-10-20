package govalidatorcustom

import (
	"github.com/go-playground/validator/v10"
)

/* data type support:
string, *string, map , *map, slice, *slice
*/

func ghSeoRobotsOption(paramValidator *validator.Validate) {
	err := paramValidator.RegisterValidation("gh_seo_robots_option", func(fl validator.FieldLevel) bool {
		val := fl.Field().String()
		switch val {
		case "index,follow", "noindex,follow", "index,nofollow", "noindex,nofollow":
			return true
		}
		return false
	})

	if err != nil {
		return
	}
}
