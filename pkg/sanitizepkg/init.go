package sanitizepkg

import (
	"github.com/microcosm-cc/bluemonday"
)

type Sanitize struct {
	BlueMonday *bluemonday.Policy
}

func NewSanitize() *Sanitize {
	//bluemonday package for html sanitize
	//https://github.com/microcosm-cc/bluemonday
	var BlueMondayStrictPolicy = bluemonday.StrictPolicy()

	var _clone = &Sanitize{
		BlueMonday: BlueMondayStrictPolicy,
	}
	return _clone
}
