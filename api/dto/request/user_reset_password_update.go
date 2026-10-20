package request

type UsersResetPasswordUpdate struct {
	Mobile         string `json:"mobile" before_validation:"trim" validate:"required,gh_printable,gh_sanitize=high,gh_phone_format_e164=mobile" sanitize_blue:"high" escape:"true"`
	NewPassword    string `json:"new_password" san:"trim" validate:"required,gh_length=8-256,gh_printable,gh_is_valid_and_secure_password" sanitize_blue:"high"`
	RepeatPassword string `json:"repeat_password" san:"trim" validate:"required,eqfield=NewPassword,gh_length=3-256"`
	Uuid           string `json:"uuid" before_validation:"trim" validate:"required,gh_printable,gh_sanitize=high"`
}
