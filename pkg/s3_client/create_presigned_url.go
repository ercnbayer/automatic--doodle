package s3client

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (svc *Service) CreateUploadPresignedUrl(
	bucketName,
	objectKey string,
	expiry time.Duration,
) (string, error) {
	putObjectInput := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}

	req, err := svc.presigner.PresignPutObject(
		context.TODO(),
		putObjectInput,
		s3.WithPresignExpires(expiry),
	)
	if err != nil {
		svc.log.Error(`Failed to sign request: `, err)
		return "", err
	}

	return req.URL, nil
}
