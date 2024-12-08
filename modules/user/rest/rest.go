package rest

import (
	"sync"
)

type Rest struct {
	AuthenticationMiddleware AuthenticationMiddleware
}

var (
	module     Rest
	moduleOnce sync.Once
)

func New(
	authenticationMiddleware AuthenticationMiddleware,
) *Rest {
	moduleOnce.Do(func() {
		module = Rest{
			AuthenticationMiddleware: authenticationMiddleware,
		}
	})

	return &module
}
