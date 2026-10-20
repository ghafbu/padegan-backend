package govalidatorcustom

import (
	"fmt"

	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"
	"github.com/go-playground/validator/v10"
)

// قانون سفارشی ghLength برای بررسی طول با min و max
func ghLength(paramValidator *validator.Validate) {
	err := paramValidator.RegisterValidation("gh_length", func(fl validator.FieldLevel) bool {
		return ghafpkg.GHTraverseVariableTypeAndApplyStringFuncM2(fl.Field(), fl.Param(), ghafpkg.GHIsLength)

		// دریافت مقدار فیلد به صورت رشته
		//value := fl.Field().Interface().(string)

		//return ghafpkg.GHIsLength(value,fl.Param())

		// جدا کردن min و max از پارامتر
		//params := strings.Split(fl.Param(), "-")

		// اگر تعداد پارامترها یکی بود، ادامه دهید و فقط min را بررسی کنید
		//if len(params) == 1 {
		//	thisMin, err1 := strconv.Atoi(params[0])
		//	if err1 != nil {
		//		return false
		//	}
		//	// فقط min را چک می‌کنیم
		//	return thisValueLen >= thisMin
		//}

		//// اگر تعداد پارامترها بیشتر از یک بود، بررسی min و max
		//if len(params) != 2 {
		//	return false // تعداد پارامترها باید دقیقاً 2 باشد
		//}
		//
		//// تبدیل min و max به عدد
		//thisMin, err1 := strconv.Atoi(params[0])
		//thisMax, err2 := strconv.Atoi(params[1])

		//// اگر هر کدام از مقادیر اشتباه بود، بازگرداندن false
		//if err1 != nil || err2 != nil {
		//	return false
		//}
		//
		//// بررسی شرایط
		//if thisValueLen >= thisMin && thisValueLen <= thisMax {
		//	return true
		//}

		//return false
	})

	// بررسی خطا در ثبت قانون سفارشی
	if err != nil {
		fmt.Println("Error registering gh_length validation:", err)
	}
}

// func ghLength(paramValidator *validator.Validate) {
// 	err := paramValidator.RegisterValidation("gh_length", func(fl validator.FieldLevel) bool {
// 		value := fl.Field().Interface().(string)
// 		// params := fl.Param()
// 		var params = make([]string, 2)
// 		params = strings.Split(fl.Param(), "-")
// 		thisMin, err1 := strconv.Atoi(params[0])
// 		thisMax, err2 := strconv.Atoi(params[1])
//
// 		if err1 != nil || err2 != nil {
// 			return false
// 		}
//
// 		var thisValueLen = utf8.RuneCountInString(value)
// 		// fmt.Printf("\n length:%v-%v| thisLength:%v | value:%q \n", min, max, thisValueLen, value)
//
// 		if thisValueLen >= thisMin && thisValueLen <= thisMax {
// 			return true
// 		} else {
// 			return false
// 		}
//
// 		// return true
// 		// return value.String() != "unknown"
// 	})
// 	if err != nil {
// 		return
// 	}
// }
//
