package service

import (
	"automatic-doodle/types"
	"fmt"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

func (svc *Service) CreateUploadUrl(
	request *types.CreateUploadUrlRequest,
) (types.CreateUploadUrlResponse, error) {
	fileExt := filepath.Ext(request.Filename)
	bucket := svc.GetBucketName(request.Type)
	objectKey := fmt.Sprintf("%s.%s", uuid.New().String(), fileExt)
	url, err := svc.s3Client.CreateUploadPresignedUrl(
		bucket,
		objectKey,
		time.Duration(time.Minute*15),
	)
	if err != nil {
		return types.CreateUploadUrlResponse{}, err
	}

	svc.logger.Info("URL:%s\n KEY:%s\n", url, objectKey)
	return types.CreateUploadUrlResponse{
		Url: url,
		Key: objectKey,
	}, nil
}
