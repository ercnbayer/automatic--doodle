package rest

import (
	"automatic-doodle/pkg/errors"
	"automatic-doodle/pkg/validator"
	"automatic-doodle/types"
	"context"

	"github.com/gofiber/fiber/v2"
)

func (r *Rest) CreateJobAdv(c *fiber.Ctx) error {
	var payload types.JobAdvRequest

	err := c.BodyParser(&payload)

	if err != nil {
		return err
	}
	err = validator.ParseBody(c, &payload)

	if err == nil {
		return err
	}

	pUser := c.Locals("user")
	if pUser == nil {
		return errors.New("middleware err", "Unauthroized user err")
	}
	userItem, ok := pUser.(types.AuthenticatedUser)

	if !ok {
		r.logger.Info("Local user is not user entity")
		return errors.NewUnauthorizedError("MiddlewaresModule")
	}

	pJob := r.jobFactory.Create(payload.Fee, payload.Description, payload.JobType, payload.StartDate, payload.EndDate, userItem.Id)

	jobItem, err := r.jobRepository.Create(pJob, context.Background())

	if err != nil {
		return c.Status(400).JSON(err)
	}

	return c.Status(200).JSON(types.JobAdvReponseFromJob(jobItem, types.UserPublicDetailsFromAUser(userItem)))

}
