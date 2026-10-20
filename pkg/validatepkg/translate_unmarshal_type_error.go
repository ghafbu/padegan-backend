package validatepkg

import (
	"encoding/json"

	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
)

func TranslateUnmarshalTypeError(lang string, ErrorList *json.UnmarshalTypeError) ([]messagepkg.Response, []CustomValidate) {
	var _mess []messagepkg.Response
	var _data []CustomValidate
	var ItemCustom CustomValidate

	ItemCustom.Name = ErrorList.Field
	//MyCustom.Message = oldBind.Value
	ItemCustom.Message = messagepkg.Response{
		Id: "json-unmarshal",
		//Text: oldBind.Value,
		Text: "unmarshal-type-error",
	}

	_mess = append(_mess, ItemCustom.Message)
	_data = append(_data, ItemCustom)

	return _mess, _data
}
