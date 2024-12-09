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
	}
}
