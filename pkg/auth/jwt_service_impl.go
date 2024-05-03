package auth

import (
	"github.com/bagusAditiaSetiawan/project-management/api/config"
	"github.com/bagusAditiaSetiawan/project-management/api/helpers"
	"github.com/bagusAditiaSetiawan/project-management/pkg/entities"
	"github.com/golang-jwt/jwt/v5"
)

type JwtServiceImpl struct{}

func NewJwtServiceImpl() *JwtServiceImpl {
	return &JwtServiceImpl{}
}

func (service *JwtServiceImpl) Sign(user entities.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
	})
	signedString, err := token.SignedString([]byte(config.Config("JWT_SECRET")))
	helpers.IfPanicHelper(err)
	return signedString
}

func (service *JwtServiceImpl) Verify(token string) bool {
	panic("implement me")
}
