package rest

import (
	"automatic-doodle/pkg/errors"
	"automatic-doodle/types"
	"context"
	"math"

	"github.com/gofiber/fiber/v2"
)

func (r *Rest) GetJobAppl(c *fiber.Ctx) error {

	pUser := c.Locals("user")

	if pUser == nil {
		return errors.New("middleware err", "Unauthroized user err")
	}

	userItem, ok := pUser.(types.AuthenticatedUser)

	if !ok {
		return errors.New("GET JOB ADV", "Token err")
	}

	var query types.GetJobApplQuery
	err := c.QueryParser(&query)

	if err != nil {
		return errors.New("Rest Get Job Appl", err.Error())
	}

	pageSize := 20 // maybe it can come from some env some variable etc i dont know

	offset := query.PageNumber * (pageSize - 1)

	if offset > math.MaxInt {
		return errors.New("Rest Get Job", "PARAM ERR")
	}

	jobAppls, err := r.jobApplService.CheckJobOwner(offset, pageSize, query.JobID, userItem.Id, context.Background())

	if err != nil {
		return errors.New(" REST GET JOB ADV ", "Not a job owner")

	}

	return c.Status(200).JSON(jobAppls)

}
