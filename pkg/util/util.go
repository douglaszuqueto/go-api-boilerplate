package util

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// GenerateID GenerateID
func GenerateID() string {
	id := uuid.New()

	return id.String()
}

// GeneratePassword generatePassword
func GeneratePassword(password string) (string, error) {
	var passwordRequired error = errors.New("Senha obrigatória")

	if len(password) == 0 {
		return "", passwordRequired
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", passwordRequired
	}

	return string(hash), nil
}
