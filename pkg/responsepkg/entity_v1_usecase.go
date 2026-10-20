package responsepkg

import "github.com/ghafbu/padegan-backend/pkg/messagepkg"

type EntityV1UseCase struct {
	Status     bool                   `json:"status"`
	StatusCode int                    `json:"status_code"`
	Message    []messagepkg.Response  `json:"message"`
	Data       interface{}            `json:"data"`
	Feature    EntityV1UseCaseFeature `json:"feature"`
}

type EntityV1UseCaseFeature struct {
	Api EntityV1UseCaseApi `json:"api"`
}

type EntityV1UseCaseApi struct {
	WithStatus  bool `json:"with_status"`
	WithMessage bool `json:"with_message"`
	WithData    bool `json:"with_data"`
}

//func (u EntityV1UseCase) Error() string {
//	return ""
//}
