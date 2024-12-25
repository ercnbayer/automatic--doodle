package service

import (
	"automatic-doodle/pkg/errors"
	"automatic-doodle/types"
	"context"

	"github.com/google/uuid"
)

func (srv *Service) CheckJobOwner(offset int, limit int, id uuid.UUID, userID uuid.UUID, ctx context.Context) ([]types.GetJobApplResponse, error) {

	job, err := srv.jobRepository.GetJobById(id, ctx)

	if err != nil {
		return nil, err
	}

	if job.JobOwner != userID {

		return nil, errors.New("JobApplService", "jobID ")
	}

	jobAppls, err := srv.jobApplRepository.GetJobApplsWithJobId(offset, limit, job.ID, ctx)

	var GetJobApplResponse []types.GetJobApplResponse

	for _, ja := range jobAppls {

		jobAppl := types.GetJobApplResponse{
			User:        types.AuthenticatedUserFromUser(ja.Edges.User),
			Description: ja.Description,
			FileKey:     ja.ObjectKey,
			BucketName:  "POST_FILES", // TO DO PROBABLY HAVE TO FIX IT CAN COME FROM ENV

		}
		GetJobApplResponse = append(GetJobApplResponse, jobAppl)

	}
	return GetJobApplResponse, nil
}
