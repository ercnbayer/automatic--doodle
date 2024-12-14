package service

import (
	"automatic-doodle/ent"
	"automatic-doodle/ent/file"
	"context"
	"time"

	"github.com/google/uuid"
)

type Logger interface {
	Trace(format string, args ...any)
	Info(format string, args ...any)
	Warning(format string, args ...any)
	Error(format string, args ...any)
	Fatal(format string, args ...any)
}

type FileFactory interface {
	Create(
		fileName,
		extention,
		contentType,
		bucket string,
		createdById uuid.UUID,
		fileType file.Type,
	) *ent.FileCreate
}

type FileRepository interface {
	Create(
		*ent.FileCreate,
		context.Context,
	) (*ent.File, error)
}

type S3Client interface {
	CreateUploadPresignedUrl(
		bucketName,
		objectKey string,
		expiry time.Duration,
	) (string, error)

	GetObjectAccesUrl(
		bucketName,
		objectKey string,
	) string
}
