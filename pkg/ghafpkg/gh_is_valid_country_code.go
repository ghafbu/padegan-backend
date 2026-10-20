package ghafpkg

import (
	"strconv"

	"github.com/nyaruka/phonenumbers"
	"golang.org/x/text/unicode/norm"
)

func GHIsValidCountryCode(code string) bool {
	normalized := norm.NFC.String(code)

	countryCode, err := strconv.Atoi(normalized)
	if err != nil {
		return false
	}

	if phonenumbers.GetRegionCodeForCountryCode(countryCode) == phonenumbers.UNKNOWN_REGION {
		return false
	}

	return true
}
