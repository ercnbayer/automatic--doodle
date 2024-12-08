package di

import (
	authenticationMiddleware "automatic-doodle/modules/auth/middleware"
	authService "automatic-doodle/modules/auth/service"
	userRest "automatic-doodle/modules/user/rest"
	"automatic-doodle/pkg/logger"

	"github.com/google/wire"
)

var (
	AuthenticationServiceProviderSet wire.ProviderSet = wire.NewSet(
		authService.New,
		wire.Bind(
			new(authenticationMiddleware.AuthenticationService),
			new(*authService.Service),
		),

		wire.InterfaceValue(
			new(authService.Logger),
			logger.New("AuthenticationService"),
		))

	AuthenticationMiddlewareProviderSet wire.ProviderSet = wire.NewSet(
		authenticationMiddleware.New,

		wire.Bind(
			new(userRest.AuthenticationMiddleware),
			new(*authenticationMiddleware.Middleware),
		),
		wire.InterfaceValue(new(authenticationMiddleware.Logger), logger.New("Authentication Middleware Logger")),
	)
)
