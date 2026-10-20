package govalidatorcustom

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ghArrayLen(paramValidator *validator.Validate) {
	err := paramValidator.RegisterValidation("gh_array_len", func(fl validator.FieldLevel) bool {
		// استخراج پارامترهای min و max از تگ
		params := strings.Split(fl.Param(), " ")
		thisMin, err1 := strconv.Atoi(params[0])
		thisMax, err2 := strconv.Atoi(params[1])

		// بررسی درستی پارامترهای min و max
		if err1 != nil || err2 != nil {
			return false
		}

		// دریافت مقدار فیلد بدون استفاده از reflection برای اشاره‌گرها
		field := fl.Field()

		// بررسی اینکه آیا اشاره‌گر است و اشاره‌گر معتبر است
		if field.Kind() == reflect.Ptr {
			if field.IsNil() {
				return false
			}
			// دسترسی به مقدار اشاره‌گر
			field = field.Elem()
		}

		// بررسی اینکه آیا نوع داده آرایه یا slice است
		switch field.Kind() {
		case reflect.Slice, reflect.Array:
			// محاسبه طول آرایه بدون reflection اضافی
			length := field.Len()
			return length >= thisMin && length <= thisMax
		default:
			// اگر نوع داده آرایه یا slice نبود
			return false
		}
	})
	if err != nil {
		return
	}
}
