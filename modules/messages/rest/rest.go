package rest

import "sync"

var (
	module     Rest
	moduleOnce sync.Once
)

type Rest struct {
	messageService           MessageService
	authenticationMiddleware AuthenticationMiddleware
}

func New(
	messageService MessageService,
	authenticationMiddleware AuthenticationMiddleware,
) *Rest {
	moduleOnce.Do(func() {
		module = Rest{
			authenticationMiddleware: authenticationMiddleware,
			messageService:           messageService,
		}
	})
	return &module
}
