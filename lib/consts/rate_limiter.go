package consts

import "time"

const (
	RateLimiterStorageTimeout = 30 * time.Second // For redis

	RateLimiterPattern         = "%s%s%s"
	RateLimiterSeparator       = ":"
	RateLimiterPrefixKey       = "rl:"
	RateLimiterKeyUserLoggedIn = "u:"
	RateLimiterKeyGuest        = "g:"
	RateLimiterCookieName      = "rl_g" // For guest and save in browser
	RateLimiterCookieHTTPOnly  = true
	RateLimiterCookieSecure    = true
	RateLimiterCookieSameSite  = "Lex"

	RateLimiterCookieExpireTime = 7 * 24 * time.Hour

	//RateLimiterOverallLimit   = 2000 // 2000 درخواست در ساعت
	RateLimiterGeneralLimit = 1000 // 1000 درخواست برای routes عمومی
	//RateLimiterAPILimit       = 500  // 500 درخواست برای APIها
	RateLimiterSensitiveLimit = 20 // 20 درخواست برای routes حساس
	RateLimiterExpiration     = 1 * time.Hour
)
