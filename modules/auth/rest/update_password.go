package rest

import (
	"automatic-doodle/pkg/errors"
	"automatic-doodle/types"

	"github.com/gofiber/fiber/v2"
)

func (r *Rest) UpdatePassword(c *fiber.Ctx) error {

	pUser := c.Locals("user")

	if pUser == nil {
		return errors.New("auth me page", "null credits.")
	}

	userItem, ok := pUser.(types.AuthenticatedUser)

	if !ok {
		return errors.NewUnauthorizedError("Me page auth fail")
	}

	var req types.UpdateUserPassword

	err := c.BodyParser(&req)

	if err != nil {
		return c.Status(400).JSON(err)
	}

	user, err := r.authenticationService.UpdatePassword(req, userItem.Id)

	if err != nil {
		return c.Status(400).JSON(err)
	}
	return c.Status(200).JSON(user)
}
