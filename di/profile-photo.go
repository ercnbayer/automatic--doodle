package di

import (
	ProfileRest "automatic-doodle/modules/profile/rest"

	"automatic-doodle/modules/router"

	logger "github.com/Wodemy-Labs/crawl"
	"github.com/google/wire"
)

var (
	ProfileRestProvider wire.ProviderSet = wire.NewSet(ProfileRest.New, wire.Bind(new(router.ProfileHandler), new(*ProfileRest.Rest)), wire.InterfaceValue(new(ProfileRest.Logger), logger.New("logger", nil)))
)
