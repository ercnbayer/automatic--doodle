package di

import (
	accessToken "automatic-doodle/modules/access-token/service"
	authService "automatic-doodle/modules/auth/service"

	logger "github.com/Wodemy-Labs/crawl"

	"github.com/google/wire"
)

var AccessTokenProvider = wire.NewSet(
	accessToken.New,
	wire.Bind(new(authService.AccessTokenService), new(*accessToken.Service)),

	wire.InterfaceValue(new(accessToken.Logger), logger.New("AccessTokenService", nil)),
)
