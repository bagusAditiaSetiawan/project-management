package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewSetupRoute(routeApi fiber.Router, db *gorm.DB, validation *validator.Validate) {
	NewRouteAuth(routeApi, db, validation)
	NewRouteProject(routeApi, db, validation)
	NewTaskRoute(routeApi, db, validation)
	NewTaskPeopleRoute(routeApi, db, validation)
}
