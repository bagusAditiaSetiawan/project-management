package project

import (
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"github.com/bagusAditiaSetiawan/project-management/pkg/entities"
)

type ProjectService interface {
	Create(request *presenter.ProjectCreateRequest) entities.Project
	Paginate(request *presenter.ProjectPaginationRequest) presenter.PaginationResponse
}
