package validatepkg

import (
	"encoding/json"

	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
	validation "github.com/go-ozzo/ozzo-validation"
)

func TranslateOzzoValidation(lang string, ErrorList validation.Errors) ([]messagepkg.Response, []CustomValidate) {
	var _mess []messagepkg.Response
	var _data []CustomValidate
	var ItemCustom CustomValidate

	//_msg := messagepkg.Response{}

	for key, value := range ErrorList {
		if value != nil {
			ItemCustom.Name = key
			//ItemCustom.Tag = OldItem.Tag()
			//ItemCustom.Param = strings.Split(OldItem.Param(), "-")
			//ItemCustom.Kind = OldItem.Kind()

			//MyCustom.Field = key
			//MyCustom.Message = value.Error()
			json.Unmarshal([]byte(value.Error()), &ItemCustom.Message)
			//MyCustom.Message = _msg
			//MyCustom.Message = Message{
			//	Code: 0,
			//	Text: value.Error(),
			//}

			_mess = append(_mess, ItemCustom.Message)
			_data = append(_data, ItemCustom)
		}
	}

	return _mess, _data
}
