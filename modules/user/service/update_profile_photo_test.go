package service

import (
	"automatic-doodle/ent"
	"automatic-doodle/ent/file"
	"automatic-doodle/ent/user"
	"automatic-doodle/types"
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
)

type LoggerMock struct {
}

func (*LoggerMock) Trace(format string, args ...any) {
	fmt.Sprintf(format, args...)
}
func (*LoggerMock) Fatal(format string, args ...any) {
	fmt.Sprintf(format, args...)
}
func (*LoggerMock) Warning(format string, args ...any) {
	fmt.Sprintf(format, args...)
}
func (*LoggerMock) Info(format string, args ...any) {
	fmt.Sprintf(format, args...)
}
func (*LoggerMock) Error(format string, args ...any) {
	fmt.Sprintf(format, args...)
}

type FileFacMock struct{}

func (*FileFacMock) Create(string,
	string,
	string, string,
	uuid.UUID,
	file.Type) *ent.FileCreate {
	return &ent.FileCreate{}
}

type FileRepMock struct{}

func (*FileRepMock) Create(*ent.FileCreate, context.Context) (*ent.File, error) {

	return &ent.File{}, nil
}

type UserRepMock struct{}

func (*UserRepMock) Create(*ent.UserCreate, context.Context) (*ent.User, error) {

	return &ent.User{}, nil
}

func (*UserRepMock) DeleteUser(uuid.UUID, context.Context) error {
	return nil
}

func (*UserRepMock) GetById(uuid.UUID, context.Context) (*ent.User, error) {

	return &ent.User{}, nil
}

func (*UserRepMock) UpdateUser(ctx context.Context, u types.UpdateUserReq, id uuid.UUID) (*ent.User, error) {
	return &ent.User{}, nil
}

type UserFacMock struct{}

func (*UserFacMock) Create(string, string, string, string, string, user.Role, user.State) *ent.UserCreate {
	return &ent.UserCreate{}
}

func TestProfilePhotoSuccess(t *testing.T) {

	srv := UserService{
		userRepository: &UserRepMock{},
		fileFactory:    &FileFacMock{},
		fileRepository: &FileRepMock{},
		userFactory:    &UserFacMock{},
		logger:         &LoggerMock{},
	}

	res, err := srv.UpdateProfilePhoto(uuid.New(), &types.UpdateProfilePhotoRequest{})

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	t.Log(res)
}
