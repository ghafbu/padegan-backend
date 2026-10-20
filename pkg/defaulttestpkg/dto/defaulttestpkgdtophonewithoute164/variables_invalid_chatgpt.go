package defaulttestpkgdtophonewithoute164

// check https://libphonenumber.appspot.com/

var variablesInvalidChatGPT = []myStruct{
	{
		Id:       1,
		Content:  "+9890002039",
		Category: "invalid",
		Scenario: "شماره موبایل به درستی ارسال شده است",
		Provider: "ChatGPT",
		Expected: responseCode200("", nil),
	},
}
