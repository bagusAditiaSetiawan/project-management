package auth

import (
	"errors"
	"github.com/bagusAditiaSetiawan/project-management/api/helpers"
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"github.com/bagusAditiaSetiawan/project-management/pkg/entities"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct{}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Register(tx *gorm.DB, registerRequest *presenter.RegisterUserRequest) (entities.User, error) {
	userExist := repository.FindByEmailUsername(tx, registerRequest.Email, registerRequest.Username)
	if userExist.ID != 0 {
		return userExist, errors.New("Username or Email is already exist")
	}
	user := entities.User{
		Username: registerRequest.Username,
		Password: registerRequest.Password,
		Email:    registerRequest.Email,
	}
	result := tx.Create(&user)
	helpers.IfPanicHelper(result.Error)

	helpers.RollbackOrCommitDb(tx)
	return user, nil
}

func (repository *UserRepositoryImpl) FindByEmailUsername(tx *gorm.DB, email string, username string) entities.User {
	user := entities.User{}
	tx.Where("email = ? OR username >= ?", email, username).First(&user)
	return user
}

func (repository *UserRepositoryImpl) FindById(tx *gorm.DB, id int) entities.User {
	user := entities.User{}
	tx.Where("id = ?", id).First(&user)
	return user
}
