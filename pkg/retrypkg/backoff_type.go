package retrypkg

type BackoffType int

const (
	ExponentialBackoff BackoffType = iota
	LinearBackoff
	ConstantBackoff
)
