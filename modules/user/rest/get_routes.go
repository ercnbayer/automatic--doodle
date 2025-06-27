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
				r.authenticationMiddleware.Auth,
				r.DeleteUser,
			},
		},
		{
			Path:   "user/view/:id",
			Method: "GET",
			Handler: []func(*fiber.Ctx) error{
				r.authenticationMiddleware.Auth,
				r.ViewUser,
			},
		},
		{
			Path:   "user/all",
			Method: "GET",
			Handler: []func(*fiber.Ctx) error{
				r.GetAllUsers,
			},
		},
		{
			Path:   "user/update",
			Method: "PATCH",
			Handler: []func(*fiber.Ctx) error{
				r.authenticationMiddleware.Auth,
				r.UpdateUser,
			},
		},
	}
}
