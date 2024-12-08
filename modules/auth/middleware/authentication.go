package middleware

import "github.com/gofiber/fiber/v2"

func (mw *Middleware) Auth(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	tokenHeaders := headers["X-Auth"]

	if len(tokenHeaders) == 0 {

		// to do throw err

	}

	token := tokenHeaders[0]

	if len(token) == 0 {

		// to do throw err
	}

	pUser, err := mw.authenticationService.GetUserByToken(&token)

	if err != nil {
		// to do throw err
	}
	c.Locals("user", pUser)

	return c.Next()

}
