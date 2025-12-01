package usecase

import (
	"WebSocket/pkg/entity"
	"WebSocket/pkg/repository/mock"
	"errors"
	"testing"
)

func TestMessageUseCase_ProcessAndDeliverMessage(t *testing.T) {
	t.Run("should return error when repository returns error", func(t *testing.T) {
		repo := &mock.MessageRepository{
			PublishFunc: func(message *entity.Message) error {
				return errors.New("some error")
			},
		}
		uc := NewMessageUseCase(repo)
		err := uc.ProcessAndDeliverMessage(&entity.Message{})
		if err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("should not return error when repository does not return error", func(t *testing.T) {
		repo := &mock.MessageRepository{
			PublishFunc: func(message *entity.Message) error {
				return nil
			},
		}
		uc := NewMessageUseCase(repo)
		err := uc.ProcessAndDeliverMessage(&entity.Message{})
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
	})
}

func TestMessageUseCase_SubscribeToMessages(t *testing.T) {
	t.Run("should return error when repository returns error", func(t *testing.T) {
		repo := &mock.MessageRepository{
			SubscribeFunc: func(topic string, handler func(message *entity.Message)) error {
				return errors.New("some error")
			},
		}
		uc := NewMessageUseCase(repo)
		err := uc.SubscribeToMessages("some-topic", nil)
		if err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("should not return error when repository does not return error", func(t *testing.T) {
		repo := &mock.MessageRepository{
			SubscribeFunc: func(topic string, handler func(message *entity.Message)) error {
				return nil
			},
		}
		uc := NewMessageUseCase(repo)
		err := uc.SubscribeToMessages("some-topic", nil)
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
	})
}
