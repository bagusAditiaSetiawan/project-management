package auth

import (
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"github.com/bagusAditiaSetiawan/project-management/pkg/entities"
)

type UserTransformService interface {
	ToResponseLoginResponse(token string) *presenter.WebResponse
	ToResponseRegisterResponse(user entities.User) *presenter.WebResponse
}
