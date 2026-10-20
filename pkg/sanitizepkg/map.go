package sanitizepkg

// map
func (r Sanitize) Map(input map[string]interface{}) map[string]interface{} {
	for key, value := range input {
		//key action
		input[key] = r.BlueMonday.Sanitize(key)

		//value action
		switch value.(type) {
		case string:
			//fmt.Println(value.(string))
			input[key] = r.BlueMonday.Sanitize(value.(string))
		case map[string]interface{}:
			input[key] = r.Map(value.(map[string]interface{}))
		}
	}

	return input
}
