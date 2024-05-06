package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

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
