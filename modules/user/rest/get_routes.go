package rest

import (
	"automatic-doodle/types"

	"github.com/gofiber/fiber/v2"
)

func (r *Rest) GetRoutes() []types.HandlerItem {
	return []types.HandlerItem{
		{
			Path:   "user/delete",
			Method: "DELETE",
			Handler: []func(*fiber.Ctx) error{
				r.AuthenticationMiddleware.Auth,
				r.DeleteUser,
			},
		},
	}
}
