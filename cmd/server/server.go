package main

import (
	"WebSocket/internal/delivery/http"
	"WebSocket/internal/repository/natspub"
	"WebSocket/internal/repository/natssub"
	"WebSocket/pkg/entity"
	"WebSocket/pkg/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/nats-io/nats.go"
	"log"
)

func main() {
	// Connect to NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Create repositories
	pubRepo := natspub.NewMessageRepository(nc)
	subRepo := natssub.NewMessageRepository(nc)

	// Create use cases
	messageUseCase := usecase.NewMessageUseCase(pubRepo)

	// Create a new Hub
	hub := http.NewHub()
	go hub.Run()

	// Subscribe to NATS and broadcast to hub
	err = subRepo.Subscribe(">", func(message *entity.Message) {
		hub.Broadcast(message)
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a new Fiber instance
	app := fiber.New()

	// Create WebSocket handler
	wsHandler := http.NewWebSocketHandler(messageUseCase, hub)

	// WebSocket route
	app.Get("/ws", wsHandler.Upgrade())

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
