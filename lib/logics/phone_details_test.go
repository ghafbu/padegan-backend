package logics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPhoneDetails(t *testing.T) {
	t.Run("Valid Phone Number", func(t *testing.T) {
		phone := "+989112223333"
		paramRequest := struct{}{}
		paramLang := "English"
		paramFieldName := "mobile"

		result, _, success := PhoneDetails(paramRequest, &phone, paramLang, paramFieldName)

		assert.True(t, success)
		assert.NotEmpty(t, result)
	})

	t.Run("Invalid Phone Number", func(t *testing.T) {
		phone := "invalid_phone"
		paramRequest := struct{}{}
		paramLang := "English"
		paramFieldName := "mobile"

		result, response, success := PhoneDetails(paramRequest, &phone, paramLang, paramFieldName)

		assert.False(t, success)
		assert.Empty(t, result)
		assert.Equal(t, 400, response.StatusCode)
	})

	t.Run("Nil Phone Pointer", func(t *testing.T) {
		var phone *string = nil
		paramRequest := struct{}{}
		paramLang := "English"
		paramFieldName := "phone"

		result, response, success := PhoneDetails(paramRequest, phone, paramLang, paramFieldName)

		assert.False(t, success)
		assert.Empty(t, result)
		assert.Equal(t, 400, response.StatusCode)
	})
}
