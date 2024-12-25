package service

import (
	"automatic-doodle/pkg/errors"
	"automatic-doodle/types"
	"context"
)

func (srv *Service) BrowseJob(offset int, limit int, identifier string, ctx context.Context) ([]types.JobAdvResponse, error) {

	jobs, err := srv.jobRepository.GetByIdentifier(identifier, ctx, offset, limit)

	if err != nil {
		return nil, errors.New("Browse Job Service Err", err.Error())
	}

	var response []types.JobAdvResponse

	for _, j := range jobs {

		singleRes := types.JobAdvReponseFromJob(j, types.UserPublicDetailsFromUser(j.Edges.User))

		response = append(response, singleRes)
	}

	return response, nil
}
