package service

import "sync"

type Service struct {
	messageRepository MessageRepository
}

var (
	module     Service
	moduleOnce sync.Once
)

func New(messageRepository MessageRepository) *Service {
	moduleOnce.Do(func() {
		module = Service{
			messageRepository: messageRepository,
		}
	})
	return &module
}
