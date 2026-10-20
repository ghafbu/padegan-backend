package methods

func Ptr[T any](v T) *T {
	// Returns pointer of the variables
	return &v
}
