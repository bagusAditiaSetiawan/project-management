package routes

import (
	"github.com/bagusAditiaSetiawan/project-management/api/handler"
	"github.com/bagusAditiaSetiawan/project-management/api/middleware"
	"github.com/bagusAditiaSetiawan/project-management/pkg/aws_cloudwatch"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewTaskRoute(app fiber.Router, db *gorm.DB, validation *validator.Validate, logger *aws_cloudwatch.AwsCloudWatchServiceImpl) fiber.Router {
	taskController := handler.InitializeTaskController(db, validation, logger)
	app.Post("/task/list", middleware.Protected(), taskController.TaskPagination)
	app.Get("/task/detail/:id", middleware.Protected(), taskController.TaskDetail)
	app.Post("/task/create", middleware.Protected(), taskController.TaskCreate)
	return app
}
