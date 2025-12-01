package natssub

import (
	"WebSocket/pkg/entity"
	"encoding/json"
	"github.com/nats-io/nats.go"
)

type MessageRepository struct {
	nc *nats.Conn
}

func NewMessageRepository(nc *nats.Conn) *MessageRepository {
	return &MessageRepository{nc: nc}
}

func (r *MessageRepository) Publish(message *entity.Message) error {
	// This repository is for subscribing only
	return nil
}

func (r *MessageRepository) Subscribe(topic string, handler func(message *entity.Message)) error {
	_, err := r.nc.Subscribe(topic, func(msg *nats.Msg) {
		var message entity.Message
		if err := json.Unmarshal(msg.Data, &message); err == nil {
			handler(&message)
		}
	})
	return err
}
