package ghafpkg

import (
	"unicode"

	"golang.org/x/text/unicode/norm"
)

// Printable Character is true => abc , 0-9 , !@#$%%... , other language
// Combining Character is true => ôkP => o\u0302kP

func GHIsPrintableWithZWJ(paramString string) bool {
	// نرمال‌سازی رشته به فرم NFC برای یکپارچه شدن حروف و علامت‌های ترکیبی
	normalized := norm.NFC.String(paramString)

	for _, r := range normalized {
		// نیم فاصله (U+200C) را به عنوان قابل چاپ در نظر می‌گیریم
		if r != '\u200C' && !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}
