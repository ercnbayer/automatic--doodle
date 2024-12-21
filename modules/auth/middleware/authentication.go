package middleware

import (
	"automatic-doodle/pkg/errors"

	"github.com/gofiber/fiber/v2"
)

func (mw *Middleware) Auth(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	tokenHeaders := headers["X-Auth"]

	if len(tokenHeaders) == 0 {

		// to do throw err
		errors.NewUnauthorizedError("auth err token headers")

	}

	token := tokenHeaders[0]

	if len(token) == 0 {

		// to do throw err
		errors.NewUnauthorizedError("auth err getting first token ")
	}

	pUser, err := mw.authenticationService.GetUserByToken(&token)

	if err != nil {
		// to do throw err
		errors.NewUnauthorizedError("auth err get token headers")
	}
	c.Locals("user", pUser)

	return c.Next()

}
