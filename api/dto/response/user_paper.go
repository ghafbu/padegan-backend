package response

type UsersPaper struct {
	CurrentPage int   `json:"current_page"`
	TotalPages  int   `json:"total_pages"`
	TotalItems  int64 `json:"total_items"`
	ItemPerPage int   `json:"item_per_page"`
}
