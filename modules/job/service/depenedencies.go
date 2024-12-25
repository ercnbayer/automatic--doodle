package service

import (
	"automatic-doodle/ent"
	"context"

	"github.com/google/uuid"
)

type JobRepository interface {
	GetByIdentifier(
		identifier string,
		ctx context.Context,
		offset int, limit int,
	) ([]*ent.Job, error)
}

type UserRepository interface {
	GetById(id uuid.UUID, ctx context.Context) (*ent.User, error)
}
