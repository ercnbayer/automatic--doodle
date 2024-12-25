package types

import (
	"automatic-doodle/ent"

	"time"
)

type JobAdvRequest struct {
	StartDate   time.Time `json:"start_date" validate:"required"`
	EndDate     time.Time `json:"end_date" validate:"required,gtefield=StartDate"`
	Fee         int       `json:"fee" validate:"required,min=1"`
	JobType     string    `json:"job_type" validate:"required,max=255"`
	Description string    `json:"description" validate:"required,max=1024"`
}

func JobAdvReponseFromJob(j *ent.Job, u UserPublicDetails) JobAdvResponse {

	return JobAdvResponse{Publisher: u, StartDate: j.StartDate, EndDate: j.EndDate, Fee: j.Fee, JobType: j.JobType, Description: j.Description}
}

type JobAdvResponse struct {
	Publisher   UserPublicDetails `json:"publisher"`
	StartDate   time.Time         `json:"start_date" validate:"required"`
	EndDate     time.Time         `json:"end_date" validate:"required,gtefield=StartDate"`
	Fee         int               `json:"fee" validate:"required,min=1"`
	JobType     string            `json:"job_type" validate:"required,max=255"`
	Description string            `json:"description" validate:"required,max=1024"`
}

type JobQuery struct {
	Identifier string `query:"params"`
	PageNumber int    `query:"pageNumber" validate:"gte=1"`
}
