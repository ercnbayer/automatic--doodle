package service

import (
	"automatic-doodle/ent"
	"context"

	"github.com/google/uuid"
)

type JobApplRepository interface {
	Create(item *ent.JobApplicationCreate, ctx context.Context) (*ent.JobApplication, error)
	GetJobApplsWithJobId(offset int, limit int, jobID uuid.UUID, ctx context.Context) ([]*ent.JobApplication, error)
}

type JobApplFactory interface {
	Create(description string, pFileKey *string, jobID uuid.UUID, userID uuid.UUID) *ent.JobApplicationCreate
}

type JobRepository interface {
	GetJobById(uuid.UUID, context.Context) (*ent.Job, error)
}

type UserRepository interface {
	GetById(id uuid.UUID, ctx context.Context) (*ent.User, error)
}
