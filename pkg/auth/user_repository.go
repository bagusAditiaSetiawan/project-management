package auth

import (
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"github.com/bagusAditiaSetiawan/project-management/pkg/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	Register(tx *gorm.DB, user *presenter.RegisterUserRequest) (entities.User, error)
	FindByEmailUsername(tx *gorm.DB, email string, username string) entities.User
	FindById(tx *gorm.DB, id int) entities.User
}
