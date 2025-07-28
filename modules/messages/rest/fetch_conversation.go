package rest

import (
	"automatic-doodle/pkg/errors"
	"automatic-doodle/types"
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (r *Rest) FetchConversationsByUserId(c *fiber.Ctx) error {

	id := c.Get("id")

	fmt.Println("api istek")

	uuid, err := uuid.Parse(id)

	if err != nil {
		fmt.Println(err)
		return c.Status(400).JSON(err)
	}

	pUser := c.Locals("pUser")

	userItem, ok := pUser.(types.AuthenticatedUser)

	if !ok {
		return c.Status(500).JSON(errors.New("Fetch Conversation", "Unknown Object"))
	}

	conversation, err := r.messageService.
		FetchConversationsByUserId(
			userItem.Id,
			uuid,
			context.Background(),
		)

	if err != nil {
		return c.Status(400).JSON(err)
	}

	fmt.Println("bak oldu iste")
	return c.Status(200).JSON(conversation)
}
