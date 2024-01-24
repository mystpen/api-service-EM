package pkg

import (
	"api-service/internal/types"
	"errors"
)

var (
	ErrEmptyField = errors.New("name and surname cannot be empty")
)

func Validate(user types.User) error {
	if user.Name == "" || user.Surname == "" {
		return ErrEmptyField
	}
	return nil
}
