package defaulttestpkgdtooneof

// check https://libphonenumber.appspot.com/

var variablesInvalidChatGPT = []myStruct{
	{
		Id:       1,
		Content:  "id",
		Category: "invalid",
		Scenario: "",
		Provider: "ChatGPT",
		Expected: responseCode200("", nil),
	}, {
		Id:       2,
		Content:  "userna",
		Category: "invalid",
		Scenario: "",
		Provider: "ChatGPT",
		Expected: responseCode400("", nil),
	},
}
