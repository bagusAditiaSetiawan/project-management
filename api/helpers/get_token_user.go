package helpers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GetTokenUser(ctx *fiber.Ctx) *jwt.Token {
	return ctx.Locals("auth").(*jwt.Token)
}
