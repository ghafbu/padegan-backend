package types

type DateDetailsWithTime struct {
	Year   *int `json:"year"`
	Month  *int `json:"month"`
	Day    *int `json:"day"`
	Hour   *int `json:"hour"`
	Minute *int `json:"minute"`
	Second *int `json:"Second"`
}
