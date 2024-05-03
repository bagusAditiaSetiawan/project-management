package project

import (
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"github.com/bagusAditiaSetiawan/project-management/pkg/entities"
	"gorm.io/gorm"
)

type ProjectRepository interface {
	Paginate(tx *gorm.DB, request *presenter.ProjectPaginationRequest) presenter.PaginationResponse
	FindById(tx *gorm.DB, id int) entities.Project
	Create(tx *gorm.DB, project entities.Project) entities.Project
}
