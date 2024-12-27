package rest

import (
	"automatic-doodle/pkg/errors"
	"automatic-doodle/pkg/validator"
	"automatic-doodle/types"
	"context"

	"github.com/gofiber/fiber/v2"
)

func (r *Rest) UpdateUser(c *fiber.Ctx) error {
	pLocalUser := c.Locals("user")

	if pLocalUser == nil {

		return errors.New("Get User", "NULL USER INFO")
	}

	userItem, ok := pLocalUser.(types.AuthenticatedUser)

	if !ok {
		return errors.NewUnauthorizedError("Get User,USER Items cast error")
	}

	var body types.UpdateUserReq
	err := c.BodyParser(&body)

	if err != nil {
		return c.Status(400).JSON(err)
	}

	err = validator.ValidateStruct(&body)

	if err != nil {
		return c.Status(400).JSON(err)
	}

	user, err := r.userRepository.UpdateUser(context.Background(), body, userItem.Id)

	if err != nil {
		return c.Status(400).JSON(err)
	}

	return c.Status(200).JSON(types.AuthenticatedUserFromUser(user))
}
