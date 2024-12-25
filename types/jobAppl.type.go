package types

import (
	"github.com/google/uuid"
)

type CreateJobApplRequest struct {
	ApplicantId   uuid.UUID `json:"applicantID"`
	JobID         uuid.UUID `json:"jobID"`
	Description   string    `json:"description"`
	HasAttachment bool      `json:"has_attachment"`
	FileKey       string
}

type CreateJobApplResponse struct {
	ApplicantId uuid.UUID `json:"applicantID"`
	JobID       uuid.UUID `json:"jobID"`
	Description string    `json:"description"`
	FileKey     string
	BucketName  string
}

type GetJobApplResponse struct {
	User        AuthenticatedUser
	FileKey     string
	BucketName  string
	Description string
}

type GetJobApplQuery struct {
	JobID      uuid.UUID `query:"id" validate:"required,uuid"`
	PageNumber int       `query:"pageNumber" validate:"gte=1"`
}

func ReturnJobApplResponse(bucketName string, FileKey string, description string, id uuid.UUID, jobID uuid.UUID) CreateJobApplResponse {
	return CreateJobApplResponse{
		BucketName:  bucketName,
		JobID:       jobID,
		Description: description,
		FileKey:     FileKey,
	}
}
