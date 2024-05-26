package task

import (
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"github.com/bagusAditiaSetiawan/project-management/pkg/entities"
)

type TaskService interface {
	Create(request *presenter.TaskCreateRequest) entities.Task
	Paginate(request *presenter.TaskPaginationRequest) presenter.PaginationResponse
	FindById(id int) entities.Task
	UpdateStatus(request *presenter.TaskUpdateStatusRequest) entities.Task
}
