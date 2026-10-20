package defaulttestpkgdtourlpath

// check https://libphonenumber.appspot.com/

var variablesInvalidChatGPT = []myStruct{
	{
		Id:       1,
		Content:  "/send/my-path",
		Category: "invalid",
		Scenario: "",
		Provider: "ChatGPT",
		Expected: responseCode200("", nil),
	},
}
