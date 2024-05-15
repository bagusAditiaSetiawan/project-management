package project

import (
	"fmt"
	"github.com/bagusAditiaSetiawan/project-management/api/exception"
	"github.com/bagusAditiaSetiawan/project-management/api/helpers"
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"github.com/bagusAditiaSetiawan/project-management/pkg/entities"
	"gorm.io/gorm"
)

type ProjectRepositoryImpl struct{}

func NewProjectRepositoryImpl() *ProjectRepositoryImpl {
	return &ProjectRepositoryImpl{}
}

func (repository *ProjectRepositoryImpl) Paginate(tx *gorm.DB, projectPagination *presenter.ProjectPaginationRequest) presenter.PaginationResponse {
	projects := []entities.Project{}
	var total int64
	query := tx.Model(projects)

	if len(projectPagination.Name) > 0 {
		query.Where("name LIKE ?", "%"+projectPagination.Name+"%")
	}
	if len(projectPagination.Description) > 0 {
		query.Where("description LIKE ?", "%"+projectPagination.Description+"%")
	}
	result := query.Count(&total)
	helpers.IfPanicHelper(result.Error)
	query.Offset(projectPagination.PaginationRequest.GetOffset()).Limit(projectPagination.Limit)
	result = query.Find(&projects)
	helpers.IfPanicHelper(result.Error)
	paginationResponse := presenter.PaginationResponse{
		List:  projects,
		Total: int(total),
	}
	return paginationResponse
}

func (repository *ProjectRepositoryImpl) FindById(tx *gorm.DB, id int) entities.Project {
	project := entities.Project{}
	result := tx.Where("id = ?", id).Preload("Tasks").First(&project)
	if result.Error != nil {
		panic(exception.NewNotFoundHandler(fmt.Sprintf("Project ID %d is not found", id)))
	}
	return project
}

func (repository *ProjectRepositoryImpl) Create(tx *gorm.DB, project entities.Project) entities.Project {
	result := tx.Create(&project)
	if result.Error != nil {
		panic(result.Error)
	}
	return project
}
