package rest

import (
	"automatic-doodle/pkg/validator"

	"automatic-doodle/types"

	"github.com/gofiber/fiber/v2"
)

func (r *Rest) CreateUploadUrl(c *fiber.Ctx) error {
	var payload types.CreateUploadUrlRequest
	if err := validator.ParseBody(c, &payload); err != nil {
		r.log.Warning(
			"Body is not valid for CreateUploadUrlRequest: %s",
			map[string]interface{}{
				"error": err,
			},
		)

		return err
	}
	response, err := r.fileService.CreateUploadUrl(&payload)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(response)
}
