// routes.go
package api

import (
	"github.com/labstack/echo/v4"
	"go-chatbot/internal/handler"
)

func RegisterRoutes(e *echo.Echo, userHandler *handler.UserHandler, chatHandler *handler.ChatHandler) {
	apiGroupUser := e.Group("/api/v1/user")
	apiGroupUser.POST("/login", userHandler.Login)
	apiGroupUser.POST("/register", userHandler.RegisterUser)

	apiGroupChat := e.Group("/api/v1/chat")
	apiGroupChat.GET("/ws/chat/:userID", chatHandler.HandleWebSocket)

	apiGroupVector := e.Group("/api/v1/vector-db")
	apiGroupVector.POST("/populate-data", userHandler.Login)
}
