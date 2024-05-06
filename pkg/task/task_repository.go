package task

import (
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"github.com/bagusAditiaSetiawan/project-management/pkg/entities"
	"gorm.io/gorm"
)

type TaskRepository interface {
	Paginate(tx *gorm.DB, request *presenter.TaskPaginationRequest) presenter.PaginationResponse
	Create(tx *gorm.DB, request *presenter.TaskCreateRequest) entities.Task
	FindById(tx *gorm.DB, id int) entities.Task
}
