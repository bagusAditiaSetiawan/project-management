package routes

import (
	"github.com/bagusAditiaSetiawan/project-management/api/handler"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewFileRoute(app fiber.Router, db *gorm.DB, validation *validator.Validate) fiber.Router {
	fileController := handler.InitializeFileController(db, validation)
	app.Post("/upload", fileController.UploadFile)
	return app
}
