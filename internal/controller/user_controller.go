// user_controller.go
package controller

import (
	"go-chatbot/internal/db/models"
	"go-chatbot/internal/service"
)

type UserController interface {
	RegisterUser(user *models.User) error
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{userService: userService}
}

func (uc *userController) RegisterUser(user *models.User) error {
	return uc.userService.RegisterUser(user)
}
