package main

import (
	"encoding/json"
	"github.com/bagusAditiaSetiawan/project-management/api/database"
	"github.com/bagusAditiaSetiawan/project-management/api/exception"
	"github.com/bagusAditiaSetiawan/project-management/api/routes"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewFiberApp() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandlerException,
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})

	app.Use(recover.New())
	return app
}

func NewValidator() *validator.Validate {
	return validator.New()
}

func NewServe() *fiber.App {
	app := NewFiberApp()
	validation := NewValidator()
	db := database.NewConnectDatabase()
	routes.NewRouteHealth(app)
	routeApi := routes.NewGroupRoute(app)
	routes.NewSetupRoute(routeApi, db, validation)
	return app
}

func main() {
	app := NewServe()
	app.Listen(":8080")
}
