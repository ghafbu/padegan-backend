package request

type UsersBatchGet struct {
	Ids           *[]uint64 `json:"ids" validate:"omitempty,dive,gt=0"`
	Mobiles       *[]string `before_validation:"trim" json:"mobiles" validate:"omitempty,dive,gh_printable,gh_sanitize=high,gh_phone_format_e164=mobile" sanitize_blue:"high" escape:"true"`
	Phones        *[]string `before_validation:"trim" json:"phones" validate:"omitempty,dive,gh_printable,gh_sanitize=high,gh_phone_format_e164=phone" sanitize_blue:"high" escape:"true"`
	NationalCodes *[]string `before_validation:"trim" json:"national_codes" validate:"omitempty,dive,gh_length=10,numeric,gh_printable,gh_sanitize=high" sanitize_blue:"high" escape:"true"`
	Emails        *[]string `before_validation:"trim,lower" json:"emails" validate:"omitempty,dive,gh_printable,gh_sanitize=high,gh_email" sanitize_blue:"high" escape:"true"`
	Usernames     *[]string `before_validation:"trim" json:"usernames" validate:"omitempty,lowercase,gh_printable,gh_username,gh_sanitize=high,gh_length=3-256"`
	Statuses      *[]uint   `json:"statuses" validate:"omitempty,dive,oneof=1 2 3"`
	OrderBy       string    `before_validation:"trim" json:"order_by" before_validation:"trim,lower" validate:"required,gh_char_abc_with_space_and_zwj,gh_printable_with_zwj,gh_sanitize=high,oneof=id family mobile mobile_country_code phone phone_country_code email birth_date national_code username org_theme_mode language profile_status status" sanitize_blue:"high" escape:"true"`
	OrderMethod   string    `before_validation:"trim" json:"order_method" before_validation:"trim,lower" validate:"required,gh_char_abc_with_space_and_zwj,gh_printable_with_zwj,gh_sanitize=high,oneof=asc desc" sanitize_blue:"high" escape:"true"`
	// Usernames   *[]string `before_validation:"trim" json:"usernames" validate:"omitempty,dive,gh_length=3-256,gh_char_abc_with_space_and_zwj,gh_printable_with_zwj,gh_word_limit=1-5,gh_sanitize=high" sanitize_blue:"high" escape:"true"`
}
