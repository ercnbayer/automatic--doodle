package rest

import (
	"automatic-doodle/pkg/errors"
	"automatic-doodle/types"

	"github.com/gofiber/fiber/v2"
)

func (r *Rest) Me(c *fiber.Ctx) error {

	pUser := c.Locals("user")

	if pUser == nil {
		return errors.New("auth me page", "null credits.")
	}

	userItem, ok := pUser.(types.AuthenticatedUser)

	if !ok {
		return errors.NewUnauthorizedError("Me page auth fail")
	}

	return c.Status(200).JSON(userItem)
}
