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
		AuthenticationServiceProviderSet,
		AuthenticationMiddlewareProviderSet,
		UserRestProviderSet,
		AuthenticationRestProviderSet,
		FileFactoryProvider,
		FileRepositoryProvider,
		FileServiceProviderSet,
		S3ClientProviderSet,
		ProfilePhotoRestProvider,
		RouterProviderSet,
		ServerProviderSet,
	))
}
