package di

import (
	"automatic-doodle/modules/router"
	"automatic-doodle/pkg/server"

	"github.com/google/wire"
)

var RouterProviderSet wire.ProviderSet = wire.NewSet(
	router.New,

	wire.Bind(new(server.Router), new(*router.Router)),
)
