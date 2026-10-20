package responsepkg

type EntityV1 struct {
	Status          bool        `json:"status"`
	StatusCode      int         `json:"status_code"`
	Message         interface{} `json:"message"`
	Data            interface{} `json:"data"`
	Path            string      `json:"path"`
	ExecuteDuration int64       `json:"execute_duration"`
}
