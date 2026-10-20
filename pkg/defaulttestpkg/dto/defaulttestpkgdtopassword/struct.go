package defaulttestpkgdtopassword

type myStruct struct {
	Id       int
	Content  string
	Category string
	Scenario string
	Provider string
	Expected expected
}

type expected struct {
	Content     string
	Description string
	Status      bool
	StatusCode  int
	Data        interface{}
}
