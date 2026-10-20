package defaulttestpkgdtopassword

import "strings"

func responseCode200(paramContent string, paramData interface{}, paramDescription ...string) expected {
	var thisContent = paramContent
	var thisData interface{} = paramData
	var thisDescription = messageExpendedCode200

	// check param message
	if len(thisDescription) > 0 {
		thisDescription = strings.Join(paramDescription, ", ")
	}

	// return
	return expected{
		Content:     thisContent,
		Description: thisDescription,
		Status:      true,
		StatusCode:  200,
		Data:        thisData,
	}
}

func responseCode400(paramContent string, paramData interface{}, paramDescription ...string) expected {
	var thisContent = paramContent
	var thisData interface{} = paramData
	var thisDescription = messageExpendedCode400

	// check param message
	if len(thisDescription) > 0 {
		thisDescription = strings.Join(paramDescription, ", ")
	}

	// return
	return expected{
		Content:     thisContent,
		Description: thisDescription,
		Status:      false,
		StatusCode:  400,
		Data:        thisData,
	}
}
