package ws

import (
	"go-vit/internal/pkg/logfmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	clientConn *websocket.Conn
	mu         sync.Mutex
}

type WebSocketClient interface {
	AddClient(*Client)
	Send(data []byte)
}

var wsUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func InitClient(w http.ResponseWriter, r *http.Request) *Client {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		logfmt.Error.Println("Failed to establish websocket")
		return nil
	}
	return &Client{clientConn: conn}
}

func (c *Client) Send(data []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	err := c.clientConn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		logfmt.Error.Println("Failed to write websocket message payload")
		c.clientConn.Close()
	}
}

func (c *Client) WaitForClose() {
	for {
		_, _, err := c.clientConn.ReadMessage()
		if err != nil {
			logfmt.Error.Println("Client disconnected")
			c.clientConn.Close()
			return
		}
	}
}
