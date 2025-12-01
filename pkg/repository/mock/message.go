package mock

import "WebSocket/pkg/entity"

type MessageRepository struct {
	PublishFunc       func(message *entity.Message) error
	SubscribeFunc     func(topic string, handler func(message *entity.Message)) error
	PublishedMessages chan *entity.Message // New channel to capture published messages
}

func (m *MessageRepository) Publish(message *entity.Message) error {
	if m.PublishedMessages != nil {
		m.PublishedMessages <- message
	}
	return m.PublishFunc(message)
}

func (m *MessageRepository) Subscribe(topic string, handler func(message *entity.Message)) error {
	return m.SubscribeFunc(topic, handler)
}
