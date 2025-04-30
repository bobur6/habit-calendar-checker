package repository

import (
	"gorm.io/gorm"
	"go-rest-project/internal/models"
)

type HabitListRepository interface {
	Create(habitList *models.HabitList) error
	FindByID(id uint) (*models.HabitList, error)
	FindByUserID(userID uint) ([]models.HabitList, error)
	Update(habitList *models.HabitList) error
	Delete(id uint) error
}

type habitListRepository struct {
	db *gorm.DB
}

func NewHabitListRepository(db *gorm.DB) HabitListRepository {
	return &habitListRepository{db: db}
}

func (r *habitListRepository) Create(habitList *models.HabitList) error {
	return r.db.Create(habitList).Error
}

func (r *habitListRepository) FindByID(id uint) (*models.HabitList, error) {
	var habitList models.HabitList
	err := r.db.Preload("Habits").First(&habitList, id).Error
	if err != nil {
		return nil, err
	}
	return &habitList, nil
}

func (r *habitListRepository) FindByUserID(userID uint) ([]models.HabitList, error) {
	var habitLists []models.HabitList
	err := r.db.Where("user_id = ? AND is_archived = ?", userID, false).
		Preload("Habits").
		Find(&habitLists).Error
	if err != nil {
		return nil, err
	}
	return habitLists, nil
}

func (r *habitListRepository) Update(habitList *models.HabitList) error {
	return r.db.Save(habitList).Error
}

func (r *habitListRepository) Delete(id uint) error {
	return r.db.Delete(&models.HabitList{}, id).Error
}