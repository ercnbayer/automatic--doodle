package rest

import (
	authMiddleWare "automatic-doodle/modules/auth/middleware"
	authService "automatic-doodle/modules/auth/service"
	UserService "automatic-doodle/modules/user/service"
	"sync"
)

type Rest struct {
	AuthenticationMiddleware authMiddleWare.Middleware
	UserService              UserService.UserService
	AuthService              authService.Service
}

var (
	module     Rest
	moduleOnce sync.Once
)

func New(
	authenticationMiddleware authMiddleWare.Middleware,
) *Rest {
	moduleOnce.Do(func() {
		module = Rest{}
	})

	return &module
}
