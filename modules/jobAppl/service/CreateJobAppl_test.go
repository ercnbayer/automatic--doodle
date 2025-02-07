package service

import (
	"automatic-doodle/ent"
	"automatic-doodle/types"
	"context"
	"testing"

	"github.com/google/uuid"
)

type JobApplRepoMock struct{}

func (*JobApplRepoMock) Create(item *ent.JobApplicationCreate, ctx context.Context) (*ent.JobApplication, error) {
	return &ent.JobApplication{}, nil
}
func (*JobApplRepoMock) GetJobApplsWithJobId(offset int, limit int, jobID uuid.UUID, ctx context.Context) ([]*ent.JobApplication, error) {

	return []*ent.JobApplication{}, nil
}

type JobApplFacMock struct{}

func (*JobApplFacMock) Create(description string, pFileKey *string, jobID uuid.UUID, userID uuid.UUID) *ent.JobApplicationCreate {
	return &ent.JobApplicationCreate{}
}

type JobRepoMock struct{}

func (*JobRepoMock) GetJobById(uuid.UUID, context.Context) (*ent.Job, error) {
	return &ent.Job{}, nil
}

type UserRepoMock struct{}

func (*UserRepoMock) GetById(id uuid.UUID, ctx context.Context) (*ent.User, error) {
	return &ent.User{}, nil
}

func TestCreateJobAppl(t *testing.T) {

	srv := Service{
		jobApplFactory:    &JobApplFacMock{},
		jobApplRepository: &JobApplRepoMock{},
		jobRepository:     &JobRepoMock{},
		userRepository:    &UserRepoMock{},
	}

	res, err := srv.CreateJobAppl(types.CreateJobApplRequest{})

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	t.Log(res)
}
