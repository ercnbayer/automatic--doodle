package service

import "sync"

type WebsocketService struct {
	messageRepository MessageRepository
	messageFactory    MessageFactory
}

var (
	module     WebsocketService
	moduleOnce sync.Once
)

func New(
	messageFactory MessageFactory,
	messageRepository MessageRepository,
) *WebsocketService {
	moduleOnce.Do(func() {
		module = WebsocketService{
			messageFactory:    messageFactory,
			messageRepository: messageRepository,
		}
	})
	return &module
}
