package handler

import (
	gws "github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"go-chatbot/internal/websocket"
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
