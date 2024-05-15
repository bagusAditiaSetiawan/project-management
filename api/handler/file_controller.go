package handler

import "github.com/gofiber/fiber/v2"

type FileController interface {
	UploadFile(ctx *fiber.Ctx) error
}
