package rest

import (
	"automatic-doodle/pkg/errors"

	"github.com/gofiber/fiber/v2"
)

// this is whole to do so dont care
func (r *Rest) DeleteUser(c *fiber.Ctx) error {
	pLocalUser := c.Locals("user")

	if pLocalUser == nil {
		// to do throw err
		return errors.New("Delete User", "NULL USER INFO")
	}

	return c.Status(200).JSON("OK")
}
