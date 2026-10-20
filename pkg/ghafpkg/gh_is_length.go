package ghafpkg

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

// EX: GHIsLength("1001","1-255") || GHIsLength("1001","4")

func GHIsLength(paramString string, paramRole string) bool {
	// get UTF-8 LEN
	thisLen := utf8.RuneCountInString(paramString)

	// اگر پارامتر حاوی '-' نباشد، بررسی مقدار دقیق
	if !strings.Contains(paramRole, "-") {
		exact, err := strconv.Atoi(paramRole)
		if err != nil {
			return false
		}
		return thisLen == exact
	}

	// اگر پارامتر حاوی '-' باشد، بررسی بازه
	limits := strings.Split(paramRole, "-")
	if len(limits) != 2 {
		return false
	}

	thisMin, err1 := strconv.Atoi(limits[0])
	thisMax, err2 := strconv.Atoi(limits[1])
	if err1 != nil || err2 != nil {
		return false
	}

	// بررسی معتبر بودن بازه
	if thisMin > thisMax {
		return false
	}

	return thisLen >= thisMin && thisLen <= thisMax
}
