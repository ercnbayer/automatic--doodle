package di

import (
	JobFactory "automatic-doodle/modules/job/factory"
	JobRepository "automatic-doodle/modules/job/repository"
	"automatic-doodle/modules/job/rest"
	"automatic-doodle/modules/job/service"
	JobApplService "automatic-doodle/modules/jobAppl/service"
	"automatic-doodle/modules/router"
	"automatic-doodle/pkg/logger"

	"github.com/google/wire"
)

var JobFactoryProviderSet wire.ProviderSet = wire.NewSet(JobFactory.New,
	wire.Bind(new(rest.JobFactory), new(*JobFactory.Factory)))

var JobRepositoryProviderSet wire.ProviderSet = wire.NewSet(JobRepository.New,
	wire.Bind(new(service.JobRepository), new(*JobRepository.Repository)),
	wire.Bind(new(JobApplService.JobRepository), new(*JobRepository.Repository)),
	wire.Bind(new(rest.JobRepository), new(*JobRepository.Repository)))

var JobServiceProviderSet wire.ProviderSet = wire.NewSet(service.New, wire.Bind(new(rest.JobService), new(*service.Service)))
var JobRestProviderSet wire.ProviderSet = wire.NewSet(rest.New, (wire.Bind(new(router.JobHandler), new(*rest.Rest))),
	wire.InterfaceValue(new(rest.Logger), logger.New("JOB API LOGGER")))
