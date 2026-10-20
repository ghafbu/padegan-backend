package ghafpkg

import (
	"regexp"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

// EX: GHIsLength("1001","1-255") || GHIsLength("1001","4")

func GHIsUsername(paramString string) bool {
	// نرمال‌سازی رشته به فرم NFC برای یکپارچه شدن حروف و علامت‌های ترکیبی
	normalized := norm.NFC.String(paramString)

	// بررسی شروع با حرف کوچک (نه عدد)
	firstChar := rune(normalized[0])
	if !unicode.IsLower(firstChar) {
		return false
	}

	// بررسی کل رشته (فقط حروف کوچک و اعداد)
	matched, _ := regexp.MatchString(`^[a-z0-9]+$`, normalized)
	return matched
}
