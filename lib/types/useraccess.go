package types

type UserAccess struct {
	Role    string      `json:"role"`
	Modules UserModules `json:"modules"`
}
