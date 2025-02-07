package service

import (
	"automatic-doodle/ent"
	"automatic-doodle/ent/file"
	"automatic-doodle/types"
	"context"
	"fmt"
	"testing"
	"time"

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

func (*FileFacMock) Create(
	fileName,
	extention,
	contentType,
	bucket string,
	createdById uuid.UUID,
	fileType file.Type,
) *ent.FileCreate {
	return &ent.FileCreate{}
}

type FileRepMock struct{}

func (*FileRepMock) Create(
	*ent.FileCreate,
	context.Context,
) (*ent.File, error) {
	return &ent.File{}, nil
}

type S3ClientMock struct{}

func (*S3ClientMock) CreateUploadPresignedUrl(
	bucketName,
	objectKey string,
	expiry time.Duration,
) (string, error) {
	return "someUrl", nil
}

func (*S3ClientMock) GetObjectAccesUrl(
	bucketName,
	objectKey string,
) string {
	return "someUrl"
}

func TestCreateFileUpload(t *testing.T) {
	srv := Service{
		logger:     &LoggerMock{},
		s3Client:   &S3ClientMock{},
		factory:    &FileFacMock{},
		repository: &FileRepMock{},
	}

	res, err := srv.CreateUploadUrl(&types.CreateUploadUrlRequest{})

	if err != nil {

		t.Log(err)

		t.Fail()
	}
	t.Log(res)
}
