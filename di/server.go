package di

import (
	"automatic-doodle/pkg/server"

	logger "github.com/Wodemy-Labs/crawl"
	"github.com/google/wire"
)

var ServerProviderSet wire.ProviderSet = wire.NewSet(
	server.New,

	wire.InterfaceValue(new(server.Logger), logger.New("ServerModule", nil)),
)
