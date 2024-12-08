package router

import "automatic-doodle/types"

type UserHandler interface {
	GetRoutes() []types.HandlerItem
}
