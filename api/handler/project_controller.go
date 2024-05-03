package handler

import "github.com/gofiber/fiber/v2"

type ProjectController interface {
	Pagination(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
}
