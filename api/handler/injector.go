//go:build wireinject
// +build wireinject

package handler

import (
	"github.com/bagusAditiaSetiawan/project-management/pkg/auth"
	"github.com/bagusAditiaSetiawan/project-management/pkg/project"
	"github.com/bagusAditiaSetiawan/project-management/pkg/task"
	"github.com/bagusAditiaSetiawan/project-management/pkg/task_people"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var authSet = wire.NewSet(auth.NewUserRepositoryImpl,
	wire.Bind(new(auth.UserRepository), new(*auth.UserRepositoryImpl)),
	auth.NewPasswordServiceImpl,
	wire.Bind(new(auth.PasswordService), new(*auth.PasswordServiceImpl)),
	auth.NewJwtServiceImpl,
	wire.Bind(new(auth.JwtService), new(*auth.JwtServiceImpl)),
	auth.NewUserServiceImpl,
	wire.Bind(new(auth.UserService), new(*auth.UserServiceImpl)),
	NewAuthControllerImpl,
	wire.Bind(new(AuthController), new(*AuthControllerImpl)),
)

var projectSet = wire.NewSet(project.NewProjectRepositoryImpl,
	wire.Bind(new(project.ProjectRepository), new(*project.ProjectRepositoryImpl)),
	project.NewProjectServiceImpl,
	wire.Bind(new(project.ProjectService), new(*project.ProjectServiceImpl)),
	NewProjectControllerImpl,
	wire.Bind(new(ProjectController), new(*ProjectControllerImpl)),
)
var taskSet = wire.NewSet(task.NewTaskRepositoryImpl,
	wire.Bind(new(task.TaskRepository), new(*task.TaskRepositoryImpl)),
	project.NewProjectRepositoryImpl,
	wire.Bind(new(project.ProjectRepository), new(*project.ProjectRepositoryImpl)),
	task.NewTaskServiceImpl,
	wire.Bind(new(task.TaskService), new(*task.TaskServiceImpl)),
	NewTaskControllerImpl,
	wire.Bind(new(TaskController), new(*TaskControllerImpl)),
)

func InitializeAuthController(db *gorm.DB, validate *validator.Validate) *AuthControllerImpl {
	wire.Build(authSet)
	return nil
}

func InitializeProjectController(db *gorm.DB, validate *validator.Validate) *ProjectControllerImpl {
	wire.Build(projectSet)
	return nil
}
func InitializeTaskController(db *gorm.DB, validate *validator.Validate) *TaskControllerImpl {
	wire.Build(taskSet)
	return nil
}

var taskPeopleSet = wire.NewSet(task_people.NewTaskPeopleRepositoryImpl,
	wire.Bind(new(task_people.TaskPeopleRepository), new(*task_people.TaskPeopleRepositoryImpl)),
	task.NewTaskRepositoryImpl,
	wire.Bind(new(task.TaskRepository), new(*task.TaskRepositoryImpl)),
	task_people.NewTaskPeopleServiceImpl,
	wire.Bind(new(task_people.TaskPeopleService), new(*task_people.TaskPeopleServiceImpl)),
	auth.NewUserRepositoryImpl,
	wire.Bind(new(auth.UserRepository), new(*auth.UserRepositoryImpl)),
	NewTaskPeopleController,
	wire.Bind(new(TaskPeopleController), new(*TaskPeopleControllerImpl)),
)

func InitializeTaskPeopleController(db *gorm.DB, validate *validator.Validate) *TaskPeopleControllerImpl {
	wire.Build(taskPeopleSet)
	return nil
}
