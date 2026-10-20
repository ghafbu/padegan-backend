package ghafpkg

import (
	"regexp"

	"golang.org/x/text/unicode/norm"
)

// تابع که بررسی می‌کند رشته فقط شامل حروف الفبای جهانی باشد.

func CharABC(paramString string) bool {
	// نرمال‌سازی رشته به فرم NFC برای یکپارچه شدن حروف و علامت‌های ترکیبی
	normalized := norm.NFC.String(paramString)

	// استفاده از عبارت منظم برای بررسی حروف الفبای همه زبان‌ها
	re := regexp.MustCompile(`^[\p{L}]+$`)
	return re.MatchString(normalized)
}
