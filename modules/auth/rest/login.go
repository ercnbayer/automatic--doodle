package rest

import (
	"automatic-doodle/pkg/errors"
	"automatic-doodle/types"

	"github.com/gofiber/fiber/v2"
)

func (r *Rest) loginUser(c *fiber.Ctx) error {
	var payload types.LoginRequest

	err := c.BodyParser(&payload)

	if err != nil {
		errors.New("User", "Invalid params")
		return c.Status(400).JSON(err)
	}

	response, err := r.authenticationService.
		Login(&payload)

	if err != nil {
		return errors.New("USER LOGIN", "SERVICE ERROR")
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    response.Token,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
		Path:     "/",
	})

	return c.Status(200).JSON(response)

}
