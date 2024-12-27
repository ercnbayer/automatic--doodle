package rest

import (
	"sync"
)

type Rest struct {
	authenticationMiddleware AuthenticationMiddleware
	userRepository           UserRepository
}

var (
	module     Rest
	moduleOnce sync.Once
)

func New(
	authenticationMiddleware AuthenticationMiddleware,
	userRepository UserRepository,
) *Rest {
	moduleOnce.Do(func() {
		module = Rest{
			authenticationMiddleware: authenticationMiddleware,
			userRepository:           userRepository,
		}
	})

	return &module
}
