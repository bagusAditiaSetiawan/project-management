package routes

import (
	"github.com/bagusAditiaSetiawan/project-management/pkg/aws"
	"github.com/bagusAditiaSetiawan/project-management/pkg/aws_cloudwatch"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewSetupRoute(routeApi fiber.Router, db *gorm.DB, validation *validator.Validate) {
	sess := aws.NewAwsSessionService()
	awsCloudWatch := aws_cloudwatch.NewCloudWatchLogsService(sess)
	loggerService := aws_cloudwatch.NewAwsCloudWatchServiceImpl(awsCloudWatch)

	NewRouteAuth(routeApi, db, validation, loggerService)
	NewRouteProject(routeApi, db, validation, loggerService)
	NewTaskRoute(routeApi, db, validation, loggerService)
	NewTaskPeopleRoute(routeApi, db, validation, loggerService)
	NewTaskPeopleRoute(routeApi, db, validation, loggerService)
	NewFileRoute(routeApi, db, validation, loggerService)
}
