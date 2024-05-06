package routes

import (
	"github.com/gofiber/fiber/v2"
)

func NewGroupRoute(app *fiber.App) fiber.Router {
	return app.Group("/api")
}
