package handler

import "github.com/gofiber/fiber/v2"

type AuthController interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	Profile(ctx *fiber.Ctx) error
}
