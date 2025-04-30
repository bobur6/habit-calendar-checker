package validators

import (
	"fmt"
	"go-rest-project/internal/models"
)

// userID должен быть подставлен из контекста в handler'ах при создании и обновлении
func ValidateHabitList(list *models.HabitList) error {
	if list.UserID == 0 {
		return fmt.Errorf("user ID is required")
	}
	if list.Name == "" {
		return fmt.Errorf("habit list name is required")
	}
	return nil
}
