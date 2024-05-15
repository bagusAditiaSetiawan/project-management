package project

import (
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"github.com/bagusAditiaSetiawan/project-management/pkg/entities"
)

type ProjectTransformService interface {
	ToCreateProject(project entities.Project) *presenter.WebResponse
	ToDetailProject(project entities.Project) *presenter.WebResponse
	ToPagination(pagination presenter.PaginationResponse) *presenter.WebResponse
}
