package request

type UsersBatchUpdate struct {
	Ids         []uint64 `json:"ids" validate:"required,gh_array_len=1 25,dive,gt=0"`
	Status      uint     `json:"status" validate:"oneof=1 2 3"`
	Page        *int     `json:"page" validate:"omitempty"`
	ItemPerPage *int     `json:"item_per_page" validate:"omitempty"`
	OrderBy     *string  `query:"order_by" before_validation:"trim,lower" validate:"omitempty,gh_char_abc_with_space_and_zwj,gh_printable_with_zwj,gh_sanitize=high,oneof=id family mobile mobile_country_code phone phone_country_code email birth_date national_code username org_theme_mode language profile_status status" sanitize_blue:"high" escape:"true"`
	OrderMethod *string  `query:"order_method" before_validation:"trim,lower" validate:"omitempty,gh_char_abc_with_space_and_zwj,gh_printable_with_zwj,gh_sanitize=high,oneof=asc desc" sanitize_blue:"high" escape:"true"`
}
