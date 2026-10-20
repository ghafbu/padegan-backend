package validatepkg

import (
	"fmt"
	"html"
	"reflect"
)

// CustomTagEscape exclaim adds punctuated enthusiasm to a string.
/*
TODO:
ایت باع هنوز برای استراکت های تو در تو اماده نیست
میتونید از تابع sanitize_blue ایده برداری بکنید
*/
func CustomTagEscape(v interface{}) error {
	val := reflect.ValueOf(v).Elem()

	// اگر struct نیست، خطا می‌دهیم
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("expected a struct")
	}

	// انجام trim برای هر فیلد که تگ "trim" دارد
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := val.Type().Field(i)

		// اگر فیلد از نوع string باشد و تگ escape داشته باشد
		if field.Kind() == reflect.String && fieldType.Tag.Get("escape") == "true" {
			// escape کردن رشته
			field.SetString(html.EscapeString(field.String()))
		}

		// اگر فیلد از نوع struct باشد (nested struct)، به صورت بازگشتی trim می‌شود
		if field.Kind() == reflect.Struct {
			err := CustomTagEscape(field.Addr().Interface())
			if err != nil {
				return err
			}
		}
	}

	return nil
}
