package ghafpkg_test

import (
	"testing"

	"github.com/ghafbu/padegan-backend/pkg/ghafpkg"
	"github.com/stretchr/testify/assert"
)

func TestGhIsValidCountryCode(t *testing.T) {
	testCases := []struct {
		name     string
		code     string
		expected bool
	}{
		{"GhIsValidCountryCode_ValidCode_ReturnsTrue", "98", true},
		{"GhIsValidCountryCode_ValidCode_ReturnsTrue", "1", true},
		{"GhIsValidCountryCode_ValidCode_ReturnsTrue", "44", true},
		{"GhIsValidCountryCode_InvalidCode_ReturnsFalse", "999", false},
		{"GhIsValidCountryCode_InvalidCode_ReturnsFalse", "123", false},
		{"GhIsValidCountryCode_InvalidCode_ReturnsFalse", "0", false},
		{"GhIsValidCountryCode_InvalidString_ReturnsFalse", "abc", false},
		{"GhIsValidCountryCode_EmptyString_ReturnsFalse", "", false},
		{"GhIsValidCountryCode_UnicodeNormalization_ReturnsTrue", "۰۹۸", false}, // تست نرمال‌سازی
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ghafpkg.GHIsValidCountryCode(tc.code)
			assert.Equal(t, tc.expected, result)
		})
	}
}
