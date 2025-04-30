package services

import (
	"errors"
	"go-rest-project/internal/models"
	"go-rest-project/internal/repository"
	"go-rest-project/internal/validators"
)

type HabitListService interface {
	CreateHabitList(habitList *models.HabitList) error
	GetHabitListByID(id uint) (*models.HabitList, error)
	GetHabitListsByUserID(userID uint) ([]models.HabitList, error)
	UpdateHabitList(habitList *models.HabitList) error
	DeleteHabitList(id uint) error
}

type habitListService struct {
	habitListRepo repository.HabitListRepository
}

func NewHabitListService(habitListRepo repository.HabitListRepository) HabitListService {
	return &habitListService{habitListRepo: habitListRepo}
}

func (s *habitListService) CreateHabitList(habitList *models.HabitList) error {
	if err := validators.ValidateHabitList(habitList); err != nil {
		return err
	}
	return s.habitListRepo.Create(habitList)
}

func (s *habitListService) GetHabitListByID(id uint) (*models.HabitList, error) {
	return s.habitListRepo.FindByID(id)
}

func (s *habitListService) GetHabitListsByUserID(userID uint) ([]models.HabitList, error) {
	return s.habitListRepo.FindByUserID(userID)
}

func (s *habitListService) UpdateHabitList(habitList *models.HabitList) error {
	if err := validators.ValidateHabitList(habitList); err != nil {
		return err
	}
	_, err := s.habitListRepo.FindByID(habitList.ID)
	if err != nil {
		return errors.New("habit list not found")
	}
	return s.habitListRepo.Update(habitList)
}

func (s *habitListService) DeleteHabitList(id uint) error {
	_, err := s.habitListRepo.FindByID(id)
	if err != nil {
		return errors.New("habit list not found")
	}
	return s.habitListRepo.Delete(id)
} 