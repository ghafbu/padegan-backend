package govalidatorcustom

import (
	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"
	"github.com/go-playground/validator/v10"
)

/* data type support:
string, *string, map , *map, slice, *slice
*/

func ghUsername(paramValidator *validator.Validate) {
	err := paramValidator.RegisterValidation("gh_username", func(fl validator.FieldLevel) bool {
		return ghafpkg.GHTraverseVariableTypeAndApplyStringFuncM1(fl.Field(), ghafpkg.GHIsUsername)
	})
	if err != nil {
		return
	}
}
