package aws_s3

import "mime/multipart"

type S3Service interface {
	UploadS3(file *multipart.FileHeader) string
	createUrl(fileName string) string
}
