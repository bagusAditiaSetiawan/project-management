package routes

import (
	"github.com/bagusAditiaSetiawan/project-management/api/handler"
	"github.com/bagusAditiaSetiawan/project-management/api/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewRouteProject(app fiber.Router, db *gorm.DB, validation *validator.Validate) fiber.Router {
	projectController := handler.InitializeProjectController(db, validation)
	app.Post("/project/create", middleware.Protected(), projectController.Create)
	app.Post("/project/list", middleware.Protected(), projectController.Pagination)
	app.Get("/project/:id", middleware.Protected(), projectController.Detail)
	return app
}
