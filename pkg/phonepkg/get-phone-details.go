package phonepkg

import (
	"fmt"

	"github.com/nyaruka/phonenumbers"
)

func GetPhoneDetails(paramPhone string) (PhoneDetails, error) {
	// Parse the phone number using phonenumbers package
	var thisDetails PhoneDetails
	parsedNumber, err := phonenumbers.Parse(paramPhone, "")
	if err != nil {
		return thisDetails, err
	}

	// **بررسی اعتبار شماره**
	if !phonenumbers.IsValidNumber(parsedNumber) {
		return thisDetails, fmt.Errorf("GetPhoneDetails: invalid phone number: %s", paramPhone)
	}

	// Create PhoneDetails from the parsed number
	thisDetails.PhoneCurrent = paramPhone
	thisDetails.CountryCode = int16(parsedNumber.GetCountryCode())
	thisDetails.CountryCodeString = fmt.Sprintf("%s%d", "+", parsedNumber.GetCountryCode())
	thisDetails.NationalNumber = fmt.Sprintf("%d", parsedNumber.GetNationalNumber())
	thisDetails.PhoneE164 = fmt.Sprintf("%s%d%s", "+", thisDetails.CountryCode, thisDetails.NationalNumber)

	return thisDetails, nil
}
