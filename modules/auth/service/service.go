package service

import (
	"sync"
)

type Service struct {
	log                    Logger
	encryptionService      Encryption
	accessTokenService     AccessTokenService
	userFactory            UserFactory
	userRepository         UserRepository
	refreshTokenRepository RefreshTokenRepository
	refreshTokenFactory    RefreshTokenFactory
}

var (
	service     Service
	serviceOnce sync.Once
)

func New(
	log Logger,
	encryptionService Encryption,
	accessTokenService AccessTokenService,
	refreshTokenRepository RefreshTokenRepository,
	refreshTokenFactory RefreshTokenFactory,
	userFactory UserFactory,
	userRepository UserRepository,
) *Service {
	serviceOnce.Do(func() {
		service = Service{
			log:                    log,
			encryptionService:      encryptionService,
			accessTokenService:     accessTokenService,
			userFactory:            userFactory,
			userRepository:         userRepository,
			refreshTokenRepository: refreshTokenRepository,
			refreshTokenFactory:    refreshTokenFactory}
	})

	return &service
}
