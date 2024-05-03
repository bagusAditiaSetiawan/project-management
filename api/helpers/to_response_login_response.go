package helpers

import (
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func ToResponseLoginResponse(token string) *presenter.WebResponse {
	return &presenter.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data: fiber.Map{
			"access_token": token,
		},
	}
}
