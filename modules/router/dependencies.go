package router

import "automatic-doodle/types"

type UserHandler interface {
	GetRoutes() []types.HandlerItem
}

type AuthHandler interface {
	GetRoutes() []types.HandlerItem
}
type Profile_PhotoHandler interface {
	GetRoutes() []types.HandlerItem
}
type FileHandler interface {
	GetRoutes() []types.HandlerItem
}
