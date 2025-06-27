//go:build wireinject
// +build wireinject

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
		RefreshTokenFactoryProvider,
		RefreshTokenRepositoryProvider,
		UserFactoryProviderSet,
		UserRepositoryProviderSet,
		UserRestProviderSet,
		UserServiceProviderSet,
		AuthenticationServiceProviderSet,
		AuthenticationMiddlewareProviderSet,

		AuthenticationRestProviderSet,
		FileFactoryProvider,
		FileRepositoryProvider,
		FileServiceProviderSet,
		FileRestProviderSet,
		JobRepositoryProviderSet,
		JobFactoryProviderSet,
		JobServiceProviderSet,
		JobRestProviderSet,
		JobApplFactoryProviderSet,
		JobApplRepositoryProviderSet,
		JobApplServiceProviderSet,
		JobApplRestProviderSet,
		S3ClientProviderSet,
		ProfileRestProvider,
		WebsocketProviderSet,
		RouterProviderSet,
		ServerProviderSet,
	))
}
