package govalidatorcustom

import (
	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"
	"github.com/go-playground/validator/v10"
)

/* data type support:
string, *string, map , *map, slice, *slice
*/

func ghCharABC(paramValidator *validator.Validate) {
	err := paramValidator.RegisterValidation("gh_char_abc", func(fl validator.FieldLevel) bool {
		return ghafpkg.GHTraverseVariableTypeAndApplyStringFuncM1(fl.Field(), ghafpkg.CharABC)
	})
	if err != nil {
		return
	}
}
