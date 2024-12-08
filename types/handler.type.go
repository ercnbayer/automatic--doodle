package types

import "github.com/gofiber/fiber/v2"

type HandlerItem struct {
	Handler []func(*fiber.Ctx) error
	Path    string
	Method  string
}
