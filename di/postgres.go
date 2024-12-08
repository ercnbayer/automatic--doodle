package di

import (
	"automatic-doodle/pkg/logger"
	"automatic-doodle/pkg/postgres"

	"github.com/google/wire"
)

var PostgresProviderSet wire.ProviderSet = wire.NewSet(
	postgres.New,

	wire.InterfaceValue(
		new(postgres.Logger),
		logger.New("PostgresModule"),
	),
)
