package auth

import (
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"github.com/bagusAditiaSetiawan/project-management/pkg/entities"
	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	Register(request *presenter.RegisterUserRequest) entities.User
	Login(request *presenter.LoginUserRequest) string
	ClaimsUserToken(token *jwt.Token) entities.User
}
