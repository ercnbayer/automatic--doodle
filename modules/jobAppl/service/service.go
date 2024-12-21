package service

import "sync"

type Service struct {
	jobApplFactory    JobApplFactory
	jobApplRepository JobApplRepository
	jobRepository     JobRepository
}

var (
	module     Service
	moduleOnce sync.Once
)

func New(jobApplFactory JobApplFactory, jobApplRepo JobApplRepository, jobRepository JobRepository) *Service {
	moduleOnce.Do(func() {
		module = Service{
			jobApplFactory:    jobApplFactory,
			jobApplRepository: jobApplRepo,
			jobRepository:     jobRepository,
		}
	})
	return &module
}
