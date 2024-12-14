package rest

import (
	"automatic-doodle/types"

	"github.com/gofiber/fiber/v2"
)

func (r *Rest) GetRoutes() []types.HandlerItem {
	return []types.HandlerItem{
		{
			Path:   "/profile-photo",
			Method: "PATCH",
			Handler: []func(*fiber.Ctx) error{
				r.authMiddleware.Auth,
				r.UpdateProfilePhoto,
			},
		},
	}
}
