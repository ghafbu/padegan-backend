package responsepkg

func NewV2(paramResponse EntityV1UseCase) interface{} {
	//var thisBackEndResponse = make(EntityV2)
	//thisBackEndResponse["status"] = paramResponse.StatusCode
	var thisBackEndResponse interface{}

	if paramResponse.Feature.Api.WithMessage && !paramResponse.Feature.Api.WithData {
		thisBackEndResponse = EntityV2WithMessage{
			Status:  paramResponse.StatusCode,
			Message: paramResponse.Message,
		}
		return thisBackEndResponse
		//thisBackEndResponse["message"] = paramResponse.Message
	}
	if paramResponse.Feature.Api.WithData && !paramResponse.Feature.Api.WithMessage {
		thisBackEndResponse = EntityV2WithData{
			Status: paramResponse.StatusCode,
			Data:   paramResponse.Data,
		}
		return thisBackEndResponse
		//thisBackEndResponse["data"] = paramResponse.Data
	}
	if paramResponse.Feature.Api.WithData && paramResponse.Feature.Api.WithMessage {
		thisBackEndResponse = EntityV2WithMessageAndData{
			Status:  paramResponse.StatusCode,
			Message: paramResponse.Message,
			Data:    paramResponse.Data,
		}
		return thisBackEndResponse
	}

	thisBackEndResponse = EntityV2JustStatus{
		Status: paramResponse.StatusCode,
	}
	return thisBackEndResponse
}
