package validators

import (
	"fmt"
	"strings"
	"go-rest-project/internal/models"
)

var (
	ErrInvalidInput = fmt.Errorf("invalid input")
)

func ValidateUser(user *models.User) error {
	if strings.TrimSpace(user.Username) == "" {
		return fmt.Errorf("%w: username is required", ErrInvalidInput)
	}
	if strings.TrimSpace(user.Email) == "" {
		return fmt.Errorf("%w: email is required", ErrInvalidInput)
	}
	if !strings.Contains(user.Email, "@") {
		return fmt.Errorf("%w: invalid email format", ErrInvalidInput)
	}
	if len(user.Password) < 6 {
		return fmt.Errorf("%w: password must be at least 6 characters long", ErrInvalidInput)
	}
	return nil
}
