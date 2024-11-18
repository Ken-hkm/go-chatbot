package websocket

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

// Manager manages WebSocket connections for multiple users.
type Manager struct {
	clients map[string]*websocket.Conn // Map userID -> WebSocket connection
	lock    sync.Mutex                 // Synchronize access to the clients map
}

// NewManager creates a new WebSocket Manager.
func NewManager() *Manager {
	return &Manager{
		clients: make(map[string]*websocket.Conn),
	}
}

// AddClient adds a new WebSocket connection for a user.
func (m *Manager) AddClient(userID string, conn *websocket.Conn) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.clients[userID] = conn
	log.Printf("User %s connected", userID)
}

// RemoveClient removes a user's WebSocket connection.
func (m *Manager) RemoveClient(userID string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if conn, exists := m.clients[userID]; exists {
		conn.Close() // Close the connection
		delete(m.clients, userID)
		log.Printf("User %s disconnected", userID)
	}
}

// SendMessage sends a message to a specific user's WebSocket connection.
func (m *Manager) SendMessage(userID, message string) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	conn, exists := m.clients[userID]
	if !exists {
		return nil // No active connection for the user
	}

	return conn.WriteMessage(websocket.TextMessage, []byte(message))
}
