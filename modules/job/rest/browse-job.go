package rest

import (
	"automatic-doodle/pkg/errors"
	"automatic-doodle/types"
	"context"

	"github.com/gofiber/fiber/v2"
)

func (r *Rest) BrowseJobs(c *fiber.Ctx) error {

	var query types.JobQuery
	err := c.QueryParser(&query)

	if err != nil {
		return errors.New("Browse Jobs", err.Error())
	}

	jobs, err := r.jobRepository.GetByIdentifier(query.Identifier, context.Background())

	if err != nil {
		return errors.New("Browse Jobs", "GetByIdentifer err")
	}

	return c.Status(200).JSON(jobs)
}
