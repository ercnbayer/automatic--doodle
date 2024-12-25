package service

import "sync"

type Service struct {
	jobApplFactory    JobApplFactory
	jobApplRepository JobApplRepository
	jobRepository     JobRepository
	userRepository    UserRepository
}

var (
	module     Service
	moduleOnce sync.Once
)

func New(jobApplFactory JobApplFactory, jobApplRepo JobApplRepository, jobRepository JobRepository, userRepository UserRepository) *Service {
	moduleOnce.Do(func() {
		module = Service{
			jobApplFactory:    jobApplFactory,
			jobApplRepository: jobApplRepo,
			jobRepository:     jobRepository,
			userRepository:    userRepository,
		}
	})
	return &module
}
