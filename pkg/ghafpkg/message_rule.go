package ghafpkg

import (
	"fmt"
)

type MessageRule struct {
	Code string `json:"code"`
	Text string `json:"text"`
}

func (m MessageRule) Error() string {
	// `{"code": "1000", "text": "Name is required"}`
	return fmt.Sprintf("{\"code\": \"%s\", \"text\": \"%s\"}", m.Code, m.Text)
}
