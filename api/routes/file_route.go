package routes

import (
	"github.com/bagusAditiaSetiawan/project-management/api/handler"
	"github.com/bagusAditiaSetiawan/project-management/pkg/aws_cloudwatch"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewFileRoute(app fiber.Router, db *gorm.DB, validation *validator.Validate, logger *aws_cloudwatch.AwsCloudWatchServiceImpl) fiber.Router {
	fileController := handler.InitializeFileController(db, validation, logger)
	app.Post("/upload", fileController.UploadFile)
	return app
}
