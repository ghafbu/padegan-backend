package methods

import (
	"fmt"
	"reflect"
	"strings"
)

// CustomTagBeforeValidation پردازش تگ‌های `before_validation` را انجام می‌دهد.
func CustomTagBeforeValidation(v interface{}) error {
	val := reflect.ValueOf(v)

	// اگر ورودی یک pointer به struct باشد، مقدار آن را استخراج کنیم
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// بررسی کنیم که struct باشد
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("expected a struct, got %s", val.Kind())
	}

	// حلقه روی فیلدهای struct
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := val.Type().Field(i)
		fieldTagString := fieldType.Tag.Get("before_validation")
		fieldTagSlice := strings.Split(fieldTagString, ",")

		// پردازش مقدار فیلد
		processField(field, fieldTagSlice)
	}

	return nil
}

// actionBeforeValidation عملیات پردازش رشته‌ها را انجام می‌دهد.
func actionBeforeValidation(paramFieldType []string, paramString string) string {
	newString := paramString

	for _, tag := range paramFieldType {
		switch tag {
		case "trim":
			newString = strings.TrimSpace(newString)
		case "lower":
			newString = strings.ToLower(newString)
		case "upper":
			newString = strings.ToUpper(newString)
		}
	}

	return newString
}

// processField مقدار ورودی را پردازش و مقدار جدید را روی آن اعمال می‌کند.
func processField(field reflect.Value, fieldTagSlice []string) {
	switch field.Kind() {
	case reflect.String:
		newFieldData := actionBeforeValidation(fieldTagSlice, field.String())
		field.SetString(newFieldData)

	case reflect.Ptr:
		if !field.IsNil() {
			elem := field.Elem()
			switch elem.Kind() {
			case reflect.String:
				newFieldData := actionBeforeValidation(fieldTagSlice, elem.String())
				field.Set(reflect.ValueOf(&newFieldData))
			case reflect.Slice, reflect.Map, reflect.Struct:
				processField(elem, fieldTagSlice)
			}
		}

	case reflect.Map:
		for _, key := range field.MapKeys() {
			valMap := field.MapIndex(key)

			// پردازش کلیدهای رشته‌ای
			if key.Kind() == reflect.String {
				newKeyData := actionBeforeValidation(fieldTagSlice, key.String())
				field.SetMapIndex(reflect.ValueOf(newKeyData), valMap)
			}

			// پردازش مقادیر رشته‌ای
			if valMap.Kind() == reflect.String {
				newValData := actionBeforeValidation(fieldTagSlice, valMap.String())
				field.SetMapIndex(key, reflect.ValueOf(newValData))
			}
		}

	case reflect.Slice, reflect.Array:
		for k := 0; k < field.Len(); k++ {
			elem := field.Index(k)
			if elem.Kind() == reflect.String {
				newFieldData := actionBeforeValidation(fieldTagSlice, elem.String())
				elem.SetString(newFieldData)
			} else if elem.Kind() == reflect.Struct || elem.Kind() == reflect.Ptr {
				processField(elem, fieldTagSlice)
			}
		}

	case reflect.Struct:
		_ = CustomTagBeforeValidation(field.Addr().Interface()) // پردازش بازگشتی برای `struct`
	}
}

/*func actionBeforeValidation(paramFieldType []string, paramString string) string {
	var newString = paramString

	for i := 0; i < len(paramFieldType); i++ {
		switch paramFieldType[i] {
		case "trim":
			newString = strings.TrimSpace(paramString)
		case "lower":
			newString = strings.ToLower(paramString)
		case "upper":
			newString = strings.ToUpper(paramString)
		}
	}

	return newString
}

func CustomTagBeforeValidation(v interface{}) error {
	val := reflect.ValueOf(v).Elem()

	// اگر struct نیست، خطا می‌دهیم
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("CustomTagBeforeValidation expected a struct")
	}

	// loop in struct field
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := val.Type().Field(i)
		fieldTagString := fieldType.Tag.Get("before_validation")
		fieldTagSlice := strings.Split(fieldTagString, ",")

		// switch field kind
		switch field.Kind() {
		case reflect.String:
			newFieldData := actionBeforeValidation(fieldTagSlice, field.String())
			field.SetString(newFieldData)
		case reflect.Ptr:
			if !field.IsNil() {
				newFieldData := actionBeforeValidation(fieldTagSlice, field.Elem().String())
				field.Set(reflect.ValueOf(&newFieldData))
			}
		case reflect.Map:
			for _, keyMap := range field.MapKeys() {
				valMap := field.MapIndex(keyMap)
				if keyMap.Kind() == reflect.String {
					newFieldData := actionBeforeValidation(fieldTagSlice, keyMap.String())
					field.SetString(newFieldData)
				}
				if valMap.Kind() == reflect.String {
					newFieldData := actionBeforeValidation(fieldTagSlice, valMap.String())
					field.SetString(newFieldData)
				}
			}
		case reflect.Slice, reflect.Array:
			for k := 0; k < field.Len(); k++ {
				if field.Index(k).Kind() == reflect.String {
					newFieldData := actionBeforeValidation(fieldTagSlice, field.Index(k).String())
					field.SetString(newFieldData)
				}
			}
		case reflect.Struct:
			err := CustomTagBeforeValidation(field.Addr().Interface())
			if err != nil {
				return err
			}
		}

	}

	return nil
}*/
