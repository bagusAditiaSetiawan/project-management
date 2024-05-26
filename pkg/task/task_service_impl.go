package task

import (
	"github.com/bagusAditiaSetiawan/project-management/api/helpers"
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"github.com/bagusAditiaSetiawan/project-management/pkg/aws_cloudwatch"
	"github.com/bagusAditiaSetiawan/project-management/pkg/entities"
	"github.com/bagusAditiaSetiawan/project-management/pkg/project"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type TaskServiceImpl struct {
	DB                *gorm.DB
	Validator         *validator.Validate
	TaskRepository    TaskRepository
	ProjectRepository project.ProjectRepository
	Logger            *aws_cloudwatch.AwsCloudWatchServiceImpl
}

func NewTaskServiceImpl(db *gorm.DB, validator *validator.Validate, taskRepository TaskRepository, projectRepository project.ProjectRepository, logger *aws_cloudwatch.AwsCloudWatchServiceImpl) *TaskServiceImpl {
	return &TaskServiceImpl{
		DB:                db,
		Validator:         validator,
		TaskRepository:    taskRepository,
		ProjectRepository: projectRepository,
		Logger:            logger,
	}
}

func (service *TaskServiceImpl) Create(request *presenter.TaskCreateRequest) entities.Task {
	err := service.Validator.Struct(request)
	helpers.IfPanicHelper(err)
	service.Logger.SendLogInfo("Process create task request", "request", request)
	tx := service.DB.Begin()
	service.ProjectRepository.FindById(tx, request.ProjectId)
	task := service.TaskRepository.Create(tx, request)
	helpers.RollbackOrCommitDb(tx)
	return task
}

func (service *TaskServiceImpl) Paginate(request *presenter.TaskPaginationRequest) presenter.PaginationResponse {
	err := service.Validator.Struct(request)
	helpers.IfPanicHelper(err)
	service.Logger.SendLogInfo("Process paginate task request", "request", request)
	tx := service.DB.Begin()
	taskResponse := service.TaskRepository.Paginate(tx, request)
	helpers.RollbackOrCommitDb(tx)
	return taskResponse
}

func (service *TaskServiceImpl) FindById(id int) entities.Task {
	tx := service.DB.Begin()
	service.Logger.SendLogInfo("Process find task request", "id", id)
	task := service.TaskRepository.FindById(tx, id)
	helpers.RollbackOrCommitDb(tx)
	return task
}

func (service *TaskServiceImpl) UpdateStatus(request *presenter.TaskUpdateStatusRequest) entities.Task {
	err := service.Validator.Struct(request)
	helpers.IfPanicHelper(err)
	tx := service.DB.Begin()
	service.Logger.SendLogInfo("Process update task request", "id", request.Id)
	task := service.TaskRepository.UpdateStatus(tx, request.Id, request.Status)
	helpers.RollbackOrCommitDb(tx)
	return task
}
