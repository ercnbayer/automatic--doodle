package di

import (
	"automatic-doodle/modules/profile_photo/rest"
	Profile_Photo_Rest "automatic-doodle/modules/profile_photo/rest"
	"automatic-doodle/modules/router"
	"automatic-doodle/pkg/logger"

	"github.com/google/wire"
)

var (
	ProfilePhotoRestProvider wire.ProviderSet = wire.NewSet(rest.New, wire.Bind(new(router.Profile_PhotoHandler), new(*Profile_Photo_Rest.Rest)), wire.InterfaceValue(new(Profile_Photo_Rest.Logger), logger.New("logger")))
)
