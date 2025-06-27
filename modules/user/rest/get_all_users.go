package rest

import (
	
	"automatic-doodle/types"
	"context"

	"github.com/gofiber/fiber/v2"
)

func (r *Rest) GetAllUsers(c *fiber.Ctx) error {
	users, err := r.userRepository.GetAllUsers(context.Background())

	if err != nil {
		return c.Status(400).JSON(err)
	}

	// Convert to AuthenticatedUser format for response
	var responseUsers []types.AuthenticatedUser
	for _, user := range users {
		responseUsers = append(responseUsers, types.AuthenticatedUserFromUser(user))
	}

	return c.Status(200).JSON(responseUsers)
} 