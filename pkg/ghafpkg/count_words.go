package ghafpkg

import (
	"unicode"

	"golang.org/x/text/unicode/norm"
)

// تابع شمارش تعداد کلمات در یک رشته
func CountWords(paramString string) int {
	normalized := norm.NFC.String(paramString)

	wordCount := 0  // شمارنده تعداد کلمات
	inWord := false // بررسی اینکه آیا درون یک کلمه هستیم یا نه

	for _, r := range normalized {
		// بررسی اینکه آیا کاراکتر حرف یا عدد است
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			if !inWord {
				wordCount++   // افزایش شمارش کلمات
				inWord = true // مشخص کردن اینکه درون یک کلمه هستیم
			}
		} else {
			inWord = false // در صورتی که کاراکتر غیرحرفی باشد، پایان کلمه در نظر گرفته می‌شود
		}
	}

	return wordCount // بازگشت تعداد کلمات
}
