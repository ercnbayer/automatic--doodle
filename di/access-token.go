package di

import (
	accessToken "automatic-doodle/modules/access-token/service"
	authService "automatic-doodle/modules/auth/service"
	"automatic-doodle/pkg/logger"

	"github.com/google/wire"
)

var AccessTokenProvider = wire.NewSet(
	accessToken.New,
	wire.Bind(new(authService.AccessTokenService), new(*accessToken.Service)),

	wire.InterfaceValue(new(accessToken.Logger), logger.New("AccessTokenService")),
)
