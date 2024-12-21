package rest

import (
	"automatic-doodle/pkg/errors"
	"automatic-doodle/types"

	"github.com/gofiber/fiber/v2"
)

func (r *Rest) CreateJobAppl(c *fiber.Ctx) error {
	var payload types.CreateJobApplRequest

	err := c.BodyParser(&payload)

	if err != nil {

		return errors.New("Cretae Job Appl EndPoint", "INVALID PAYLOAD")
	}
	pUser := c.Locals("user")
	userItem, ok := pUser.(types.AuthenticatedUser)

	if !ok {
		r.logger.Info("Local user is not user entity")
		return errors.NewUnauthorizedError("MiddlewaresModule")
	}
	if payload.ApplicantId != userItem.Id {
		// payload sender trying to mess with api

		return errors.New("Create Job api", "Invalid ID")
	}

	response, err := r.jobApplService.CreateJobAppl(payload)

	if err != nil {

		return errors.New("job appll Service", "Service Error")
	}

	return c.Status(200).JSON(response)
}
