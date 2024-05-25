package aws_s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/bagusAditiaSetiawan/project-management/api/config"
	"github.com/bagusAditiaSetiawan/project-management/api/helpers"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
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
	openFile, err := file.Open()
	helpers.IfPanicHelper(err)
	filename := strconv.Itoa(int(time.Now().Unix())) + strings.ToLower(file.Filename)
	_, err = service.S3Service.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(config.Config("AWS_S3_BUCKET")),
		Key:    aws.String("public/" + filename),
		Body:   openFile,
	})
	helpers.IfPanicHelper(err)
	return filename
}

func (service *AwsS3ServiceImpl) createUrl(fileName string) string {
	req, _ := service.S3Service.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(config.Config("AWS_S3_BUCKET")),
		Key:    aws.String(fileName),
	})
	url, err := req.Presign(15 * time.Minute)
	helpers.IfPanicHelper(err)
	return url
}
