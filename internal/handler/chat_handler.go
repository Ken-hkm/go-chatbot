package handler

import (
	gws "github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"go-chatbot/internal/auth"
	"go-chatbot/internal/websocket"
	"log"
	"net/http"
	"os"
)

// WebSocket upgrader settings
var upgrader = gws.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type ChatHandler struct {
	manager *websocket.Manager // Refers to your custom Manager struct in internal/websocket
}

// Create a new ChatHandler with the provided manager
func NewChatHandler(manager *websocket.Manager) *ChatHandler {
	return &ChatHandler{manager: manager}
}

// HandleWebSocket handles the WebSocket connections and processes messages
func (h *ChatHandler) HandleWebSocket(c echo.Context) error {
	// Extract the token and validate as described in your existing code.
	userIDFromToken := "0"
	if os.Getenv("VALIDATE_TOKEN") == "true" {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing token")
		}

		userIDFromURL := c.Param("userID")
		if userIDFromURL == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Missing userID in request path")
		}

		userIDFromToken, err := auth.ValidateToken(token)
		if err != nil {
			log.Println("Error validating token:", err)
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		if userIDFromURL != userIDFromToken {
			return echo.NewHTTPError(http.StatusUnauthorized, "UserID mismatch")
		}
	}

	// Upgrade to WebSocket
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println("Failed to upgrade WebSocket:", err)
		return err
	}
	defer conn.Close()

	// Add the connection to the manager
	h.manager.AddClient(userIDFromToken, conn)
	defer h.manager.RemoveClient(userIDFromToken)

	// Read messages from the WebSocket
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading from WebSocket for user %s: %v", userIDFromToken, err)
			break
		}
		log.Printf("Received message from user %s: %s", userIDFromToken, msg)
		aiResponse, err := h.getLLMResponse(c, string(msg))
		if err := conn.WriteMessage(gws.TextMessage, []byte(aiResponse)); err != nil {
			log.Printf("Error writing to WebSocket for user %s: %v", userIDFromToken, err)
			break
		}
	}

	return nil
}
