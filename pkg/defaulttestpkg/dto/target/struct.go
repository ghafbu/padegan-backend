package defaulttestpkg

type DTOTargetStruct struct {
	Id       int
	Content  string
	Category string
	Scenario string
	Provider string
	Expected Expected
}

type Expected struct {
	Description string
	Status      bool
	StatusCode  int
	Message     interface{}
}
