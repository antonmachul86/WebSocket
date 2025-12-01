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
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	pubRepo := natspub.NewMessageRepository(nc)
	subRepo := natssub.NewMessageRepository(nc)

	messageUseCase := usecase.NewMessageUseCase(pubRepo)

	hub := http.NewHub()
	go hub.Run()

	err = subRepo.Subscribe(">", func(message *entity.Message) {
		hub.Broadcast(message)
	})
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	wsHandler := http.NewWebSocketHandler(messageUseCase, hub)

	app.Get("/ws", wsHandler.Upgrade())

	log.Fatal(app.Listen(":3000"))
}
