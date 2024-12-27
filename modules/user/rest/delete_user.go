package rest

import (
	"automatic-doodle/pkg/errors"
	"automatic-doodle/types"
	"context"

	"github.com/gofiber/fiber/v2"
)

func (r *Rest) DeleteUser(c *fiber.Ctx) error {
	pLocalUser := c.Locals("user")

	if pLocalUser == nil {

		return errors.New("Delete User", "NULL USER INFO")
	}

	userItems, ok := pLocalUser.(types.AuthenticatedUser)

	if !ok {
		return errors.NewUnauthorizedError("Delete User,USER Items caps error")
	}

	err := r.userRepository.DeleteUser(userItems.Id, context.Background())

	if err != nil {
		return c.Status(400).JSON(err)
	}

	return c.Status(200).JSON("OK")
}
