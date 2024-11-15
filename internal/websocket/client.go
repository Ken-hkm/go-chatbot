package websocket

import (
	"github.com/gorilla/websocket"
	"log"
)

func (c *Client) Read(manager *Manager) {
	defer func() {
		manager.Unregister <- c
		c.Conn.Close()
	}()
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}
		manager.Broadcast <- message
	}
}

func (c *Client) Write() {
	defer c.Conn.Close()
	for message := range c.Send {
		err := c.Conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("write error:", err)
			break
		}
	}
}
