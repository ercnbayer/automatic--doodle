package rest

import (
	"automatic-doodle/types"

	"github.com/gofiber/fiber/v2"
)

func (r *Rest) GetRoutes() []types.HandlerItem {

	return []types.HandlerItem{
		types.HandlerItem{
			Method: "POST",
			Path:   "/jobAppl",
			Handler: []func(*fiber.Ctx) error{
				r.authMiddleware.Auth,
				r.CreateJobAppl,
			},
		},
		types.HandlerItem{
			Method: "GET",
			Path:   "/GetJobAppl/:id",
			Handler: []func(*fiber.Ctx) error{
				r.authMiddleware.Auth,
				r.GetJobAppl,
			},
		},
	}
}
