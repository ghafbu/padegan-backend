package govalidatorcustom

import (
	"github.com/ghafbu/padegan-backend/pkg/phonepkg"
	"github.com/go-playground/validator/v10"
	"github.com/nyaruka/phonenumbers"
)

func ghPhoneFormatE164(paramValidator *validator.Validate) {
	err := paramValidator.RegisterValidation("gh_phone_format_e164", func(fl validator.FieldLevel) bool {
		// +989112253390 is true
		// +9809112253390 is false
		// 09112253390 is false

		// Get phone
		thisPhone := fl.Field().String()

		//fmt.Println("thisPhone:", thisPhone)

		// get phone parse
		thisPhoneParse, err := phonenumbers.Parse(thisPhone, "")
		if err != nil {
			return false
		}

		//ghafpkg.GHPrintLnWithColor(ghafpkg.GHPrintColorPaletteWarning, "phonenumbers.IsValidNumber(thisPhoneParse):", phonenumbers.IsValidNumber(thisPhoneParse))

		// check basic standard
		if !phonenumbers.IsValidNumber(thisPhoneParse) || !phonenumbers.IsPossibleNumber(thisPhoneParse) {
			return false
		}

		// get phone details
		thisPhoneDetails, err := phonepkg.GetPhoneDetails(thisPhone)
		if err != nil {
			return false
		}

		// get phone type | return int | ex: mobile=1
		thisPhoneType := phonenumbers.GetNumberType(thisPhoneParse)

		// get param
		thisTagParam := fl.Param()

		// e164 by +9809112253390 is true BUT +989112253390 is true
		//if thisTagParam == "mobile" {
		//	if thisPhoneType == phonenumbers.MOBILE {
		//		return true
		//	} else {
		//		return false
		//	}
		//} else if thisTagParam == "phone" {
		//	if thisPhoneType == phonenumbers.FIXED_LINE {
		//		return true
		//	} else {
		//		return false
		//	}
		//} else if thisTagParam == "fax" {
		//	if thisPhoneType == phonenumbers.FIXED_LINE {
		//		return true
		//	} else {
		//		return false
		//	}
		//} else if thisTagParam == "mobile-or-phone" {
		//	if thisPhoneType == phonenumbers.FIXED_LINE_OR_MOBILE {
		//		return true
		//	} else {

		//		return false
		//	}
		//}

		//ghafpkg.GHPrintLnWithColor(ghafpkg.GHPrintColorPaletteWarning, "thisPhoneType:", thisPhoneType)
		//ghafpkg.GHPrintLnWithColor(ghafpkg.GHPrintColorPaletteWarning, "thisTagParam:", thisTagParam)

		if thisTagParam == "mobile" {
			if thisPhoneType != phonenumbers.MOBILE {
				return false
			}
		} else if thisTagParam == "phone" {
			if thisPhoneType != phonenumbers.FIXED_LINE {
				return false
			}
		} else if thisTagParam == "fax" {
			if thisPhoneType != phonenumbers.FIXED_LINE {
				return false
			}
		} else if thisTagParam == "mobile-or-phone" {
			if thisPhoneType != phonenumbers.FIXED_LINE_OR_MOBILE {
				return false
			}
		}

		//fmt.Println(thisTagParam, thisPhoneDetails.PhoneE164, thisPhoneDetails.CountryCode, thisPhoneDetails.NationalNumber, thisPhone)

		// check e164 standard , check 0 in prefix
		//ghafpkg.GHPrintLnWithColor(ghafpkg.GHPrintColorPaletteWarning, "thisPhone:", thisPhone)
		//ghafpkg.GHPrintLnWithColor(ghafpkg.GHPrintColorPaletteWarning, "thisPhoneDetails.PhoneE164:", thisPhoneDetails.PhoneE164)
		if thisPhone != thisPhoneDetails.PhoneE164 {
			return false
		}

		return true
	})
	if err != nil {
		return
	}
}

//func IsE164(paramPhone string, paramPhoneConvertToE164 string) bool {
//	// check e164 standard , check 0 in prefix
//	if paramPhone != paramPhoneConvertToE164 {
//		return false
//	} else {
//		return true
//	}
//}
