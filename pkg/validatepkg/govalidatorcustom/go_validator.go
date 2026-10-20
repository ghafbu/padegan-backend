package govalidatorcustom

import (
	"github.com/go-playground/validator/v10"
)

func GoValidator(passwordEntropy float64, paramValidator *validator.Validate) {
	// CheckStatus
	// Tag Name: status
	ghCheckStatus(paramValidator)

	// Length
	// Tag Name: length=3-256
	ghLength(paramValidator)

	// ArrayLen
	// Tag Name: length=3-256
	ghArrayLen(paramValidator)

	// Slugify
	// Tag Name: slugify
	// EX : 影師 => ying-shi
	ghSlugify(paramValidator)

	// Repeat
	// Tag Name: (RepeatPassword) `repeat=Password`
	ghRepeat(paramValidator)

	// PhoneFormatE164
	// Tag Name: phone_format_e164
	// EX: +989112253390 is true
	ghPhoneFormatE164(paramValidator)

	// PhoneByCountryCode
	ghPhoneByCountryCode(paramValidator)

	// RequiredBool
	ghRequiredBool(paramValidator)

	// GEndsWith
	ghGEndsWith(paramValidator)

	// StartsWithPlus
	ghStartsWithPlus(paramValidator)

	// IsValidAndSecurePassword
	ghIsValidAndSecurePassword(passwordEntropy, paramValidator)

	// IsSanitize
	ghSanitize(paramValidator)

	// IsPrintable
	ghPrintable(paramValidator)
	ghPrintableWithZWJ(paramValidator)

	// WordLimit
	ghWordLimit(paramValidator)

	// Char ABC
	ghCharABC(paramValidator)
	ghCharABCWithSpaceAndZWJ(paramValidator)

	// Numeric
	// Numeric(paramValidator)

	// ghEmail
	ghEmail(paramValidator)

	// ghUsername
	ghUsername(paramValidator)

	// IsValidCountryCode
	ghIsValidCountryCode(paramValidator)

	// SeoRobotsOption
	ghSeoRobotsOption(paramValidator)
}
