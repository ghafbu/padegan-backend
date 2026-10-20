package types

type UserPermission struct {
	OnlyMySelf bool `json:"only_my_self"`
	Create     bool `json:"create"`
	Delete     bool `json:"delete"`
	Edit       bool `json:"edit"`
}
