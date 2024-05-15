package aws

import (
	"flag"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/bagusAditiaSetiawan/project-management/api/config"
)

func NewAwsSessionService() *session.Session {
	region := config.Config("AWS_S3_REGION")
	sess := session.Must(session.NewSession(&aws.Config{
		Region: flag.String("region", "", region),
	}))
	return sess
}
