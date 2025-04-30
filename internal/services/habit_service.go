package services

import (
	"errors"
	"go-rest-project/internal/models"
	"go-rest-project/internal/repository"
	"go-rest-project/internal/validators"
)

type HabitService interface {
	CreateHabit(habit *models.Habit) error
	GetHabitByID(id uint) (*models.Habit, error)
	GetHabitsByHabitListID(habitListID uint) ([]models.Habit, error)
	UpdateHabit(habit *models.Habit) error
	DeleteHabit(id uint) error
}

type habitService struct {
	habitRepo      repository.HabitRepository
	habitListRepo  repository.HabitListRepository
}

func NewHabitService(habitRepo repository.HabitRepository, habitListRepo repository.HabitListRepository) HabitService {
	return &habitService{habitRepo: habitRepo, habitListRepo: habitListRepo}
}

func (s *habitService) CreateHabit(habit *models.Habit) error {
	if err := validators.ValidateHabit(habit); err != nil {
		return err
	}
	// Check if the habit list exists
	_, err := s.habitListRepo.FindByID(habit.HabitListID)
	if err != nil {
		return errors.New("referenced habit list does not exist")
	}
	return s.habitRepo.Create(habit)
}

func (s *habitService) GetHabitByID(id uint) (*models.Habit, error) {
	return s.habitRepo.FindByID(id)
}

func (s *habitService) GetHabitsByHabitListID(habitListID uint) ([]models.Habit, error) {
	return s.habitRepo.FindByHabitListID(habitListID)
}

func (s *habitService) UpdateHabit(habit *models.Habit) error {
	if err := validators.ValidateHabit(habit); err != nil {
		return err
	}
	return s.habitRepo.Update(habit)
}

func (s *habitService) DeleteHabit(id uint) error {
	_, err := s.habitRepo.FindByID(id)
	if err != nil {
		return errors.New("habit not found")
	}
	return s.habitRepo.Delete(id)
} 