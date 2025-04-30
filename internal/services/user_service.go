package services

import (
	"errors"
	"fmt"
	"go-rest-project/internal/models"
	"go-rest-project/internal/repository"
	"go-rest-project/internal/validators"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserNotFound       = errors.New("user not found")
	ErrUsernameExists     = errors.New("username already exists")
	ErrEmailExists        = errors.New("email already exists")
	ErrInvalidInput       = errors.New("invalid input")
)

type UserService interface {
	Register(user *models.User) error
	Login(username, password string) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}


func (s *userService) Register(user *models.User) error {
	if err := validators.ValidateUser(user); err != nil {
		return err
	}

	
	existingUser, _ := s.userRepo.FindByUsername(user.Username)
	if existingUser != nil {
		return ErrUsernameExists
	}

	
	existingUser, _ = s.userRepo.FindByEmail(user.Email)
	if existingUser != nil {
		return ErrEmailExists
	}

	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	user.Password = string(hashedPassword)
	user.Role = "user" 

	return s.userRepo.Create(user)
}

func (s *userService) Login(username, password string) (*models.User, error) {
	if strings.TrimSpace(username) == "" || strings.TrimSpace(password) == "" {
		return nil, ErrInvalidCredentials
	}

	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	
	now := time.Now()
	user.LastLoginAt = &now
	if err := s.userRepo.Update(user); err != nil {
		
		fmt.Printf("Failed to update last login time: %v\n", err)
	}

	return user, nil
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	if id == 0 {
		return nil, fmt.Errorf("%w: invalid user ID", ErrInvalidInput)
	}

	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (s *userService) UpdateUser(user *models.User) error {
	if user.ID == 0 {
		return fmt.Errorf("%w: user ID is required", ErrInvalidInput)
	}

	
	existingUser, err := s.userRepo.FindByID(user.ID)
	if err != nil {
		return ErrUserNotFound
	}

	
	if user.Username != "" && user.Username != existingUser.Username {
		if u, _ := s.userRepo.FindByUsername(user.Username); u != nil {
			return ErrUsernameExists
		}
	}

	
	if user.Email != "" && user.Email != existingUser.Email {
		if u, _ := s.userRepo.FindByEmail(user.Email); u != nil {
			return ErrEmailExists
		}
	}

	
	if user.Password != "" {
		// Если пароль меняется, обязательно хешируем его
		if len(user.Password) < 6 {
			return fmt.Errorf("password must be at least 6 characters long")
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("failed to hash password: %w", err)
		}
		user.Password = string(hashedPassword)
	} else {
		// Если пароль не меняется, оставляем старый хеш
		user.Password = existingUser.Password
	}

	
	user.Role = existingUser.Role
	user.LastLoginAt = existingUser.LastLoginAt

	return s.userRepo.Update(user)
}

func (s *userService) DeleteUser(id uint) error {
	if id == 0 {
		return fmt.Errorf("%w: invalid user ID", ErrInvalidInput)
	}

	
	if _, err := s.userRepo.FindByID(id); err != nil {
		return ErrUserNotFound
	}

	return s.userRepo.Delete(id)
}
