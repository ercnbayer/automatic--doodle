package di

import (
	"automatic-doodle/ent"
	"automatic-doodle/pkg/server"

	"github.com/google/wire"
)

func DBBuilder() *ent.Client {
	panic(wire.Build(
		ConfigProviderSet,
		PostgresProviderSet,
	))
}

func Wire(
	db *ent.Client,
) *server.Server {
	panic(wire.Build(
		ConfigProviderSet,
		EncryptionProviderSet,
		AccessTokenProvider,

		UserFactoryProviderSet,
		UserRepositoryProviderSet,
		AuthenticationServiceProviderSet,
		AuthenticationMiddlewareProviderSet,
		UserRestProviderSet,

		RouterProviderSet,
		ServerProviderSet,
	))
}
