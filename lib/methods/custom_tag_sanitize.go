package methods

import (
	"fmt"
	"html"
	"reflect"

	"github.com/ghafbu/padegan-backend/pkg/sanitizepkg"
	"github.com/go-sanitize/sanitize"
)

// CustomTagSanitizeStandard exclaim adds punctuated enthusiasm to a string.
// Create Custom Sanitize Tag
func CustomTagSanitizeStandard(s sanitize.Sanitizer, structValue reflect.Value, idx int) error {
	fieldValue := structValue.Field(idx)
	if fieldValue.Kind() == reflect.Ptr && fieldValue.IsNil() {
		return nil
	}

	if fieldValue.Kind() == reflect.Ptr && !fieldValue.IsNil() {
		fieldValue = fieldValue.Elem()
	}

	if fieldValue.Kind() != reflect.String {
		return fmt.Errorf("standard: invalid type %q", fieldValue.Kind().String())
	}

	v := fieldValue.Interface().(string)
	//if strings.HasSuffix(v, "...") {
	//	v = v[:len(v)-3] + "!"
	//} else if strings.HasSuffix(v, ".") {
	//	v = v[:len(v)-1] + "!"
	//} else {
	//	v += "!"
	//}
	var _sanitize = sanitizepkg.NewSanitize()
	v = _sanitize.String(v)
	// escape string
	v = html.EscapeString(v)

	fieldValue.SetString(v)

	return nil
}
