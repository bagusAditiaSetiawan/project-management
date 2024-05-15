package routes

import (
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"net/http"
)

func NewRouteHealth(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(&presenter.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   "Status OK",
		})
	})
	app.Get("/monitor", monitor.New(monitor.Config{
		Title:   "monitor service",
		APIOnly: true,
	}))
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.JSON(&presenter.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   "Serve Api",
		})
	})
}
