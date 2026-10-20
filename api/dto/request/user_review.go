package request

type UsersReview struct {
	Name              *string `query:"name" before_validation:"trim,lower" validate:"omitempty,gh_length=3-256,gh_char_abc_with_space_and_zwj,gh_printable_with_zwj,gh_word_limit=1-5,gh_sanitize=high" sanitize_blue:"high" escape:"true"`
	Family            *string `query:"family" before_validation:"trim,lower"  validate:"omitempty,gh_length=3-256,gh_char_abc_with_space_and_zwj,gh_printable_with_zwj,gh_word_limit=1-5,gh_sanitize=high" sanitize_blue:"high" escape:"true"`
	Mobile            *string `query:"mobile" before_validation:"trim" validate:"omitempty,gh_printable,gh_sanitize=high,gh_phone_format_e164=mobile" sanitize_blue:"high" escape:"true"`
	MobileCountryCode *string `query:"mobile_country_code" before_validation:"trim,lower" validate:"omitempty,gh_char_abc_with_space_and_zwj,gh_printable_with_zwj,gh_sanitize=high,gh_is_valid_country_code" sanitize_blue:"high" escape:"true"`
	Phone             *string `query:"phone" before_validation:"trim" validate:"omitempty,gh_printable,gh_sanitize=high,gh_phone_format_e164=phone" sanitize_blue:"high" escape:"true"`
	PhoneCountryCode  *string `query:"phone_country_code" before_validation:"trim,lower" validate:"omitempty,gh_char_abc_with_space_and_zwj,gh_printable_with_zwj,gh_sanitize=high,gh_is_valid_country_code" sanitize_blue:"high" escape:"true"`
	Email             *string `query:"email" before_validation:"trim,lower" validate:"omitempty,gh_printable,gh_sanitize=high,gh_email" sanitize_blue:"high" escape:"true"`
	BirthDate         *string `query:"birth_date" before_validation:"trim" validate:"omitempty,gh_printable,gh_sanitize=high,datetime=2006-01-02" sanitize_blue:"high" escape:"true"`
	NationalCode      *string `query:"national_code" before_validation:"trim,lower" validate:"omitempty,gh_length=10,numeric,gh_printable,gh_sanitize=high"  sanitize_blue:"high" escape:"true"`
	Username          *string `query:"username" before_validation:"trim" validate:"omitempty,lowercase,gh_printable,gh_username,gh_sanitize=high,gh_length=3-256"`
	Gender            *uint   `query:"gender" validate:"omitempty,oneof=1 2"`
	ProfileImage      *string `query:"profile_image" before_validation:"trim" validate:"omitempty,url,gh_length=3-1024,gh_char_abc_with_space_and_zwj,gh_printable_with_zwj,gh_sanitize=high"  sanitize_blue:"high"`
	OrgThemeMode      *uint   `query:"org_theme_mode" validate:"omitempty,oneof=1 2 3"` // oneof= 1 2 3
	Language          *uint   `query:"language" validate:"omitempty,oneof=1 2"`         // oneof= 1 2
	ProfileStatus     *uint   `query:"profile_status" validate:"omitempty,oneof=1 2"`   // oneof= 1 2
	Status            *uint   `query:"status" validate:"omitempty,oneof=1 2 3"`         // oneof= 1 2 3
	CreatedAtStart    *string `query:"created_at_start" validate:"omitempty,gh_printable,gh_sanitize=high,datetime=2006-01-02T15:04:05Z07:00"`
	CreatedAtEnd      *string `query:"created_at_end" validate:"omitempty,gh_printable,gh_sanitize=high,datetime=2006-01-02T15:04:05Z07:00,gtefield=CreatedAtStart"`
	UpdatedAtStart    *string `query:"updated_at_start" validate:"omitempty,gh_printable,gh_sanitize=high,datetime=2006-01-02T15:04:05Z07:00"`
	UpdatedAtEnd      *string `query:"updated_at_start" validate:"omitempty,gh_printable,gh_sanitize=high,datetime=2006-01-02T15:04:05Z07:00,gtefield=UpdatedAtStart"`
	Page              *int    `query:"page" validate:"omitempty"`
	ItemPerPage       *int    `query:"item_per_page" validate:"omitempty"`
	OrderBy           *string `query:"order_by" before_validation:"trim,lower" validate:"omitempty,gh_char_abc_with_space_and_zwj,gh_printable_with_zwj,gh_sanitize=high,oneof=id family mobile mobile_country_code phone phone_country_code email birth_date national_code username org_theme_mode language profile_status status" sanitize_blue:"high" escape:"true"`
	OrderMethod       *string `query:"order_method" before_validation:"trim,lower" validate:"omitempty,gh_char_abc_with_space_and_zwj,gh_printable_with_zwj,gh_sanitize=high,oneof=asc desc" sanitize_blue:"high" escape:"true"`
	// DeletedAtStart    *string `query:"deleted_at_start" validate:"omitempty,gh_printable,gh_sanitize=high,datetime=2006-01-02T15:04:05Z07:00"`
	// DeletedAtEnd      *string `query:"deleted_at_start" validate:"omitempty,gh_printable,gh_sanitize=high,datetime=2006-01-02T15:04:05Z07:00,gtefield=DeletedAtStart"`
}
