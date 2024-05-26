package routes

import (
	"github.com/bagusAditiaSetiawan/project-management/pkg/aws_cloudwatch"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewSetupRoute(routeApi fiber.Router, db *gorm.DB, validation *validator.Validate, loggerService *aws_cloudwatch.AwsCloudWatchServiceImpl) {
	NewRouteAuth(routeApi, db, validation, loggerService)
	NewRouteProject(routeApi, db, validation, loggerService)
	NewTaskRoute(routeApi, db, validation, loggerService)
	NewTaskPeopleRoute(routeApi, db, validation, loggerService)
	NewTaskPeopleRoute(routeApi, db, validation, loggerService)
	NewFileRoute(routeApi, db, validation, loggerService)
}
