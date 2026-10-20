package ghafpkg

import (
	"encoding/json"
	"fmt"
)

// ToJSON converts an interface to a JSON bytes.
func ToJSON(v interface{}) ([]byte, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return []byte{}, fmt.Errorf("error marshaling to JSON: %v", err)
	}
	return bytes, nil
}

// FromJSON converts a JSON bytes to an interface.
func FromJSON(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
