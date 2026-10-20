package responsepkg

import (
	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
)

// type myTest1 struct {
//	Name    string  `json:"name"`
//	Feature myTest2 `json:"Feature"`
// }
// type myTest2 struct {
//	Api myTest3 `json:"api"`
// }
//
// type myTest3 struct {
//	WithStatus  bool `json:"with_status"`
//	WithMessage bool `json:"with_message"`
//	WithData    bool `json:"with_data"`
// }

// func (u myTest1) Error() string {
//	return ""
// }

func Code401(
	paramMessage []messagepkg.Response,
	paramData interface{},
	paramFeature ...EntityV1UseCaseFeature,
) EntityV1UseCase {
	// SnakeCase message id
	// paramMessage.Id = messagepkg.UtlBuildIdText(paramMessage.Id)

	// Generate response
	var thisOutput = EntityV1UseCase{}
	thisOutput.Status = false
	thisOutput.StatusCode = 401
	thisOutput.Message = paramMessage
	thisOutput.Data = paramData

	// fmt.Println("thisOutput00 :", myTest1{})
	// fmt.Println("thisOutput00 :", EntityV1UseCase{})
	// fmt.Println("thisOutput-message :", paramMessage)
	// fmt.Println("thisOutput-data :", paramData)
	// fmt.Println("thisOutput1 :", thisOutput)

	// Feature add
	if len(paramFeature) > 0 {
		if Feature, ok := any(paramFeature[0]).(EntityV1UseCaseFeature); ok {
			thisOutput.Feature = Feature
		}
	}

	// fmt.Println("thisOutput2:", thisOutput)
	// Log: app response
	// if LoggerTransform, ok := any(paramLogger).(interface.ILogger); ok {
	//	LoggerCategoryAppResponse(LoggerTransform, thisOutput)
	// }

	return thisOutput
}
