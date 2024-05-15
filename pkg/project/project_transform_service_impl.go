package project

import (
	"github.com/bagusAditiaSetiawan/project-management/api/helpers"
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"github.com/bagusAditiaSetiawan/project-management/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

type ProjectTransformServiceImpl struct{}

func NewProjectTransformServiceImpl() *ProjectTransformServiceImpl {
	return &ProjectTransformServiceImpl{}
}

func (service *ProjectTransformServiceImpl) ToCreateProject(project entities.Project) *presenter.WebResponse {
	transformed := fiber.Map{
		"id":          project.ID,
		"name":        project.Name,
		"description": project.Description,
	}
	return helpers.ToCreatedResponse(transformed)
}

func (service *ProjectTransformServiceImpl) ToDetailProject(project entities.Project) *presenter.WebResponse {
	transformed := fiber.Map{
		"id":          project.ID,
		"name":        project.Name,
		"description": project.Description,
		"tasks":       project.Tasks,
	}
	return helpers.ToReadResponse(transformed)
}

func (service *ProjectTransformServiceImpl) ToPagination(pagination presenter.PaginationResponse) *presenter.WebResponse {
	projects := pagination.List.([]entities.Project)
	transformed := []fiber.Map{}
	for _, project := range projects {
		data := fiber.Map{
			"id":          project.ID,
			"name":        project.Name,
			"description": project.Description,
		}
		transformed = append(transformed, data)
	}
	pagination.List = transformed
	return helpers.ToResponsePagination(pagination)
}
