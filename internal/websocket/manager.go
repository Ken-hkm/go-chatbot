package websocket

import (
	"github.com/gorilla/websocket"
	"sync"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Send chan []byte
}

type Manager struct {
	Clients    map[string]*Client
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
	mu         sync.Mutex
}

func NewManager() *Manager {
	return &Manager{
		Clients:    make(map[string]*Client),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (m *Manager) Start() {
	for {
		select {
		case client := <-m.Register:
			m.mu.Lock()
			m.Clients[client.ID] = client
			m.mu.Unlock()
		case client := <-m.Unregister:
			m.mu.Lock()
			if _, ok := m.Clients[client.ID]; ok {
				close(client.Send)
				delete(m.Clients, client.ID)
			}
			m.mu.Unlock()
		case message := <-m.Broadcast:
			m.mu.Lock()
			for _, client := range m.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(m.Clients, client.ID)
				}
			}
			m.mu.Unlock()
		}
	}
}
