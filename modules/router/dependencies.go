package router

import "automatic-doodle/types"

type UserHandler interface {
	GetRoutes() []types.HandlerItem
}

type AuthHandler interface {
	GetRoutes() []types.HandlerItem
}
