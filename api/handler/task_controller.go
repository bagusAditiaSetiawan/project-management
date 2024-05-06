package handler

import "github.com/gofiber/fiber/v2"

type TaskController interface {
	TaskPagination(ctx *fiber.Ctx) error
	TaskCreate(ctx *fiber.Ctx) error
}
