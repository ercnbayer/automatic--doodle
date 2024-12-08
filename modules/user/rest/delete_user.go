package rest

import "github.com/gofiber/fiber/v2"

func (r *Rest) DeleteUser(c *fiber.Ctx) error {
	pLocalUser := c.Locals("user")

	if pLocalUser == nil {
		// to do throw err
	}

	return c.Status(200).JSON("OK")
}
