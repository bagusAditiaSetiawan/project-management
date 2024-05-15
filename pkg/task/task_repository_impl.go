package task

import (
	"fmt"
	"github.com/bagusAditiaSetiawan/project-management/api/exception"
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"github.com/bagusAditiaSetiawan/project-management/pkg/entities"
	"gorm.io/gorm"
)

type TaskRepositoryImpl struct{}

func NewTaskRepositoryImpl() *TaskRepositoryImpl {
	return &TaskRepositoryImpl{}
}

func (repository *TaskRepositoryImpl) Paginate(tx *gorm.DB, request *presenter.TaskPaginationRequest) presenter.PaginationResponse {
	tasks := []entities.Task{}
	var total int64
	query := tx.Model(tasks)
	if len(request.Name) > 0 {
		query.Where("name = ?", "%"+request.Name+"%")
	}

	if len(request.Description) > 0 {
		query.Where("description LIKE ?", "%"+request.Description+"%")
	}
	if len(request.Status) > 0 {
		query.Where("status = ?", request.Status)
	}
	query.Count(&total)
	query.Offset(request.GetOffset()).Limit(request.Limit)
	query.Find(&tasks)

	tasksPagination := presenter.PaginationResponse{
		List:  tasks,
		Total: int(total),
	}
	return tasksPagination
}

func (repository *TaskRepositoryImpl) Create(tx *gorm.DB, request *presenter.TaskCreateRequest) entities.Task {
	task := entities.Task{
		Name:        request.Name,
		Description: request.Description,
		ProjectID:   request.ProjectId,
		Status:      request.Status,
	}
	result := tx.Create(&task)
	if result.Error != nil {
		panic("Create task is failed")
	}
	return task
}

func (repository *TaskRepositoryImpl) FindById(tx *gorm.DB, id int) entities.Task {
	task := entities.Task{}
	result := tx.Where("id = ? ", id).Preload("TaskPeoples").Preload("TaskPeoples.User").Preload("Project").First(&task)
	if result.Error != nil {
		panic(exception.NewNotFoundHandler(fmt.Sprintf("Tasks ID  %d is not found", id)))
	}
	return task
}
