package rest

import (
	"automatic-doodle/types"

	"github.com/gofiber/fiber/v2"
)

func (r *Rest) GetRoutes() []types.HandlerItem {
	return []types.HandlerItem{
		{
			Path:   "/files",
			Method: "POST",
			Handler: []func(*fiber.Ctx) error{
				r.authenticationMiddleware.Auth,
				r.CreateUploadUrl,
			},
		},
	}
}
