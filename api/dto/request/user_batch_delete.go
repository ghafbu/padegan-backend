package request

type UsersBatchDelete struct {
	Ids []uint64 `json:"ids" validate:"required,gh_array_len=1 100,dive,gt=0"`
}
