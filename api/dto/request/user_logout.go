package request

type UsersLogout struct {
	Id uint64 `json:"id" validate:"required"`
	// Or mobile
	// Mobile string `json:"mobile" san:"standard,trim" validate:"required,gh_length=11,mobile=n0911"`
}
