package responsepkg

import (
	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
)

type Code500CustomResponseData struct {
	Request          interface{}
	RequestOwnerName string
	Err              string
}

func Code500(
	paramMessage []messagepkg.Response,
	paramData interface{},
	paramFeature ...EntityV1UseCaseFeature,
) EntityV1UseCase {
	// SnakeCase message id
	// paramMessage.Id = messagepkg.UtlBuildIdText(paramMessage.Id)

	// Generate response
	var thisOutput = EntityV1UseCase{}
	thisOutput.Status = false
	thisOutput.StatusCode = 500
	thisOutput.Message = paramMessage
	thisOutput.Data = paramData

	// Feature add
	if len(paramFeature) > 0 {
		if Feature, ok := any(paramFeature[0]).(EntityV1UseCaseFeature); ok {
			thisOutput.Feature = Feature
		}
	}

	// DEV MODE
	thisOutput.Feature = EntityV1UseCaseFeature{Api: EntityV1UseCaseApi{
		WithStatus:  true,
		WithMessage: true,
		WithData:    true,
	}}

	return thisOutput
}
