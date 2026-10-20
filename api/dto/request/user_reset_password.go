package request

type UsersResetPassword struct {
	Mobile string `json:"mobile" before_validation:"trim" validate:"required,gh_printable,gh_sanitize=high,gh_phone_format_e164=mobile" sanitize_blue:"high" escape:"true"`
}
