package govalidatorcustom

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

func ghRepeat(paramValidator *validator.Validate) {
	err := paramValidator.RegisterValidation("gh_repeat", func(fl validator.FieldLevel) bool {
		// Password         *string `json:"password" san:"trim" validate:"omitempty,length=3-256"`
		// RepeatPassword   *string `json:"repeat_password" san:"trim" validate:"omitempty,repeat=Password"`

		// داده‌ی فعلی (فیلد `RepeatPassword`)
		thisCurrentData := fl.Field().String()

		// ساختار والد ریکوئسست
		parent := fl.Parent()

		// نام فیلد والد (که در تگ `repeat` مشخص شده)
		thisParentFieldName := fl.Param()

		// پیدا کردن فیلد `Password`
		field := parent.FieldByName(thisParentFieldName)
		if !field.IsValid() || (field.Kind() != reflect.String && field.Kind() != reflect.Pointer) {
			return false
		}

		// بررسی نوع فیلد `Password`
		if field.Kind() == reflect.Pointer {
			if field.IsNil() {
				// اگر `Password` nil است، تایید درست است
				return true
			}
			ptrValue := field.Elem()
			if ptrValue.Kind() == reflect.String {
				thisParentData := ptrValue.String()
				// اگر مقدار `Password` موجود است، مقایسه باید انجام شود
				return thisCurrentData == thisParentData
			}
			return false
		}

		// در غیر این صورت فیلد `Password` از نوع رشته است
		thisParentData := field.String()
		return thisCurrentData == thisParentData
	})
	if err != nil {
		return
	}
}
