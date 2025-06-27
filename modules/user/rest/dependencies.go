package rest

import (
	"automatic-doodle/ent"
	"automatic-doodle/types"
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AuthenticationMiddleware interface {
	Auth(c *fiber.Ctx) error
}

type UserRepository interface {
	//Create(*ent.UserCreate, context.Context) (*ent.User, error)
	DeleteUser(id uuid.UUID, ctx context.Context) error
	GetById(id uuid.UUID, ctx context.Context) (*ent.User, error)
	GetAllUsers(ctx context.Context) ([]*ent.User, error)
	UpdateUser(context.Context, types.UpdateUserReq, uuid.UUID) (*ent.User, error)
}
