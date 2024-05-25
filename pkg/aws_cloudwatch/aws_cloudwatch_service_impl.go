package aws_cloudwatch

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/bagusAditiaSetiawan/project-management/api/config"
	"github.com/bagusAditiaSetiawan/project-management/api/helpers"
	"time"
)

type AwsCloudWatchServiceImpl struct {
	Logging *cloudwatchlogs.CloudWatchLogs
}

func NewCloudWatchLogsService(sess *session.Session) *cloudwatchlogs.CloudWatchLogs {
	return cloudwatchlogs.New(sess)
}

func NewAwsCloudWatchServiceImpl(logging *cloudwatchlogs.CloudWatchLogs) *AwsCloudWatchServiceImpl {
	return &AwsCloudWatchServiceImpl{
		Logging: logging,
	}
}

func (service *AwsCloudWatchServiceImpl) SendLogInfo(a ...any) bool {
	group := config.Config("AWS_CLOUDWATCH_GROUP")
	stream := config.Config("AWS_CLOUDWATCH_STREAM")
	var message string

	for _, item := range a {
		message += fmt.Sprint(item) + " "
	}

	_, err := service.Logging.PutLogEvents(&cloudwatchlogs.PutLogEventsInput{
		LogGroupName:  aws.String(group),
		LogStreamName: aws.String(stream),
		LogEvents: []*cloudwatchlogs.InputLogEvent{
			{
				Message:   aws.String(message),
				Timestamp: aws.Int64(time.Now().UnixNano() / int64(time.Millisecond)),
			},
		},
	})
	helpers.IfPanicHelper(err)
	return true
}
