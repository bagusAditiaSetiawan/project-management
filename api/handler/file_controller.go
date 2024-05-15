package handler

import "github.com/gofiber/fiber/v2"

type FileController interface {
	uploadFile(ctx *fiber.Ctx) error
}
