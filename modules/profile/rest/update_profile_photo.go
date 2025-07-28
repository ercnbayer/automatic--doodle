package rest

import (
	"automatic-doodle/pkg/errors"
	"automatic-doodle/pkg/validator"
	"automatic-doodle/types"

	"github.com/gofiber/fiber/v2"
)

func (r *Rest) UpdateProfilePhoto(c *fiber.Ctx) error {
	var payload types.UpdateProfilePhotoRequest

	if err := validator.ParseBody(c, &payload); err != nil {
		r.logger.Warning(
			"Body is not valid for UpdateProfilePhoto: %s",
			map[string]interface{}{
				"error": err,
			},
		)

		return err
	}

	pUser := c.Locals("user")

	userItem, ok := pUser.(types.AuthenticatedUser)
	if !ok {
		r.logger.Info("Local user is not user entity")
		return errors.NewUnauthorizedError("ProfileRest")
	}
	response, err := r.userService.UpdateProfilePhoto(userItem.Id, &payload)
	if err != nil {
		return err
	}

	return c.
		Status(200).
		JSON(response)
}
