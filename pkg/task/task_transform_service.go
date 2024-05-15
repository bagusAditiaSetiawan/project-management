package task

import (
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"github.com/bagusAditiaSetiawan/project-management/pkg/entities"
)

type TaskTransformService interface {
	PaginateTask(response presenter.PaginationResponse) *presenter.WebResponse
	DetailTask(task entities.Task) *presenter.WebResponse
	CreateTask(task entities.Task) *presenter.WebResponse
}
