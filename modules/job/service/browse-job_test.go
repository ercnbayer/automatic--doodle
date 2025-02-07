package service

import (
	"automatic-doodle/ent"
	"context"
	"testing"

	"github.com/google/uuid"
)

type UserRepoMock struct {
}

func (*UserRepoMock) GetById(id uuid.UUID, ctx context.Context) (*ent.User, error) {
	return &ent.User{}, nil
}

type JobRepoMock struct{}

func (*JobRepoMock) GetByIdentifier(
	identifier string,
	ctx context.Context,
	offset int, limit int,
) ([]*ent.Job, error) {
	return []*ent.Job{}, nil
}

func TestBrowseJobSuccess(t *testing.T) {

	srv := Service{
		userRepository: &UserRepoMock{},
		jobRepository:  &JobRepoMock{},
	}

	res, err := srv.BrowseJob(1, 1, "", context.Background())

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	t.Log(res)

}
