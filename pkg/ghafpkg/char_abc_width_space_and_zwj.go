package ghafpkg

import (
	"regexp"

	"golang.org/x/text/unicode/norm"
)

// تابع myWorldABC که بررسی می‌کند رشته فقط شامل حروف الفبای جهانی باشد.
/*

\p{L}: حروف الفبا از تمام زبان‌ها
\x{200C}: نیم‌فاصله (U+200C)
: فاصله (space)

*/

func CharABCWithSpaceAndZWJ(paramString string) bool {
	// نرمال‌سازی رشته به فرم NFC برای یکپارچه شدن حروف و علامت‌های ترکیبی
	normalized := norm.NFC.String(paramString)
	// عبارت منظم به‌روز شده شامل:
	// \p{L} => حروف (Letters)
	// \p{M} => علامت‌های ترکیبی (Marks)
	// \x{200C} => نیم‌فاصله
	// space => فاصله
	re := regexp.MustCompile(`^[\p{L}\p{M}\x{200C} ]+$`)
	return re.MatchString(normalized)
}
