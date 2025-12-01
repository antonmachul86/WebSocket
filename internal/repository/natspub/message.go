package natspub

import (
	"WebSocket/pkg/entity"
	"github.com/nats-io/nats.go"
)

type MessageRepository struct {
	nc *nats.Conn
}

func NewMessageRepository(nc *nats.Conn) *MessageRepository {
	return &MessageRepository{nc: nc}
}

func (r *MessageRepository) Publish(message *entity.Message) error {
	return r.nc.Publish(message.To, message.ToJson())
}

func (r *MessageRepository) Subscribe(topic string, handler func(message *entity.Message)) error {
	// This repository is for publishing only
	return nil
}
