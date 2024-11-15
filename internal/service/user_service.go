// user_service.go
package service

import (
	"errors"
	"go-chatbot/internal/auth"
	"go-chatbot/internal/db/models"
	"go-chatbot/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UserService interface {
	RegisterUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	Login(username, password string) (string, error)
}

type userService struct {
	userRepo  repository.UserRepository
	jwtSecret string
}

func NewUserService(userRepo repository.UserRepository, jwtSecret string) UserService {
	return &userService{userRepo: userRepo, jwtSecret: jwtSecret}
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

func (s *userService) Login(username, password string) (string, error) {
	// Check if user exists and password matches
	user, err := s.userRepo.GetUserByUsername(username)
	hashedPass, err := hashPassword(password)

	log.Println(user.Password + " : " + hashedPass)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	tokenString, err := auth.GenerateToken(user)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
