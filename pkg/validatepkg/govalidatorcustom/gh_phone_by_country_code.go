package govalidatorcustom

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/nyaruka/phonenumbers"
)

func ghPhoneByCountryCode(paramValidator *validator.Validate) {
	err := paramValidator.RegisterValidation("gh_phone_by_country_code", func(fl validator.FieldLevel) bool {
		// Phone             *string    `json:"phone" san:"standard,trim" validate:"omitempty,phone_get_parent=PhoneCountryCode"`
		// PhoneCountryCode  *string    `json:"phone_country_code" san:"standard,trim" validate:"omitempty"`

		//fmt.Println("phone_by_country_code:")
		// Find the parent struct
		parent := fl.Parent()

		// Find the name of the country code field specified in the tag
		countryCodeFieldName := fl.Param()

		//fmt.Println("phone 1:", fl.Field().String())
		//fmt.Println("phone 2 (countryCodeFieldName):", countryCodeFieldName)
		// Find the country code value
		field := parent.FieldByName(countryCodeFieldName)
		if !field.IsValid() || (field.Kind() != reflect.String && field.Kind() != reflect.Pointer) {
			//fmt.Println("phone 3 - type error:", field.Kind())
			return false
		}

		var countryCode string
		if field.Kind() == reflect.String {
			countryCode = field.String()
		} else if field.Kind() == reflect.Ptr {
			ptrValue := field.Elem()
			if ptrValue.IsValid() && ptrValue.Kind() == reflect.String {
				countryCode = ptrValue.String()
			} else {
				return false
			}
		} else {
			return false
		}

		// Get the mobile number
		mobile := fl.Field().String()

		//println("mobile:", mobile)

		// Combine country code and mobile number & Validate the phone number
		combineCountryCodeAndMobile := countryCode + mobile

		num, err := phonenumbers.Parse(
			combineCountryCodeAndMobile, "")
		if err != nil || !phonenumbers.IsValidNumber(num) {
			//fmt.Println("phone 4 (full):", combineCountryCodeAndMobile)
			//fmt.Println("phone 5 (full extract):", fmt.Sprintf("%s%s", countryCode, mobile))
			return false
		}
		//fmt.Println("phone 6 (full extract):", fmt.Sprintf("%s%s", countryCode, mobile))
		return true
	})
	if err != nil {
		return
	}
}
