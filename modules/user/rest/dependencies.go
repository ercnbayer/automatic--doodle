package rest

import (
	"github.com/gofiber/fiber/v2"
)

type AuthenticationMiddleware interface {
	Auth(c *fiber.Ctx) error
}
