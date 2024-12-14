package s3client

import (
	"fmt"

	"automatic-doodle/types"

	"automatic-doodle/pkg/env"
)

func (svc *Service) GetObjectAccesUrl(
	bucketName,
	objectKey string,
) string {
	if env.GO_ENV == types.GoEnvProduction {
		return fmt.Sprintf(
			"https://%s.s3.%s.amazonaws.com/%s",
			bucketName,
			svc.region,
			objectKey,
		)
	} else {
		return fmt.Sprintf("%s/%s/%s", svc.endpoint, bucketName, objectKey)
	}
}
