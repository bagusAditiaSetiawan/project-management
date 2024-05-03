//go:build wireinject
// +build wireinject

package handler

import (
	"github.com/bagusAditiaSetiawan/project-management/src/repository"
	"github.com/bagusAditiaSetiawan/project-management/src/service"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var authSet = wire.NewSet(repository.NewUserRepositoryImpl,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
	service.NewPasswordServiceImpl,
	wire.Bind(new(service.PasswordService), new(*service.PasswordServiceImpl)),
	service.NewJwtServiceImpl,
	wire.Bind(new(service.JwtService), new(*service.JwtServiceImpl)),
	service.NewUserServiceImpl,
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
	NewAuthControllerImpl,
	wire.Bind(new(AuthController), new(*AuthControllerImpl)),
)

var projectSet = wire.NewSet(repository.NewProjectRepositoryImpl,
	wire.Bind(new(repository.ProjectRepository), new(*repository.ProjectRepositoryImpl)),
	service.NewProjectServiceImpl,
	wire.Bind(new(service.ProjectService), new(*service.ProjectServiceImpl)),
	NewProjectControllerImpl,
	wire.Bind(new(ProjectController), new(*ProjectControllerImpl)),
)

func InitializeAuthController(db *gorm.DB, validate *validator.Validate) *AuthControllerImpl {
	wire.Build(authSet)
	return nil
}

func InitializeProjectController(db *gorm.DB, validate *validator.Validate) *ProjectControllerImpl {
	wire.Build(projectSet)
	return nil
}
