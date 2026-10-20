package response

type UsersAuthVerify struct {
	User   UsersPublic `json:"member"`
	Tokens UsersToken  `json:"tokens"`
}
