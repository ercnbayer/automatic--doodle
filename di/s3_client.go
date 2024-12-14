package di

import (
	"automatic-doodle/pkg/logger"

	s3client "automatic-doodle/pkg/s3_client"

	fileService "automatic-doodle/modules/file/service"

	"github.com/google/wire"
)

var S3ClientProviderSet wire.ProviderSet = wire.NewSet(
	s3client.New,

	wire.Bind(
		new(fileService.S3Client),
		new(*s3client.Service),
	),

	wire.InterfaceValue(
		new(s3client.Logger),
		logger.New("S3Client"),
	),
)
