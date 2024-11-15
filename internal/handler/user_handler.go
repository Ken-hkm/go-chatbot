package handler

import (
	"github.com/labstack/echo/v4"
	"go-chatbot/internal/controller"
	"go-chatbot/internal/db/models"
	"net/http"
)

type UserHandler struct {
	userController controller.UserController
}

// NewUserHandler creates a new instance of UserHandler
func NewUserHandler(userController controller.UserController) *UserHandler {
	return &UserHandler{userController: userController}
}

// RegisterRequest struct represents the JSON structure of the registration request
type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

// RegisterUser handles the registration of a new user
func (h *UserHandler) RegisterUser(c echo.Context) error {
	var req RegisterRequest

	// Bind and validate request
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Map request to the User model
	user := &models.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	// Call the controller to register the user
	if err := h.userController.RegisterUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to register user"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "User registered successfully"})
}

// LoginRequest represents the payload for login
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Login handles user login
func (h *UserHandler) Login(c echo.Context) error {
	var loginReq LoginRequest
	if err := c.Bind(&loginReq); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request"})
	}

	if err := c.Validate(&loginReq); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "validation failed"})
	}

	token, err := h.userController.Login(loginReq.Username, loginReq.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "invalid credentials"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
