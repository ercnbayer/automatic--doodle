package service

import (
	"automatic-doodle/ent"
	"automatic-doodle/pkg/errors"
	"automatic-doodle/types"
	"context"
)

func (srv *Service) CreateJobAppl(request types.CreateJobApplRequest) (types.CreateJobApplResponse, error) {

	var item *ent.JobApplicationCreate

	ctx := context.Background()

	_, err := srv.jobRepository.GetJobById(request.JobID, ctx)
	if err != nil {
		return types.CreateJobApplResponse{}, errors.New("Create Job Appl", "Not a valid job")
	}

	if request.HasAttachment {

		item = srv.jobApplFactory.Create(request.Description, &request.FileKey, request.JobID, request.ApplicantId)

	} else {
		item = srv.jobApplFactory.Create(request.Description, nil, request.JobID, request.ApplicantId)
	}

	jobAppl, err := srv.jobApplRepository.Create(item, ctx)

	if err != nil {
		return types.CreateJobApplResponse{}, errors.New("Create Job Appl", "Create Err")
	}

	response := types.ReturnJobApplResponse("post_files", request.FileKey, request.Description, request.ApplicantId, jobAppl.JobID)

	return response, nil

}
