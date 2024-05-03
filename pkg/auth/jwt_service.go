package auth

import "github.com/bagusAditiaSetiawan/project-management/pkg/entities"

type JwtService interface {
	Sign(user entities.User) string
	Verify(token string) bool
}
