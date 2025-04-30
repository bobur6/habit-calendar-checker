package validators

import (
	"fmt"
	"go-rest-project/internal/models"
)

func ValidateHabit(habit *models.Habit) error {
	if habit.HabitListID == 0 {
		return fmt.Errorf("habit list ID is required")
	}
	if habit.Name == "" {
		return fmt.Errorf("habit name is required")
	}
	return nil
}
