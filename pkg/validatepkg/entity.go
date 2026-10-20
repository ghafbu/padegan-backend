package validatepkg

import (
	"reflect"

	"github.com/ghafbu/padegan-backend/pkg/messagepkg"
)

//type Message struct {
//	Id   string `json:"id"`
//	Text string `json:"text"`
//}

type CustomValidate struct {
	Name    string              `json:"name"`
	Json    string              `json:"json"`
	Tag     string              `json:"tag"`
	Param   []string            `json:"param"`
	Kind    reflect.Kind        `json:"kind"`
	Message messagepkg.Response `json:"message"`

	//Description string              `json:"description"`
	//Other       interface{}         `json:"other"`
}
