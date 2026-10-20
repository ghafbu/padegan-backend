package response

type UsersResetPasswordVerify struct {
	Mobile string `json:"mobile"`
	Uuid   string `json:"uuid"`
}
