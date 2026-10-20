package response

type UsersBatchGet struct {
	Users []UsersPublic `json:"members"`
}
