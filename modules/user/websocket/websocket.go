package websocket

import "sync"

type Websocket struct {
	authenticationMiddleware AuthenticationMiddleware
	userRepository           UserRepository
	hub                      *Hub
	websocketService         WebsocketService
}

var (
	module     Websocket
	moduleOnce sync.Once
)

func New(
	authenticationMiddleware AuthenticationMiddleware,
	userRepository UserRepository,
	websocketService WebsocketService,

) *Websocket {
	moduleOnce.Do(func() {
		module = Websocket{
			authenticationMiddleware: authenticationMiddleware,
			userRepository:           userRepository,
			hub:                      NewHub(websocketService),
		}
		go module.hub.Run()
	})

	return &module
}
