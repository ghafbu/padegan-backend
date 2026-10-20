package ghafpkg

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func BcryptPassword(password string) (string, error) {
	if len(password) > 72 {
		return "", fmt.Errorf("password length exceeds 72 bytes")
	}

	var encryptedPass, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return password, fmt.Errorf("unexpected error: %v", err)
	}

	return string(encryptedPass), nil
}
