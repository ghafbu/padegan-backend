package phonepkg

import (
	"github.com/nyaruka/phonenumbers"
)

func GetPhoneStandard(paramMobile string, paramFormat string) (string, error) {
	// When you want [ +9809112253390 || +989112253390 ] convert to standard number
	var thisMobileParse, thisErr = phonenumbers.Parse(paramMobile, "")
	if thisErr != nil {
		return "", thisErr
	}

	//get mobile
	var mobile string
	if paramFormat == "national" {
		// standard is 0911 225 3390
		mobile = phonenumbers.Format(thisMobileParse, phonenumbers.NATIONAL)
	} else if paramFormat == "international" {
		// standard is +98 911 225 3390
		mobile = phonenumbers.Format(thisMobileParse, phonenumbers.INTERNATIONAL)
	} else if paramFormat == "e164" {
		// standard is +989112253390
		mobile = phonenumbers.Format(thisMobileParse, phonenumbers.E164)
	}

	//return
	return mobile, nil
}
