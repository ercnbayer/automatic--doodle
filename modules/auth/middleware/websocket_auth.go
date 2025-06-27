package middleware

import (
	"automatic-doodle/pkg/errors"

	"github.com/gofiber/fiber/v2"
)

func (mw *Middleware) WebSocketAuth(c *fiber.Ctx) error {
	// Cookie'den token al
	token := c.Cookies("token")

	if token == "" {
		return errors.NewUnauthorizedError("auth err: token not found in cookie")
	}

	pUser, err := mw.authenticationService.GetUserByToken(&token)
	if err != nil {
		return errors.NewUnauthorizedError("auth err: invalid token")
	}
	c.Locals("pUser", pUser)

	return c.Next()
}
