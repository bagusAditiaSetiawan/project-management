package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/bagusAditiaSetiawan/project-management/api/config"
)

func NewAwsSessionService() *session.Session {
	region := aws.String(config.Config("AWS_S3_REGION"))
	secret := config.Config("AWS_S3_SECRET")
	key := config.Config("AWS_S3_KEY")
	cred := credentials.NewStaticCredentials(key, secret, "")
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      region,
		Credentials: cred,
	}))
	return sess
}
