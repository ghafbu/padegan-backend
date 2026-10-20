package responsepkg

import (
	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
)

func Code502(
	paramMessage []messagepkg.Response,
	paramData interface{},
	paramFeature ...EntityV1UseCaseFeature,
) EntityV1UseCase {
	// SnakeCase message id
	//paramMessage.Id = messagepkg.UtlBuildIdText(paramMessage.Id)

	// Generate response
	var thisOutput = EntityV1UseCase{}
	thisOutput.Status = false
	thisOutput.StatusCode = 502
	thisOutput.Message = paramMessage
	thisOutput.Data = paramData

	//Feature add
	if len(paramFeature) > 0 {
		if Feature, ok := any(paramFeature[0]).(EntityV1UseCaseFeature); ok {
			thisOutput.Feature = Feature
		}
	}

	return thisOutput
}
