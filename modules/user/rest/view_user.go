package rest

import (
	"automatic-doodle/pkg/errors"
	"automatic-doodle/types"
	"context"

	"github.com/gofiber/fiber/v2"
)

func (r *Rest) ViewUser(c *fiber.Ctx) error {
	pLocalUser := c.Locals("user")

	if pLocalUser == nil {

		return errors.New("Get User", "NULL USER INFO")
	}

	_, ok := pLocalUser.(types.AuthenticatedUser)

	if !ok {
		return errors.NewUnauthorizedError("Get User,USER Items cast error")
	}

	var params types.ViewUserReq
	err := c.ParamsParser(&params)

	if err != nil {
		return c.Status(400).JSON(err)
	}

	user, err := r.userRepository.GetById(params.Id, context.Background())

	if err != nil {
		return c.Status(400).JSON(err)
	}

	return c.Status(200).JSON(types.AuthenticatedUserFromUser(user))
}
