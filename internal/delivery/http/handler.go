package http

import (
	"WebSocket/pkg/entity"
	"WebSocket/pkg/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"log"
)

type WebSocketHandler struct {
	uc  *usecase.MessageUseCase
	hub *Hub
}

func NewWebSocketHandler(uc *usecase.MessageUseCase, hub *Hub) *WebSocketHandler {
	return &WebSocketHandler{
		uc:  uc,
		hub: hub,
	}
}

func (h *WebSocketHandler) Upgrade() fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
		token := c.Query("token")
		if token == "" {
			log.Println("Authorization failed: no token")
			c.Close()
			return
		}

		if token != "secret" {
			log.Println("Authorization failed: invalid token")
			c.Close()
			return
		}

		userID := c.Query("user_id")
		if userID == "" {
			log.Println("user_id is required")
			c.Close()
			return
		}

		client := &Client{Conn: c, Send: make(chan *entity.Message, 256)}
		h.hub.register <- client

		go client.WritePump()
		h.readPump(client)
	})
}

func (h *WebSocketHandler) readPump(client *Client) {
	defer func() {
		h.hub.unregister <- client
		client.Conn.Close()
	}()
	for {
		var message entity.Message
		if err := client.Conn.ReadJSON(&message); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message.SenderID = client.Conn.Query("user_id")
		if err := h.uc.ProcessAndDeliverMessage(&message); err != nil {
			log.Println("publish:", err)
		}
	}
}
