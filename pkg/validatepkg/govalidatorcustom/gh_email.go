package govalidatorcustom

import (
	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"
	"github.com/go-playground/validator/v10"
)

/* data type support:
string, *string, map , *map, slice, *slice
*/

func ghEmail(paramValidator *validator.Validate) {
	err := paramValidator.RegisterValidation("gh_email", func(fl validator.FieldLevel) bool {
		return ghafpkg.GHTraverseVariableTypeAndApplyStringFuncM1(fl.Field(), ghafpkg.GHIsEmail)

		//field := fl.Field()
		//switch field.Kind() {
		//case reflect.String:
		//	return ghafpkg.GHIsEmail(field.String())
		//case reflect.Ptr:
		//	if field.IsNil() {
		//		return true
		//	}
		//	return ghafpkg.GHIsEmail(field.Elem().String())
		//case reflect.Map:
		//	for _, key := range field.MapKeys() {
		//		val := field.MapIndex(key)
		//		//if key.Kind() == reflect.String && !ghafpkg.GHIsEmail(key.String()) {
		//		//	return false
		//		//}
		//		if val.Kind() == reflect.String && !ghafpkg.GHIsEmail(val.String()) {
		//			return false
		//		}
		//	}
		//	return true
		//case reflect.Slice, reflect.Array:
		//	for i := 0; i < field.Len(); i++ {
		//		if field.Index(i).Kind() == reflect.String && !ghafpkg.GHIsEmail(field.Index(i).String()) {
		//			return false
		//		}
		//	}
		//	return true
		//default:
		//	return false
		//}
	})
	if err != nil {
		return
	}
}
