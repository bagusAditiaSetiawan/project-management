package aws_cloudwatch

type AwsCloudWatchService interface {
	SendLogInfo(a ...any) bool
}
