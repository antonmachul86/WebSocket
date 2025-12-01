package repository

import (
	"WebSocket/pkg/entity"
)

type MessageRepository interface {
	Publish(message *entity.Message) error
	Subscribe(topic string, handler func(message *entity.Message)) error
}
