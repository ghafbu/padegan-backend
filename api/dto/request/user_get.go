package request

type UsersGet struct {
	Target string `query:"target" before_validation:"trim,lower" validate:"required,gh_char_abc_with_space_and_zwj,gh_printable_with_zwj,gh_sanitize=high,oneof=id mobile national_code email username" sanitize_blue:"high" escape:"true"`
}
