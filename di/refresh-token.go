package di

import (
	authenticationService "automatic-doodle/modules/auth/service"
	RefreshTokenFactory "automatic-doodle/modules/refresh-token/factory"
	RefreshTokenRepository "automatic-doodle/modules/refresh-token/repository"

	"github.com/google/wire"
)

var (
	RefreshTokenFactoryProvider wire.ProviderSet = wire.NewSet(RefreshTokenFactory.New,
		wire.Bind(new(authenticationService.RefreshTokenFactory), new(*RefreshTokenFactory.Factory)))

	RefreshTokenRepositoryProvider wire.ProviderSet = wire.NewSet(RefreshTokenRepository.New, wire.Bind(new(authenticationService.RefreshTokenRepository), new(*RefreshTokenRepository.Repository)))
)
