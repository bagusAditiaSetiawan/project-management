package routes

import (
	"github.com/bagusAditiaSetiawan/project-management/api/handler"
	"github.com/bagusAditiaSetiawan/project-management/api/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewTaskPeopleRoute(app fiber.Router, db *gorm.DB, validation *validator.Validate) fiber.Router {
	taskPeopleController := handler.InitializeTaskPeopleController(db, validation)
	app.Post("/task-people/create", middleware.Protected(), taskPeopleController.TaskPeopleCreate)
	return app
}
