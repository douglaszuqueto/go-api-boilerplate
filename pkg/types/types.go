package types

import (
	"errors"
	"regexp"
)

// UUID type
type UUID string

var uuidRegex = regexp.MustCompile("^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$")

var errInvalidUUID = errors.New("ID inválido")

// IsValid isValid
func (u *UUID) IsValid() error {
	if !uuidRegex.MatchString(string(*u)) {
		return errInvalidUUID
	}
	return nil
}

// String string
func (u *UUID) String() string {
	return string(*u)
}
