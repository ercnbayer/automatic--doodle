package service

import "sync"

type Service struct {
	logger     Logger
	repository FileRepository
	factory    FileFactory
	s3Client   S3Client
}

var (
	module     Service
	moduleOnce sync.Once
)

func New(logger Logger, repository FileRepository, factory FileFactory, s3Client S3Client) *Service {
	moduleOnce.Do(func() {

		module =
			Service{
				logger:     logger,
				repository: repository,
				factory:    factory,
				s3Client:   s3Client,
			}
	})
	return &module
}
