package routes

import (
	"github.com/bagusAditiaSetiawan/project-management/api/handler"
	"github.com/bagusAditiaSetiawan/project-management/api/middleware"
	"github.com/bagusAditiaSetiawan/project-management/pkg/aws_cloudwatch"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewTaskPeopleRoute(app fiber.Router, db *gorm.DB, validation *validator.Validate, logger *aws_cloudwatch.AwsCloudWatchServiceImpl) fiber.Router {
	taskPeopleController := handler.InitializeTaskPeopleController(db, validation, logger)
	app.Post("/task-people/create", middleware.Protected(), taskPeopleController.TaskPeopleCreate)
	return app
}
