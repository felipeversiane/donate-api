package aws

import (
	"log/slog"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/felipeversiane/donate-api/internal/infra/config"
)

type cloudService struct {
	session *session.Session
}

type CloudServiceInterface interface {
	GetSession() *session.Session
}

func NewCloudService(config config.CloudServiceConfig) CloudServiceInterface {
	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String(config.Region),
		Credentials:      credentials.NewStaticCredentials(config.AccessKey, config.SecretAccessKey, ""),
		Endpoint:         aws.String(config.Endpoint),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		slog.Error("failed to create a new aws session", slog.Any("error", err))
	}
	return &cloudService{sess}
}

func (c *cloudService) GetSession() *session.Session {
	return c.session
}
