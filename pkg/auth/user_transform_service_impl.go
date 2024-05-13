package auth

import (
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"github.com/bagusAditiaSetiawan/project-management/pkg/entities"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type UserTransformServiceImpl struct{}

func NewUserTransformServiceImpl() *UserTransformServiceImpl {
	return &UserTransformServiceImpl{}
}

func (service *UserTransformServiceImpl) ToResponseLoginResponse(token string) *presenter.WebResponse {
	return &presenter.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data: fiber.Map{
			"access_token": token,
		},
	}
}

func (service *UserTransformServiceImpl) ToResponseRegisterResponse(user entities.User) *presenter.WebResponse {
	return &presenter.WebResponse{
		Code:   http.StatusCreated,
		Status: "Created",
		Data: fiber.Map{
			"username": user.Username,
			"email":    user.Email,
		},
	}
}
