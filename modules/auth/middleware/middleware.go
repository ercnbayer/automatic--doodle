package middleware

import "sync"

type Middleware struct {
	authenticationService AuthenticationService
	log                   Logger
}

var (
	module     Middleware
	moduleOnce sync.Once
)

func New(
	log Logger,
	authenticationService AuthenticationService,
) *Middleware {
	moduleOnce.Do(func() {
		module = Middleware{
			log:                   log,
			authenticationService: authenticationService,
		}
	})

	return &module
}
