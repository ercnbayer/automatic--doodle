package router

import "automatic-doodle/types"

type UserHandler interface {
	GetRoutes() []types.HandlerItem
}

type AuthHandler interface {
	GetRoutes() []types.HandlerItem
}
type ProfileHandler interface {
	GetRoutes() []types.HandlerItem
}
type FileHandler interface {
	GetRoutes() []types.HandlerItem
}
type JobHandler interface {
	GetRoutes() []types.HandlerItem
}

type JobApplHandler interface {
	GetRoutes() []types.HandlerItem
}

type WebsocketHandler interface {
	GetRoutes() []types.HandlerItem
}
