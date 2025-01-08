package di

import (
	authService "automatic-doodle/modules/auth/service"
	refreshTokenFactory "automatic-doodle/modules/refresh-token/factory"
	userFactory "automatic-doodle/modules/user/factory"
	"automatic-doodle/pkg/encryption"

	logger "github.com/Wodemy-Labs/crawl"

	"github.com/google/wire"
)

var EncryptionProviderSet wire.ProviderSet = wire.NewSet(
	encryption.New,

	wire.Bind(
		new(authService.Encryption),
		new(*encryption.EncryptionModule),
	),
	wire.Bind(
		new(userFactory.Encryption),
		new(*encryption.EncryptionModule),
	),
	wire.Bind(new(refreshTokenFactory.Encryption), new(*encryption.EncryptionModule)),
	wire.InterfaceValue(
		new(encryption.Logger),
		logger.New("EncryptionModule", nil),
	),
)
