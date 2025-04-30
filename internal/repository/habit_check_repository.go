package repository

import (
	"time"
	"gorm.io/gorm"
	"go-rest-project/internal/models"
)

type HabitCheckRepository interface {
	Create(check *models.HabitCheck) error
	FindByID(id uint) (*models.HabitCheck, error)
	FindByHabitID(habitID uint) ([]models.HabitCheck, error)
	FindByHabitIDAndDate(habitID uint, date time.Time) (*models.HabitCheck, error)
	Update(check *models.HabitCheck) error
	Delete(id uint) error
}

type habitCheckRepository struct {
	db *gorm.DB
}

func NewHabitCheckRepository(db *gorm.DB) HabitCheckRepository {
	return &habitCheckRepository{db: db}
}

func (r *habitCheckRepository) Create(check *models.HabitCheck) error {
	return r.db.Create(check).Error
}

func (r *habitCheckRepository) FindByID(id uint) (*models.HabitCheck, error) {
	var check models.HabitCheck
	err := r.db.First(&check, id).Error
	if err != nil {
		return nil, err
	}
	return &check, nil
}

func (r *habitCheckRepository) FindByHabitID(habitID uint) ([]models.HabitCheck, error) {
	var checks []models.HabitCheck
	err := r.db.Where("habit_id = ?", habitID).
		Order("date DESC").
		Find(&checks).Error
	if err != nil {
		return nil, err
	}
	return checks, nil
}

func (r *habitCheckRepository) FindByHabitIDAndDate(habitID uint, date time.Time) (*models.HabitCheck, error) {
	var check models.HabitCheck
	err := r.db.Where("habit_id = ? AND date = ?", habitID, date.Format("2006-01-02")).
		First(&check).Error
	if err != nil {
		return nil, err
	}
	return &check, nil
}

func (r *habitCheckRepository) Update(check *models.HabitCheck) error {
	return r.db.Save(check).Error
}

func (r *habitCheckRepository) Delete(id uint) error {
	return r.db.Delete(&models.HabitCheck{}, id).Error
} 