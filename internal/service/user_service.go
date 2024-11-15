// user_service.go
package service

import (
	"errors"
	"go-chatbot/internal/db/models"
	"go-chatbot/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) RegisterUser(user *models.User) error {
	existingUser, err := s.userRepo.FindByEmail(user.Email)
	if err == nil && existingUser != nil {
		return errors.New("email already registered")
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return s.userRepo.CreateUser(user)
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.userRepo.FindByID(id)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
