package ws

import (
	"encoding/json"
	"go-vit/internal/mon/models"
	"go-vit/internal/pkg/logfmt"
	"sync"
)

type WebSocketServer interface {
	AddClient(*Client)
	RemoveClient(*Client)
	Broadcast(*models.ResourceStats)
}

type ClientWebSocketServer struct {
	clients map[*Client]bool
	mu      sync.Mutex
}

func InitWebSocketServer() *ClientWebSocketServer {
	return &ClientWebSocketServer{clients: make(map[*Client]bool)}
}

func (s *ClientWebSocketServer) AddClient(client *Client) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if client != nil {
		s.clients[client] = true
	}
}

func (s *ClientWebSocketServer) RemoveClient(client *Client) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.clients, client)
}

func (s *ClientWebSocketServer) Broadcast(stats *models.ResourceStats) {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := json.Marshal(stats)
	if err != nil {
		logfmt.Error.Println("Failed to serialize payload:", err)
		return
	}

	for client := range s.clients {
		client.Send(data)
	}
}
