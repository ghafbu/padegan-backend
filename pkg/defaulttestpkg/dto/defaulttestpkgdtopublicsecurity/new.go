package defaulttestpkgdtopublicsecurity

func NewDefaultTestDTOPublicSecurity() []myStruct {
	var thisProperty []myStruct

	thisProperty = append(thisProperty, variablesSecurityChatGPT...)
	thisProperty = append(thisProperty, variablesSecurityDeepSeek...)
	thisProperty = append(thisProperty, variablesSecurityQwen...)

	return thisProperty
}
