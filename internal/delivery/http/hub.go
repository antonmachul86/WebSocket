package http

import (
	"WebSocket/pkg/entity"
	"github.com/gofiber/websocket/v2"
	"log"
	"sync"
)

type Client struct {
	Conn *websocket.Conn
	Send chan *entity.Message
}

type Hub struct {
	clients    map[string]*Client
	register   chan *Client
	unregister chan *Client
	broadcast  chan *entity.Message
	mu         sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[string]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan *entity.Message),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client.Conn.Query("user_id")] = client
			h.mu.Unlock()
		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client.Conn.Query("user_id")]; ok {
				delete(h.clients, client.Conn.Query("user_id"))
				close(client.Send)
			}
			h.mu.Unlock()
		case message := <-h.broadcast:
			h.mu.Lock()
			if client, ok := h.clients[message.To]; ok {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.clients, message.To)
				}
			}
			h.mu.Unlock()
		}
	}
}

func (h *Hub) Broadcast(message *entity.Message) {
	h.broadcast <- message
}

func (c *Client) WritePump() {
	defer func() {
		c.Conn.Close()
	}()
	for {
		message, ok := <-c.Send
		if !ok {
			c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}

		if err := c.Conn.WriteJSON(message); err != nil {
			log.Println("write:", err)
			return
		}
	}
}
