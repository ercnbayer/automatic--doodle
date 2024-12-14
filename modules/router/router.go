package router

import (
	"automatic-doodle/pkg/utils"
	"automatic-doodle/types"
	"sync"
)

var (
	module     Router
	moduleOnce sync.Once
)

func New(
	userHandlers UserHandler,
	authHandlers AuthHandler,
	profile_photo Profile_PhotoHandler,
	fileHandler FileHandler,
) *Router {
	moduleOnce.Do(func() {
		module = Router{
			Handlers: utils.MergeArrays(
				userHandlers.GetRoutes(),
				authHandlers.GetRoutes(),
				profile_photo.GetRoutes(),
				fileHandler.GetRoutes(),
			),
		}
	})

	return &module
}

type Router struct {
	Handlers []types.HandlerItem
}

func (r *Router) GetRoutes() *[]types.HandlerItem {
	return &r.Handlers
}
