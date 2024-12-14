package s3client

import (
	"context"
	"sync"

	"automatic-doodle/types"

	"automatic-doodle/pkg/env"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	module     Service
	moduleOnce sync.Once
)

func New(
	log Logger,
	cfg ConfigModule,
) *Service {
	moduleOnce.Do(func() {
		endpoint, err := cfg.GetConfig("AWS_ENDPOINT")
		if err != nil {
			log.Fatal(`Missing env value "AWS_ENDPOINT": `, err)
			return
		}

		region, err := cfg.GetConfig("AWS_REGION")
		if err != nil {
			log.Fatal(`Missing env value "AWS_REGION": `, err)
			return
		}

		accessKey, err := cfg.GetConfig("AWS_ACCESS_KEY")
		if err != nil {
			log.Fatal(`Missing env value "AWS_ACCESS_KEY": `, err)
			return
		}

		secretKey, err := cfg.GetConfig("AWS_SECRET_KEY")
		if err != nil {
			log.Fatal(`Missing env value "AWS_SECRET_KEY": `, err)
			return
		}

		var cfg aws.Config
		creds := credentials.NewStaticCredentialsProvider(
			accessKey,
			secretKey,
			"",
		)
		if env.GO_ENV == types.GoEnvProduction {
			cfg, err = config.LoadDefaultConfig(
				context.TODO(),
				config.WithRegion(region),
				config.WithCredentialsProvider(creds),
				config.WithBaseEndpoint(endpoint),
			)
			if err != nil {
				log.Fatal(`AWS config initialize failed with: `, err)
				return
			}
		}

		s3Service := s3.NewFromConfig(cfg)
		presigner := s3.NewPresignClient(s3Service)
		module = Service{
			log:       log,
			s3Service: s3Service,
			presigner: presigner,
			region:    region,
			endpoint:  endpoint,
			accessKey: accessKey,
			secretKey: secretKey,
		}
	})

	return &module
}

type Service struct {
	s3Service *s3.Client
	presigner *s3.PresignClient
	log       Logger
	region    string
	endpoint  string
	accessKey string
	secretKey string
}
