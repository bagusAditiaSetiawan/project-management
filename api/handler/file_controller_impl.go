package handler

import (
	"github.com/bagusAditiaSetiawan/project-management/api/helpers"
	"github.com/bagusAditiaSetiawan/project-management/pkg/aws_s3"
	"github.com/gofiber/fiber/v2"
)

type FileControllerImpl struct {
	AwsS3Service aws_s3.S3Service
}

func NewFileControllerImpl(awsS3Service aws_s3.S3Service) *FileControllerImpl {
	return &FileControllerImpl{
		AwsS3Service: awsS3Service,
	}
}

func (controller *FileControllerImpl) uploadFile(ctx *fiber.Ctx) error {
	fileHeader, err := ctx.FormFile("file")
	helpers.IfPanicHelper(err)

	controller.AwsS3Service.UploadS3(fileHeader)
	panic("implement me")
}
