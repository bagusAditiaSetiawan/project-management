package routes

import (
	"github.com/bagusAditiaSetiawan/project-management/api/handler"
	"github.com/bagusAditiaSetiawan/project-management/api/middleware"
	"github.com/bagusAditiaSetiawan/project-management/pkg/aws_cloudwatch"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewRouteProject(app fiber.Router, db *gorm.DB, validation *validator.Validate, logger *aws_cloudwatch.AwsCloudWatchServiceImpl) fiber.Router {
	projectController := handler.InitializeProjectController(db, validation, logger)
	app.Post("/project/create", middleware.Protected(), projectController.Create)
	app.Post("/project/list", middleware.Protected(), projectController.Pagination)
	app.Get("/project/:id", middleware.Protected(), projectController.Detail)
	return app
}
