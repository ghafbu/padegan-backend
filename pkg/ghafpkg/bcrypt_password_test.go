package ghafpkg

import (
	"strings"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

//go:generate go test -v -race -count=1 ./...

func TestBcryptPassword_ValidPassword_ReturnsEncryptedPassword(t *testing.T) {
	t.Parallel()

	password := "securePassword123"
	encryptedPass, err := BcryptPassword(password)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if encryptedPass == password {
		t.Error("Expected encrypted password to be different from the original password")
	}

	if len(encryptedPass) == 0 {
		t.Error("Expected encrypted password to be non-empty")
	}
}

func TestBcryptPassword_PasswordExceeds72Bytes_ReturnsError(t *testing.T) {
	t.Parallel()

	longPassword := strings.Repeat("a", 73)
	_, err := BcryptPassword(longPassword)

	if err == nil {
		t.Error("Expected an error for password exceeding 72 bytes, but got none")
	}

	expectedError := "password length exceeds 72 bytes"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', but got '%s'", expectedError, err.Error())
	}
}

func TestBcryptPassword_EmptyPassword_ReturnsEncryptedPassword(t *testing.T) {
	t.Parallel()

	password := ""
	encryptedPass, err := BcryptPassword(password)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if len(encryptedPass) == 0 {
		t.Error("Expected encrypted password to be non-empty")
	}
}

func TestBcryptPassword_ValidPassword_ReturnsDifferentEncryptedPasswords(t *testing.T) {
	t.Parallel()

	password := "securePassword123"
	encryptedPass1, err1 := BcryptPassword(password)
	encryptedPass2, err2 := BcryptPassword(password)

	if err1 != nil || err2 != nil {
		t.Errorf("Expected no errors, but got: %v and %v", err1, err2)
	}

	if encryptedPass1 == encryptedPass2 {
		t.Error("Expected encrypted passwords to be different due to bcrypt salt")
	}
}

func TestBcryptPassword_NilPassword_ReturnsError(t *testing.T) {
	t.Parallel()

	_, err := BcryptPassword("")

	if err != nil {
		t.Errorf("Expected no error for empty password, but got: %v", err)
	}
}

func TestBcryptPassword_ValidPassword_CanVerifyWithBcrypt(t *testing.T) {
	t.Parallel()

	password := "securePassword123"
	encryptedPass, err := BcryptPassword(password)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(encryptedPass), []byte(password))
	if err != nil {
		t.Errorf("Expected password verification to succeed, but got: %v", err)
	}
}
