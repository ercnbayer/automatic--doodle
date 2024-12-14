package di

import (
	accessTokenService "automatic-doodle/modules/access-token/service"
	"automatic-doodle/pkg/config"
	"automatic-doodle/pkg/encryption"
	"automatic-doodle/pkg/logger"
	"automatic-doodle/pkg/postgres"
	s3client "automatic-doodle/pkg/s3_client"
	"automatic-doodle/pkg/server"

	"github.com/google/wire"
)

var ConfigProviderSet wire.ProviderSet = wire.NewSet(
	config.New,

	wire.Bind(
		new(accessTokenService.ConfigModule),
		new(*config.ConfigModule),
	),
	wire.Bind(
		new(encryption.ConfigModule),
		new(*config.ConfigModule),
	),
	wire.Bind(
		new(postgres.ConfigModule),
		new(*config.ConfigModule),
	),
	wire.Bind(
		new(server.ConfigModule),
		new(*config.ConfigModule),
	),
	wire.Bind(new(s3client.ConfigModule), new(*config.ConfigModule)),

	wire.InterfaceValue(
		new(config.Logger),
		logger.New("ConfigModule"),
	),
)
