package defaulttestpkg

type DTOString1To256Struct struct {
	Id       int
	Content  string
	Category string
	Scenario string
	Provider string
	Expected Expected
}

type Expected struct {
	Content     string
	Description string
	Status      bool
	StatusCode  int
	Message     interface{}
}
