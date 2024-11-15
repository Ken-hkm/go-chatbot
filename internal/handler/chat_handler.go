package handler

import (
	gws "github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"go-chatbot/internal/auth"
	"go-chatbot/internal/websocket"
	"log"
	"net/http"
)

var upgrader = gws.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type ChatHandler struct {
	manager *websocket.Manager // Refers to your custom Manager struct in internal/websocket
}

func NewChatHandler(manager *websocket.Manager) *ChatHandler {
	return &ChatHandler{manager: manager}
}

func (h *ChatHandler) HandleWebSocket(c echo.Context) error {
	// Extract the token from the Authorization header (Bearer token)
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Missing token")
	}
	// Extract userID from the URL or request parameter
	userIDFromURL := c.Param("userID") // Assuming the URL has a userID parameter like /ws/chat/:userID
	if userIDFromURL == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing userID in request path")
	}
	// Validate the JWT token and extract userID from the claims
	userIDFromToken, err := auth.ValidateToken(token)
	if err != nil {
		log.Println("Error validating token:", err)
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error()) // Provide the specific error from ValidateToken
	}

	// Check if the userID from the token matches the userID in the URL path
	if userIDFromURL != userIDFromToken {
		return echo.NewHTTPError(http.StatusUnauthorized, "UserID mismatch")
	}

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil) // Uses the gws (Gorilla WebSocket) upgrader
	if err != nil {
		return err
	}

	client := &websocket.Client{ // Refers to your Client struct in internal/websocket
		ID:   c.Param("userID"),
		Conn: conn,
		Send: make(chan []byte),
	}

	h.manager.Register <- client

	go client.Read(h.manager)
	go client.Write()

	return nil
}
