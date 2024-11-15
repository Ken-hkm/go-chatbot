// routes.go
package api

import (
	"github.com/labstack/echo/v4"
	"go-chatbot/internal/handler"
)

func RegisterRoutes(e *echo.Echo, userHandler *handler.UserHandler) {
	e.POST("/login", userHandler.Login)
	e.POST("/register", userHandler.RegisterUser)

}
