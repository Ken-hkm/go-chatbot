package handler

import (
	"bytes"
	gws "github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"go-chatbot/internal/auth"
	"go-chatbot/internal/websocket"
	"io"
	"log"
	"net/http"
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

	// Upgrade the connection to WebSocket
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println("Failed to upgrade to WebSocket:", err)
		return err
	}
	defer conn.Close()

	// Create a new client and register it with the manager
	client := &websocket.Client{
		ID:   c.Param("userID"),
		Conn: conn,
		Send: make(chan []byte, 2), // Buffered channel with a size of 10
	}
	h.manager.Register <- client

	// Start goroutines to read and write WebSocket messages concurrently
	go client.Read(h.manager) // Read messages in a separate goroutine
	go client.Write()         // Write messages in a separate goroutine

	log.Println("WebSocket connection established for user:", userIDFromURL)
	// Process incoming chat messages
	for {
		log.Println("Waiting for message...")
		// Read the next message
		messageType, reader, err := conn.NextReader()
		if err != nil {
			log.Printf("Read error or connection closed: %v", err)
			break
		}
		// Check the message type (TextMessage or BinaryMessage)
		if messageType != gws.TextMessage {
			log.Println("Non-text message received, ignoring.")
			continue
		}

		// Read the entire message payload (handling fragmented frames)
		var messageBuffer bytes.Buffer
		if _, err := io.Copy(&messageBuffer, reader); err != nil {
			log.Printf("Error reading message payload: %v", err)
			continue
		}

		message := messageBuffer.String()

		log.Printf("Received message from user %s: %s", userIDFromURL, message)

		// Generate a fresh context for each request and send it to the LLM
		go func(msg string) {
			aiResponse, err := h.getLLMResponse(c, msg)
			if err != nil {
				log.Println("Error getting AI response:", err)
				return
			}
			log.Println("AI Response:", aiResponse)
			client.Send <- []byte(aiResponse)
			log.Println("done send response")
		}(string(message))
	}
	return nil
}
