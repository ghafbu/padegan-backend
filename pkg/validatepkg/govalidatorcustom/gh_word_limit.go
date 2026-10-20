package govalidatorcustom

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"
	"github.com/go-playground/validator/v10"
)

/* data type support:
string, *string
// map , *map, slice, *slice

Use: word_limit=10 || word_limit=10-256
*/

func ghWordLimit(paramValidator *validator.Validate) {
	err := paramValidator.RegisterValidation("gh_word_limit", func(fl validator.FieldLevel) bool {
		field := fl.Field()
		thisValueCount := 0
		// type
		//thisType := "string"
		//thisKeyValueCount := 0
		//thisValueValueCount := 0

		// get param
		thisTagParam := fl.Param()

		// check type
		switch field.Kind() {
		case reflect.String:
			thisValueCount = ghafpkg.CountWords(field.String())
		case reflect.Ptr:
			if field.IsNil() {
				return true
			}
			thisValueCount = ghafpkg.CountWords(field.Elem().String())
		//case reflect.Map:
		//	thisType="map"
		//
		//	for _, key := range field.MapKeys() {
		//		val := field.MapIndex(key)
		//		if key.Kind() == reflect.String {
		//			thisKeyValueCount = ghafpkg.CountWords(val.String())
		//		}
		//		if val.Kind() == reflect.String {
		//			thisValueValueCount = ghafpkg.CountWords(val.String())
		//		}
		//	}
		//case reflect.Slice, reflect.Array:
		//	for i := 0; i < field.Len(); i++ {
		//		if field.Index(i).Kind() == reflect.String {
		//			thisValueCount = ghafpkg.CountWords(field.Index(i).String())
		//		}
		//	}
		default:
			return false
		}

		// print console
		//ghafpkg.GHPrintLn("thisValueCount:", thisValueCount)

		if strings.Contains(thisTagParam, "-") {
			// اگر مقدار به‌صورت بازه‌ای (مثلاً 3-256) مشخص شده باشد
			parts := strings.Split(thisTagParam, "-")
			if len(parts) != 2 {
				return false // در صورت عدم رعایت فرمت صحیح، مقدار نامعتبر است
			}

			minWords, err1 := strconv.Atoi(parts[0]) // تبدیل مقدار حداقل کلمات به عدد
			maxWords, err2 := strconv.Atoi(parts[1]) // تبدیل مقدار حداکثر کلمات به عدد
			if err1 != nil || err2 != nil {
				return false // اگر مقدار صحیح نبود، مقدار نامعتبر است
			}

			// بررسی اینکه تعداد کلمات در محدوده مجاز قرار دارد
			return thisValueCount >= minWords && thisValueCount <= maxWords
		} else {
			// اگر مقدار دقیق (مثلاً 10) تعیین شده باشد
			exactWords, err := strconv.Atoi(thisTagParam) // تبدیل مقدار دقیق کلمات به عدد
			if err != nil {
				return false // اگر مقدار معتبر نبود، مقدار نامعتبر است
			}

			//ghafpkg.GHPrintLn("")
			return thisValueCount == exactWords // بررسی اینکه تعداد کلمات دقیقاً برابر مقدار تعیین‌شده باشد
		}
	})
	if err != nil {
		return
	}
}
