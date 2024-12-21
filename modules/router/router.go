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
	profileHandler ProfileHandler,
	fileHandler FileHandler,
	jobHandler JobHandler,
	jobApplHandler JobApplHandler,
) *Router {
	moduleOnce.Do(func() {
		module = Router{
			Handlers: utils.MergeArrays(
				userHandlers.GetRoutes(),
				authHandlers.GetRoutes(),
				profileHandler.GetRoutes(),
				fileHandler.GetRoutes(),
				jobHandler.GetRoutes(),
				jobApplHandler.GetRoutes(),
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
