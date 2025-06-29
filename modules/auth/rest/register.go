package rest

import (
	"automatic-doodle/pkg/validator"
	"automatic-doodle/types"

	"github.com/gofiber/fiber/v2"
)

func (r *Rest) Register(c *fiber.Ctx) error {
	var payload types.RegisterRequest

	err := validator.ParseBody(c, &payload)

	if err != nil {

		r.log.Warning(
			"Body is not valid for Register: %s",
			map[string]interface{}{
				"error": err,
			},
		)
		return c.Status(400).JSON(err)

	}

	tokens, RegisterErr := r.authenticationService.Register(&payload)

	if RegisterErr != nil {
		r.log.Info("service err")
		return c.Status(400).JSON("Bad Req")
	}
	r.log.Info("Success")

	return c.Status(200).JSON(tokens)
}
