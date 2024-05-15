package aws_s3

import (
	"flag"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/bagusAditiaSetiawan/project-management/api/config"
	"github.com/bagusAditiaSetiawan/project-management/api/helpers"
	"mime/multipart"
)

type AwsS3ServiceImpl struct {
	S3Service *s3.S3
}

func NewS3Service(sess *session.Session) *s3.S3 {
	return s3.New(sess)
}

func NewAwsS3ServiceImpl(s3 *s3.S3) *AwsS3ServiceImpl {
	return &AwsS3ServiceImpl{
		S3Service: s3,
	}
}

func (service *AwsS3ServiceImpl) UploadS3(file *multipart.FileHeader) string {
	bucket := flag.String("bucket", "", config.Config("AWS_S3_BUCKET"))
	key := flag.String("key", "", "public/"+file.Filename)
	openFile, err := file.Open()
	helpers.IfPanicHelper(err)

	response, err := service.S3Service.PutObject(&s3.PutObjectInput{
		Bucket: bucket,
		Key:    key,
		Body:   openFile,
	})
	helpers.IfPanicHelper(err)
	return response.String()
}
