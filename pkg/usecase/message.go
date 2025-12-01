package usecase

import (
	"WebSocket/pkg/entity"
	"WebSocket/pkg/repository"
)

type MessageUseCase struct {
	repo repository.MessageRepository
}

func NewMessageUseCase(repo repository.MessageRepository) *MessageUseCase {
	return &MessageUseCase{repo: repo}
}

func (uc *MessageUseCase) ProcessAndDeliverMessage(message *entity.Message) error {
	// Here you can add any business logic before publishing the message
	return uc.repo.Publish(message)
}

func (uc *MessageUseCase) SubscribeToMessages(topic string, handler func(message *entity.Message)) error {
	return uc.repo.Subscribe(topic, handler)
}
