package responsepkg

import "github.com/ghafbu/padegan-backend/pkg/messagepkg"

//type EntityV2 map[string]interface{}

type EntityV2JustStatus struct {
	Status int `json:"status"`
}

type EntityV2WithMessage struct {
	Status  int                   `json:"status"`
	Message []messagepkg.Response `json:"message"`
}
type EntityV2WithData struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}
type EntityV2WithMessageAndData struct {
	Status  int                   `json:"status"`
	Message []messagepkg.Response `json:"message"`
	Data    interface{}           `json:"data"`
}
