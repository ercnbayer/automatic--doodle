package websocket

import "sync"

type Websocket struct {
	authenticationMiddleware AuthenticationMiddleware
	userRepository           UserRepository
	hub                      *Hub
}

var (
	module     Websocket
	moduleOnce sync.Once
)

func New(
	authenticationMiddleware AuthenticationMiddleware,
	userRepository UserRepository,
) *Websocket {
	moduleOnce.Do(func() {
		module = Websocket{
			authenticationMiddleware: authenticationMiddleware,
			userRepository:           userRepository,
			hub:                      NewHub(),
		}
		go module.hub.Run()
	})

	return &module
}
