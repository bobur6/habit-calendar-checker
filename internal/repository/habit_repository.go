package repository

import (
	"gorm.io/gorm"
	"go-rest-project/internal/models"
)

type HabitRepository interface {
	Create(habit *models.Habit) error
	FindByID(id uint) (*models.Habit, error)
	FindByHabitListID(habitListID uint) ([]models.Habit, error)
	Update(habit *models.Habit) error
	Delete(id uint) error
}

type habitRepository struct {
	db *gorm.DB
}

func NewHabitRepository(db *gorm.DB) HabitRepository {
	return &habitRepository{db: db}
}

func (r *habitRepository) Create(habit *models.Habit) error {
	return r.db.Create(habit).Error
}

func (r *habitRepository) FindByID(id uint) (*models.Habit, error) {
	var habit models.Habit
	err := r.db.Preload("Checks").First(&habit, id).Error
	if err != nil {
		return nil, err
	}
	return &habit, nil
}

func (r *habitRepository) FindByHabitListID(habitListID uint) ([]models.Habit, error) {
	var habits []models.Habit
	err := r.db.Where("habit_list_id = ? AND is_archived = ?", habitListID, false).
		Preload("Checks").
		Find(&habits).Error
	if err != nil {
		return nil, err
	}
	return habits, nil
}

func (r *habitRepository) Update(habit *models.Habit) error {
	return r.db.Save(habit).Error
}

func (r *habitRepository) Delete(id uint) error {
	return r.db.Delete(&models.Habit{}, id).Error
}