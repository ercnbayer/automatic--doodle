package service

import (
	"sync"
)

type Service struct {
	log                Logger
	encryptionService  Encryption
	accessTokenService AccessTokenService
	userFactory        UserFactory
	userRepository     UserRepository
}

var (
	service     Service
	serviceOnce sync.Once
)

func New(
	log Logger,
	encryptionService Encryption,
	accessTokenService AccessTokenService,

	userFactory UserFactory,
	userRepository UserRepository,
) *Service {
	serviceOnce.Do(func() {
		service = Service{
			log:                log,
			encryptionService:  encryptionService,
			accessTokenService: accessTokenService,
			userFactory:        userFactory,
			userRepository:     userRepository,
		}
	})

	return &service
}
