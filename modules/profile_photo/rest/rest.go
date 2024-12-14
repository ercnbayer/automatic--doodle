package rest

import "sync"

var (
	module     Rest
	moduleOnce sync.Once
)

type Rest struct {
	authMiddleware AuthMiddleware
	userService    UserService
	logger         Logger
}

func New(authMiddleware AuthMiddleware, userService UserService, logger Logger) *Rest {
	moduleOnce.Do(func() {
		module = Rest{
			authMiddleware: authMiddleware,
			userService:    userService,
			logger:         logger,
		}
	})
	return &module
}
