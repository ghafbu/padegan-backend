package responsepkg

import (
	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
)

// Code403
//
// این کد به معنای این است که سرور درخواست را فهمیده، اما اجازه انجام آن را به دلیل محدودیت‌های دسترسی نمی‌دهد.
//
// /*
func Code403(
	paramMessage []messagepkg.Response,
	paramData interface{},
	paramFeature ...EntityV1UseCaseFeature,
) EntityV1UseCase {
	// SnakeCase message id
	// paramMessage.Id = messagepkg.UtlBuildIdText(paramMessage.Id)

	// Generate response
	var thisOutput = EntityV1UseCase{}
	thisOutput.Status = false
	thisOutput.StatusCode = 403
	thisOutput.Message = paramMessage
	thisOutput.Data = paramData

	// Feature add
	if len(paramFeature) > 0 {
		if Feature, ok := any(paramFeature[0]).(EntityV1UseCaseFeature); ok {
			thisOutput.Feature = Feature
		}
	}

	// Log: app response
	// if LoggerTransform, ok := any(paramLogger).(interface.ILogger); ok {
	//	LoggerCategoryAppResponse(LoggerTransform, thisOutput)
	// }

	return thisOutput
}
