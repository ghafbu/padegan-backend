package govalidatorcustom

import (
	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"
	"github.com/go-playground/validator/v10"
)

/* data type support:
string, *string, map , *map, slice, *slice
*/

func ghPrintableWithZWJ(paramValidator *validator.Validate) {
	err := paramValidator.RegisterValidation("gh_printable_with_zwj", func(fl validator.FieldLevel) bool {
		return ghafpkg.GHTraverseVariableTypeAndApplyStringFuncM1(fl.Field(), ghafpkg.GHIsPrintableWithZWJ)

		//field := fl.Field()
		//switch field.Kind() {
		//case reflect.String:
		//	return ghafpkg.GHIsPrintableWithZWJ(field.String())
		//case reflect.Ptr:
		//	if field.IsNil() {
		//		return true
		//	}
		//	return ghafpkg.GHIsPrintableWithZWJ(field.Elem().String())
		//case reflect.Map:
		//	for _, key := range field.MapKeys() {
		//		val := field.MapIndex(key)
		//		if key.Kind() == reflect.String && !ghafpkg.GHIsPrintableWithZWJ(key.String()) {
		//			return false
		//		}
		//		if val.Kind() == reflect.String && !ghafpkg.GHIsPrintableWithZWJ(val.String()) {
		//			return false
		//		}
		//	}
		//	return true
		//case reflect.Slice, reflect.Array:
		//	for i := 0; i < field.Len(); i++ {
		//		if field.Index(i).Kind() == reflect.String && !ghafpkg.GHIsPrintableWithZWJ(field.Index(i).String()) {
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
