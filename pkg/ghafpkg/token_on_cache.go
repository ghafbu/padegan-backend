package ghafpkg

import (
	"fmt"
	"strconv"
)

type Type interface {
	uint64 | string
}

const (
	_keyPattern = "%s:%s"
)

func BuildTokenKeyPattern[K Type](keyDefault string, key K) string {
	if s, ok := any(key).(string); ok {
		return fmt.Sprintf(_keyPattern, keyDefault, s)
	}

	if s, ok := any(key).(uint64); ok {
		return fmt.Sprintf(_keyPattern, keyDefault, strconv.FormatUint(s, 10))
	}

	return ""
	//return fmt.Sprintf(_keyPattern, keyDefault, key)
}

func BuildKeyPattern(defaultKey, key string) string {
	return fmt.Sprintf("%s:%s", defaultKey, key)
}
