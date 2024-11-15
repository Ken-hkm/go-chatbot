package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-chatbot/api"
	"go-chatbot/internal/controller"
	"go-chatbot/internal/db"
	"go-chatbot/internal/handler"
	"go-chatbot/internal/repository"
	"go-chatbot/internal/service"
	"go-chatbot/internal/websocket"
	"log"
	"net/http"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	// Initialize the database
	db.InitDB()
	database := db.DB

	// Set up repositories
	userRepo := repository.NewUserRepository(database)

	// Initialize the JWT secret from environment variables or a default value
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "your-default-jwt-secret"
	}
	// Set up services
	userService := service.NewUserService(userRepo, jwtSecret)

	// Set up controllers
	userController := controller.NewUserController(userService)

	// Set up handlers
	userHandler := handler.NewUserHandler(userController)

	// Initialize WebSocket manager
	manager := websocket.NewManager()
	go manager.Start() // Run the WebSocket manager in a separate goroutine
	// Initialize ChatHandler with the WebSocket manager
	chatHandler := handler.NewChatHandler(manager)
	//init echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//e.Use(internalMiddleware.AuthMiddleware) // Custom middleware

	// Set up request validation
	e.Validator = &CustomValidator{validator: validator.New()}

	// Register routes
	api.RegisterRoutes(e, userHandler, chatHandler)
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	// Start server
	// Get the host and port from environment variables, default to "0.0.0.0:8080" if not set
	host := os.Getenv("SERVER_ADDRESS")
	if host == "" {
		host = "0.0.0.0" // Default to all interfaces if not set
	}
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080" // Default to port 8080 if not set
	}
	log.Printf("Starting server on port %s...", port)
	if err := e.Start(host + ":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// CustomValidator is a custom request validator that uses go-playground/validator
type CustomValidator struct {
	validator *validator.Validate
}

// Validate is a method to validate requests using the CustomValidator
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
