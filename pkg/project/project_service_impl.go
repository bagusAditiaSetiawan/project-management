package project

import (
	"github.com/bagusAditiaSetiawan/project-management/api/helpers"
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"github.com/bagusAditiaSetiawan/project-management/pkg/entities"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ProjectServiceImpl struct {
	DB                *gorm.DB
	ProjectRepository ProjectRepository
	Validator         *validator.Validate
}

func NewProjectServiceImpl(db *gorm.DB, projectRepository ProjectRepository, validator *validator.Validate) *ProjectServiceImpl {
	return &ProjectServiceImpl{
		DB:                db,
		ProjectRepository: projectRepository,
		Validator:         validator,
	}
}

func (service *ProjectServiceImpl) Create(request *presenter.ProjectCreateRequest) entities.Project {
	errorValidate := service.Validator.Struct(request)
	helpers.IfPanicHelper(errorValidate)

	tx := service.DB.Begin()

	project := entities.Project{
		Name:        request.Name,
		Description: request.Description,
		IsActive:    request.IsActive,
	}
	project = service.ProjectRepository.Create(tx, project)
	defer helpers.RollbackOrCommitDb(tx)
	return project
}

func (service *ProjectServiceImpl) Paginate(request *presenter.ProjectPaginationRequest) presenter.PaginationResponse {
	errorValidate := service.Validator.Struct(request)

	helpers.IfPanicHelper(errorValidate)

	tx := service.DB.Begin()
	paginationResponse := service.ProjectRepository.Paginate(tx, request)
	defer helpers.RollbackOrCommitDb(tx)
	return paginationResponse
}

func (service *ProjectServiceImpl) FindById(id int) entities.Project {
	tx := service.DB.Begin()

	project := service.ProjectRepository.FindById(tx, id)
	defer helpers.RollbackOrCommitDb(tx)
	return project

}
