package di

import (
	"automatic-doodle/modules/jobAppl/factory"
	"automatic-doodle/modules/jobAppl/repository"
	"automatic-doodle/modules/jobAppl/rest"
	"automatic-doodle/modules/jobAppl/service"
	"automatic-doodle/modules/router"
	"automatic-doodle/pkg/logger"

	"github.com/google/wire"
)

var JobApplFactoryProviderSet wire.ProviderSet = wire.NewSet(factory.New, wire.Bind(new(service.JobApplFactory), new(*factory.Factory)))

var JobApplRepositoryProviderSet wire.ProviderSet = wire.NewSet(repository.New, wire.Bind(new(service.JobApplRepository), new(*repository.Repository)))

var JobApplServiceProviderSet wire.ProviderSet = wire.NewSet(service.New, wire.Bind(new(rest.JobApplService), new(*service.Service)))

var JobApplRestProviderSet wire.ProviderSet = wire.NewSet(rest.New, wire.Bind(new(router.JobApplHandler), new(*rest.Rest)), wire.InterfaceValue(new(rest.Logger), logger.New("JobApplRest Module")))
