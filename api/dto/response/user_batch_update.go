package response

type UsersBatchUpdate struct {
	Users []UsersPublic `json:"members"`
	Paper UsersPaper    `json:"paper"`
}
