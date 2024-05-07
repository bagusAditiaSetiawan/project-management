package task_people

import (
	"github.com/bagusAditiaSetiawan/project-management/api/helpers"
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"github.com/bagusAditiaSetiawan/project-management/pkg/auth"
	"github.com/bagusAditiaSetiawan/project-management/pkg/entities"
	"github.com/bagusAditiaSetiawan/project-management/pkg/task"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type TaskPeopleServiceImpl struct {
	DB                   *gorm.DB
	UserRepository       auth.UserRepository
	TaskPeopleRepository TaskPeopleRepository
	TaskRepository       task.TaskRepository
	Validation           *validator.Validate
}

func NewTaskPeopleServiceImpl(db *gorm.DB, userRepository auth.UserRepository, taskPeopleRepository TaskPeopleRepository, taskRepository task.TaskRepository, validation *validator.Validate) *TaskPeopleServiceImpl {
	return &TaskPeopleServiceImpl{
		DB:                   db,
		UserRepository:       userRepository,
		TaskPeopleRepository: taskPeopleRepository,
		TaskRepository:       taskRepository,
		Validation:           validation,
	}
}

func (service *TaskPeopleServiceImpl) Create(request *presenter.TaskPeopleCreateRequest) entities.TaskPeople {
	err := service.Validation.Struct(request)
	helpers.IfPanicHelper(err)

	tx := service.DB.Begin()
	service.TaskRepository.FindById(tx, request.TaskID)
	service.UserRepository.FindByIdOrFail(tx, request.UserID)
	taskPeople := entities.TaskPeople{
		UserID: request.UserID,
		TaskId: request.TaskID,
		Role:   request.Role,
	}
	taskPeople = service.TaskPeopleRepository.Create(tx, taskPeople)

	defer helpers.RollbackOrCommitDb(tx)
	return taskPeople
}
