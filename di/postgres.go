package di

import (
	"automatic-doodle/pkg/postgres"

	logger "github.com/Wodemy-Labs/crawl"

	"github.com/google/wire"
)

var PostgresProviderSet wire.ProviderSet = wire.NewSet(
	postgres.New,

	wire.InterfaceValue(
		new(postgres.Logger),
		logger.New("PostgresModule", nil),
	),
)
