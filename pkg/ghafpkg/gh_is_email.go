package ghafpkg

import (
	"net/mail"
	"regexp"
	"strings"

	"github.com/asaskevich/govalidator"
	"golang.org/x/net/idna"
	"golang.org/x/net/publicsuffix"
	"golang.org/x/text/unicode/norm"
)

// Printable Character is true => abc , 0-9 , !@#$%%... , other language
// Combining Character is true => ôkP => o\u0302kP

func GHIsEmail(email string) bool {
	email = norm.NFC.String(email)

	if !govalidator.IsEmail(email) {
		return false
	}

	_, err1 := mail.ParseAddress(email)
	if err1 != nil {
		return false
	}

	// 1. بررسی طول کل ایمیل (حداکثر 254 کاراکتر)
	if len(email) > 254 {
		return false
	}

	// بخش محلی (Local Part) با استفاده از Whitelist دقیق
	//// Local Part (امن + منطبق بر RFC بدون کاراکترهای خطرناک)
	localPartRegex := regexp.MustCompile(`^[a-zA-Z0-9._+-]+(?:\.[a-zA-Z0-9._+-]+)*$`)

	//// Domain Part (رفع مشکل هیفن در ابتدا/انتهای دامنه)
	domainPartRegex := regexp.MustCompile(`^(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z0-9]{2,}$`)

	// 3. جدا کردن بخش محلی و دامنه
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	local, domain := parts[0], parts[1]

	// 4. بررسی بخش محلی
	if !localPartRegex.MatchString(local) {
		return false
	}

	// 4. بررسی طول بخش محلی (حداکثر 64 کاراکتر)
	if len(local) > 64 {
		return false
	}

	// 5. بررسی نقطه در ابتدا/انتهای بخش محلی (ممنوع)
	if strings.HasPrefix(local, ".") || strings.HasSuffix(local, ".") {
		return false
	}

	// 6. بررسی بخش دامنه
	if !domainPartRegex.MatchString(domain) {
		return false
	}

	// بررسی معتبر بودن TLD با publicsuffix
	tld, err2 := publicsuffix.PublicSuffix(domain)
	if !err2 || len(tld) < 2 || len(tld) > 63 {
		return false // ❌ پسوند دامنه نامعتبر است
	}

	// 6. بررسی بخش‌های دامنه مطابق RFC 1034
	domainParts := strings.Split(domain, ".")
	for _, part := range domainParts {
		if strings.HasPrefix(part, "xn--") {
			decoded, err := idna.ToUnicode(part)
			if err != nil || decoded == "" {
				return false
			}
			// بررسی اینکه برچسب دیکد شده تنها شامل حروف انگلیسی، ارقام و خط تیره باشد
			asciiRegex := regexp.MustCompile(`^[a-zA-Z0-9-]+$`)
			if !asciiRegex.MatchString(decoded) {
				return false
			}
		}

		// 6-1. طول بخش (1-63 کاراکتر)
		if len(part) < 1 || len(part) > 63 {
			return false
		}

		// 8-2. کاراکترهای مجاز: حروف، اعداد، خط تیره (نه در ابتدا/انتها)
		if partRegex := regexp.MustCompile(`^[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?$`); !partRegex.MatchString(part) {
			return false
		}
	}

	return true
}

func GHIsEmail2(paramString string) bool {
	// نرمال‌سازی رشته به فرم NFC برای یکپارچه شدن حروف و علامت‌های ترکیبی
	normalized := norm.NFC.String(paramString)

	_, err1 := mail.ParseAddress(normalized)
	if err1 != nil {
		return false
	}

	// الگوی Regex برای بررسی ساختار ایمیل
	var emailRegex = `^[a-zA-Z0-9][a-zA-Z0-9._%+-]*@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,63}$`
	re := regexp.MustCompile(emailRegex)

	// بررسی فرمت اولیه ایمیل
	if !re.MatchString(normalized) {
		return false
	}

	// تجزیه ایمیل به نام کاربری و دامنه
	parts := strings.Split(normalized, "@")
	if len(parts) != 2 {
		return false
	}

	username, domain := parts[0], parts[1]

	// بررسی حداقل طول نام کاربری (حداقل 3 کاراکتر)
	if len(username) < 3 || len(username) > 64 {
		return false
	}

	// بررسی حداقل طول دامنه (حداقل 3 کاراکتر)
	if len(domain) < 3 || len(domain) > 255 {
		return false
	}

	// بررسی عدم وجود ".." در نام دامنه
	if strings.Contains(domain, "..") {
		return false
	}

	// بررسی عدم وجود "." یا "-" در ابتدا یا انتهای دامنه
	if strings.HasPrefix(domain, "-") || strings.HasPrefix(domain, ".") ||
		strings.HasSuffix(domain, "-") || strings.HasSuffix(domain, ".") {
		return false
	}

	// بررسی تعداد صحیح @ (نباید بیش از یک عدد باشد)
	if strings.Count(normalized, "@") != 1 {
		return false
	}

	// بررسی معتبر بودن TLD با استفاده از publicsuffix
	tld, err := publicsuffix.PublicSuffix(domain)
	if !err || tld == "" || len(tld) < 2 || len(tld) > 63 {
		return false
	}

	return true
}
