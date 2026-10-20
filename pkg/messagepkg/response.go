package messagepkg

import "fmt"

type Response struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

func (m Response) Build() string {
	// `{"code": "1000", "text": "Name is required"}`
	return fmt.Sprintf("{\"id\": \"%s\", \"text\": \"%s\"}", m.Id, m.Text)
}
