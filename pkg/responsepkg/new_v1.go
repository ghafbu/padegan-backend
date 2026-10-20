package responsepkg

import "reflect"

type CloneEntityV1 struct {
	responseEntity EntityV1
}

func NewV1() *CloneEntityV1 {
	e := EntityV1{}

	return &CloneEntityV1{responseEntity: e}
}

func (r CloneEntityV1) SetStatus(status bool) CloneEntityV1 {
	r.responseEntity.Status = status

	return r
}

func (r CloneEntityV1) SetStatusCode(statusCode int) CloneEntityV1 {
	r.responseEntity.StatusCode = statusCode

	return r
}

func (r CloneEntityV1) SetMessage(message interface{}) CloneEntityV1 {
	if reflect.ValueOf(message).Kind() != reflect.Slice {
		r.responseEntity.Message = []interface{}{message}

		return r
	}

	r.responseEntity.Message = message

	return r
}

func (r CloneEntityV1) SetData(data interface{}) CloneEntityV1 {
	//r.responseEntity.Data = data
	if reflect.ValueOf(data).Kind() != reflect.Slice {
		r.responseEntity.Data = []interface{}{data}

		return r
	}

	r.responseEntity.Data = data

	return r
}

func (r CloneEntityV1) SetPath(path string) CloneEntityV1 {
	r.responseEntity.Path = path

	return r
}

func (r CloneEntityV1) SetExecutionDuration(executionDuration int64) CloneEntityV1 {
	r.responseEntity.ExecuteDuration = executionDuration

	return r
}

func (r CloneEntityV1) Build() EntityV1 {

	return r.responseEntity
}
