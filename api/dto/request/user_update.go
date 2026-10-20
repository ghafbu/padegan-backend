package request

type UsersUpdate struct {
	Name   string `json:"name" before_validation:"trim,lower" validate:"required,gh_length=3-256,gh_char_abc_with_space_and_zwj,gh_printable_with_zwj,gh_word_limit=1-5,gh_sanitize=high" sanitize_blue:"high" escape:"true"`
	Family string `json:"family" before_validation:"trim,lower"  validate:"required,gh_length=3-256,gh_char_abc_with_space_and_zwj,gh_printable_with_zwj,gh_word_limit=1-5,gh_sanitize=high" sanitize_blue:"high" escape:"true"`
	// Mobile            string  `json:"mobile" before_validation:"trim" validate:"required,gh_printable,gh_sanitize=high,gh_phone_format_e164=mobile" sanitize_blue:"high" escape:"true"`
	// MobileCountryCode *string `json:"mobile_country_code" before_validation:"trim" validate:"omitempty,gh_printable,gh_sanitize=high" sanitize_blue:"high" escape:"true"`
	Phone            *string `json:"phone" before_validation:"trim" validate:"omitempty,gh_printable,gh_sanitize=high,gh_phone_format_e164=phone" sanitize_blue:"high" escape:"true"`
	PhoneCountryCode *string `json:"phone_country_code" before_validation:"trim" validate:"omitempty,gh_printable,gh_sanitize=high" sanitize_blue:"high" escape:"true"`
	Email            *string `json:"email" before_validation:"trim,lower" validate:"omitempty,gh_printable,gh_sanitize=high,gh_email" sanitize_blue:"high" escape:"true"`
	BirthDate        *string `json:"birth_date" before_validation:"trim" validate:"omitempty,gh_printable,gh_sanitize=high,datetime=2006-01-02" sanitize_blue:"high" escape:"true"`
	NationalCode     *string `json:"national_code" before_validation:"trim,lower" validate:"omitempty,gh_length=10,numeric,gh_printable,gh_sanitize=high"  sanitize_blue:"high" escape:"true"`
	Username         *string `json:"username" before_validation:"trim" validate:"omitempty,lowercase,gh_printable,gh_username,gh_sanitize=high,gh_length=3-256"`
	Gender           *uint   `json:"gender" validate:"omitempty,oneof=1 2"`
	ProfileImage     *string `json:"profile_image" before_validation:"trim" validate:"omitempty,url,gh_length=3-1024,gh_printable_with_zwj,gh_sanitize=high"  sanitize_blue:"high"`
	OrgThemeMode     *uint   `json:"org_theme_mode" validate:"omitempty,oneof=1 2 3"` // oneof= 1 2 3
	Language         *uint   `json:"language" validate:"omitempty,oneof=1 2"`         // oneof= 1 2
	ProfileStatus    *uint   `json:"profile_status" validate:"omitempty,oneof=1 2"`   // oneof= 1 2
	Status           *uint   `json:"status" validate:"omitempty,oneof=1 2 3"`         // oneof= 1 2 3
	// Username      *string `json:"username" before_validation:"trim" validate:"omitempty,gh_length=3-256,gh_char_abc_with_space_and_zwj,gh_printable_with_zwj,gh_word_limit=1-5,gh_sanitize=high"  sanitize_blue:"high" escape:"true"`
}
