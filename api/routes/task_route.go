package routes

import (
	"github.com/bagusAditiaSetiawan/project-management/api/handler"
	"github.com/bagusAditiaSetiawan/project-management/api/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewTaskRoute(app fiber.Router, db *gorm.DB, validation *validator.Validate) fiber.Router {
	taskController := handler.InitializeTaskController(db, validation)
	app.Post("/task/list", middleware.Protected(), taskController.TaskPagination)
	app.Post("/task/create", middleware.Protected(), taskController.TaskCreate)
	return app
}
