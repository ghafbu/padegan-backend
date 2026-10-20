package response

type UsersReview struct {
	Users []UsersPublic `json:"members"`
	Paper UsersPaper    `json:"paper"`
}
