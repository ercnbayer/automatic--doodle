package service

import "sync"

type Service struct {
	jobRepository  JobRepository
	userRepository UserRepository
}

var (
	module     Service
	moduleOnce sync.Once
)

func New(jobRepository JobRepository, userRepository UserRepository) *Service {
	moduleOnce.Do(func() {
		module = Service{
			jobRepository:  jobRepository,
			userRepository: userRepository,
		}
	})
	return &module
}
