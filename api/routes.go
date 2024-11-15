// routes.go
package api

import (
	"github.com/labstack/echo/v4"
	"go-chatbot/internal/handler"
)

func RegisterRoutes(e *echo.Echo, userHandler *handler.UserHandler) {
	apiGroupUser := e.Group("/api/v1/user")

	apiGroupUser.POST("/login", userHandler.Login)
	apiGroupUser.POST("/register", userHandler.RegisterUser)
}
