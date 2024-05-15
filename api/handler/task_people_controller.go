package handler

import "github.com/gofiber/fiber/v2"

type TaskPeopleController interface {
	TaskPeopleCreate(ctx *fiber.Ctx) error
}
