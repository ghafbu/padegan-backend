package slugpkg

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// Make ایجاد اسلاگ از رشته ورودی
func Make(text string, separator ...string) string {
	var sep string
	if len(separator) > 0 {
		sep = separator[0]
	} else {
		sep = "-"
	}

	// تبدیل حروف به کوچک (حروف انگلیسی)
	text = strings.ToLower(text)

	// تبدیل اعداد فارسی به اعداد انگلیسی
	text = convertPersianDigits(text)

	// حذف کاراکترهای غیرمجاز، به جز حروف فارسی و انگلیسی و اعداد
	pattern := fmt.Sprintf(`[^\p{L}\p{N}\s\x{200c}%s]`, separator)
	re := regexp.MustCompile(pattern)
	text = re.ReplaceAllString(text, "")

	// جایگزینی نیم‌فاصله با خط تیره
	text = strings.ReplaceAll(text, "\u200c", sep)

	// جایگزینی فاصله‌ها و خطوط زیر با خط تیره
	re = regexp.MustCompile(`[\s_-]+`)
	text = re.ReplaceAllString(text, sep)

	// حذف خطوط تیره اضافی در ابتدا و انتهای رشته
	text = strings.Trim(text, sep)

	// حذف خطوط تیره اضافی در نقاط مختلف رشته
	text = cleanSep(text, sep)

	return text
}

var persianDigits = map[rune]string{
	'۰': "0",
	'۱': "1",
	'۲': "2",
	'۳': "3",
	'۴': "4",
	'۵': "5",
	'۶': "6",
	'۷': "7",
	'۸': "8",
	'۹': "9",
}

// cleanUpDashes حذف خطوط تیره اضافی
func cleanSep(text string, sep string) string {
	var result strings.Builder
	var sepRune = []rune(sep)
	dashFound := false
	for _, char := range text {
		if char == sepRune[0] {
			if !dashFound {
				result.WriteRune(char)
				dashFound = true
			}
		} else {
			result.WriteRune(char)
			dashFound = false
		}
	}
	return result.String()
}

// isLetterOrDigit بررسی اینکه آیا کاراکتر ورودی حرف یا عدد است
func isLetterOrDigit(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

// convertPersianDigits تبدیل اعداد فارسی به اعداد انگلیسی
func convertPersianDigits(text string) string {
	result := ""
	for _, char := range text {
		if digit, ok := persianDigits[char]; ok {
			result += digit
		} else {
			result += string(char)
		}
	}
	return result
}
