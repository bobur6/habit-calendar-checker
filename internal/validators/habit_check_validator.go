package validators

import (
	"fmt"
	"go-rest-project/internal/models"
	"time"
)

func ValidateHabitCheck(check *models.HabitCheck) error {
	if check.HabitID == 0 {
		return fmt.Errorf("habit ID is required")
	}
	if check.Date.IsZero() {
		check.Date = time.Now()
	}
	return nil
}
