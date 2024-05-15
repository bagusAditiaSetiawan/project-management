package routes

import (
	"github.com/bagusAditiaSetiawan/project-management/api/handler"
	"github.com/bagusAditiaSetiawan/project-management/api/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewRouteAuth(app fiber.Router, db *gorm.DB, validation *validator.Validate) fiber.Router {
	authController := handler.InitializeAuthController(db, validation)
	app.Post("/register", authController.Register)
	app.Post("/login", authController.Login)
	app.Get("/profile", middleware.Protected(), authController.Profile)
	return app
}
