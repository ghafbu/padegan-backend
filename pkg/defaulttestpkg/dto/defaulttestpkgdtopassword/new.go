package defaulttestpkgdtopassword

func NewDefaultTestDTOPassword() []myStruct {
	var thisProperty []myStruct

	thisProperty = append(thisProperty, variablesInvalidChatGPT...)
	thisProperty = append(thisProperty, variablesInvalidDeepSeek...)
	thisProperty = append(thisProperty, variablesInvalidQwen...)
	// thisProperty = append(thisProperty, variablesInvalidManual...)
	//
	thisProperty = append(thisProperty, variablesSecurityChatGPT...)
	thisProperty = append(thisProperty, variablesSecurityDeepSeek...)
	thisProperty = append(thisProperty, variablesSecurityQwen...)

	return thisProperty
}
