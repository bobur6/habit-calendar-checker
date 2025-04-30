package services

import (
	"go-rest-project/internal/models"
	"go-rest-project/internal/repository"
	"go-rest-project/internal/validators"
	"time"
)

type HabitCheckService interface {
	CreateHabitCheck(check *models.HabitCheck) error
	GetHabitCheckByID(id uint) (*models.HabitCheck, error)
	GetHabitChecksByHabitID(habitID uint) ([]models.HabitCheck, error)
	GetHabitCheckByHabitIDAndDate(habitID uint, date time.Time) (*models.HabitCheck, error)
	UpdateHabitCheck(check *models.HabitCheck) error
	DeleteHabitCheck(id uint) error
}

type habitCheckService struct {
	habitCheckRepo repository.HabitCheckRepository
}

func NewHabitCheckService(habitCheckRepo repository.HabitCheckRepository) HabitCheckService {
	return &habitCheckService{habitCheckRepo: habitCheckRepo}
}

func (s *habitCheckService) CreateHabitCheck(check *models.HabitCheck) error {
	if err := validators.ValidateHabitCheck(check); err != nil {
		return err
	}
	return s.habitCheckRepo.Create(check)
}

func (s *habitCheckService) GetHabitCheckByID(id uint) (*models.HabitCheck, error) {
	return s.habitCheckRepo.FindByID(id)
}

func (s *habitCheckService) GetHabitChecksByHabitID(habitID uint) ([]models.HabitCheck, error) {
	return s.habitCheckRepo.FindByHabitID(habitID)
}

func (s *habitCheckService) GetHabitCheckByHabitIDAndDate(habitID uint, date time.Time) (*models.HabitCheck, error) {
	return s.habitCheckRepo.FindByHabitIDAndDate(habitID, date)
}

func (s *habitCheckService) UpdateHabitCheck(check *models.HabitCheck) error {
	return s.habitCheckRepo.Update(check)
}

func (s *habitCheckService) DeleteHabitCheck(id uint) error {
	return s.habitCheckRepo.Delete(id)
}
