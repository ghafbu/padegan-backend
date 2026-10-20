package ghafpkg

import (
	"encoding/json"
	"fmt"
)

func GHGetJsons(paramData ...interface{}) (string, error) {
	var a = ""
	for _, d := range paramData {
		// تبدیل داده به JSON با تورفتگی مناسب
		jsonBytes, err := json.MarshalIndent(d, "", "  ")
		if err != nil {
			fmt.Printf("GHPrint Error: JSON: %v\n", err)
			return "", err
		} else {
			a = a + string(jsonBytes)
		}
	}

	return a, nil
}

func GHGetJson(paramData interface{}) (string, error) {
	// تبدیل داده به JSON با تورفتگی مناسب
	jsonBytes, err := json.MarshalIndent(paramData, "", "  ")
	if err != nil {
		fmt.Printf("GHPrint Error: JSON: %v\n", err)
		return "", err
	}

	return string(jsonBytes), nil
}
