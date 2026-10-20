package defaulttestpkgdtoemail

var variablesInvalidManual = []myStruct{
	{
		Id:       1,
		Content:  "ali!@gmail.com",
		Category: "invalid",
		Scenario: "ایمیل اشتباه",
		Provider: "Manual",
		Expected: responseCode400("", nil),
	},
	{
		Id:       2,
		Content:  "ali#@gmail.com",
		Category: "invalid",
		Scenario: "ایمیل اشتباه",
		Provider: "Manual",
		Expected: responseCode400("", nil),
	},
	{
		Id:       3,
		Content:  "ali$@gmail.com",
		Category: "invalid",
		Scenario: "ایمیل اشتباه",
		Provider: "Manual",
		Expected: responseCode400("", nil),
	},
}
