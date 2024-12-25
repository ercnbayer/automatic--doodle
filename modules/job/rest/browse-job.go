package rest

import (
	"automatic-doodle/pkg/errors"
	"automatic-doodle/types"
	"context"
	"math"

	"github.com/gofiber/fiber/v2"
)

func (r *Rest) BrowseJobs(c *fiber.Ctx) error {

	var query types.JobQuery
	err := c.QueryParser(&query)

	if err != nil {
		return errors.New("Browse Jobs", err.Error())
	}

	pageSize := 20 // maybe it can come from some env some variable etc i dont know

	offset := query.PageNumber * (pageSize - 1)

	if offset > math.MaxInt {
		return errors.New("Browse Jobs", "PARAM ERR")
	}

	jobs, err := r.jobService.BrowseJob(offset, pageSize, query.Identifier, context.Background())

	if err != nil {
		return errors.New("Browse Jobs", err.Error())
	}

	return c.Status(200).JSON(jobs)
}
