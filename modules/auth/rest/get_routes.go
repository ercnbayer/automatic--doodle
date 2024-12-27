package rest

import (
	"automatic-doodle/types"

	"github.com/gofiber/fiber/v2"
)

func (r *Rest) GetRoutes() []types.HandlerItem {
	return []types.HandlerItem{
		{
			Path:   "/auth/login",
			Method: "POST",
			Handler: []func(*fiber.Ctx) error{
				r.loginUser,
			},
		},
		{
			Path:   "/auth/register",
			Method: "POST",
			Handler: []func(*fiber.Ctx) error{

				r.Register,
			},
		},
		{
			Path:   "/auth/me",
			Method: "GET",
			Handler: []func(*fiber.Ctx) error{

				r.authenticationMiddleware.Auth,
				r.Me,
			},
		},

		{
			Path:   "/auth/me",
			Method: "PATCH",
			Handler: []func(*fiber.Ctx) error{

				r.authenticationMiddleware.Auth,
				r.UpdatePassword,
			},
		},
	}

}
