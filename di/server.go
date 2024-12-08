package di

import (
	"automatic-doodle/pkg/logger"
	"automatic-doodle/pkg/server"

	"github.com/google/wire"
)

var ServerProviderSet wire.ProviderSet = wire.NewSet(
	server.New,

	wire.InterfaceValue(new(server.Logger), logger.New("ServerModule")),
)
