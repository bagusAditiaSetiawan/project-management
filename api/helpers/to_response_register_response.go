package helpers

import (
	"net/http"

	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"github.com/bagusAditiaSetiawan/project-management/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

func ToResponseRegisterResponse(user entities.User) *presenter.WebResponse {
	return &presenter.WebResponse{
		Code:   http.StatusCreated,
		Status: "Created",
		Data: fiber.Map{
			"username": user.Username,
			"email":    user.Email,
		},
	}
}
