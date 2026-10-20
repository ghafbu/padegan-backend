package ghafpkg

import (
	"errors"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

var (
	errCannotEmpty  = errors.New("chars cannot be empty")
	errMustBeGTZero = errors.New("length must be greater than zero")
)

func GenCode(length int, chars string) (string, error) {
	if length <= 0 {
		return "", errMustBeGTZero
	}
	if len(chars) == 0 {
		return "", errCannotEmpty
	}

	rand.NewSource(time.Now().UnixNano())
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}

	return string(result), nil
}

func GenUUID() string {
	return uuid.New().String()
}
