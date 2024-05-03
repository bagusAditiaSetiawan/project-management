package routes

import (
	"github.com/bagusAditiaSetiawan/project-management/api/handler"
	"github.com/bagusAditiaSetiawan/project-management/api/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"gorm.io/gorm"
)

func NewGroupRoute(app *fiber.App) fiber.Router {
	return app.Group("/api")
}

func NewRouteProject(app fiber.Router, db *gorm.DB, validation *validator.Validate) fiber.Router {
	projectController := handler.InitializeProjectController(db, validation)
	app.Post("/project/create", middleware.Protected(), projectController.Create)
	app.Post("/project/list", middleware.Protected(), projectController.Pagination)
	return app
}

func NewRouteHealth(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"Code":   fiber.StatusOK,
			"Status": "OK",
			"Data":   "Hello World",
		})
	})
	app.Get("/monitor", monitor.New(monitor.Config{
		Title:   "monitor service",
		APIOnly: true,
	}))
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"Code":   fiber.StatusOK,
			"Status": "OK",
			"Data":   "Serve Restapi",
		})
	})
}

func NewRouteAuth(app fiber.Router, db *gorm.DB, validation *validator.Validate) fiber.Router {
	authController := handler.InitializeAuthController(db, validation)
	app.Post("/register", authController.Register)
	app.Post("/login", authController.Login)
	app.Get("/profile", middleware.Protected(), authController.Profile)
	return app
}
