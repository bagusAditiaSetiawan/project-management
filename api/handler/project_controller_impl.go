package handler

import (
	"github.com/bagusAditiaSetiawan/project-management/api/exception"
	"github.com/bagusAditiaSetiawan/project-management/api/helpers"
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"github.com/bagusAditiaSetiawan/project-management/pkg/project"
	"github.com/gofiber/fiber/v2"
)

type ProjectControllerImpl struct {
	ProjectService project.ProjectService
}

func NewProjectControllerImpl(projectService project.ProjectService) *ProjectControllerImpl {
	return &ProjectControllerImpl{
		ProjectService: projectService,
	}
}

func (controller *ProjectControllerImpl) Pagination(ctx *fiber.Ctx) error {
	projectPagination := new(presenter.ProjectPaginationRequest)

	err := ctx.BodyParser(&projectPagination)
	if err != nil {
		panic(exception.NewErrorBodyException("Malformed request, please check your request"))
	}

	responsePagination := controller.ProjectService.Paginate(projectPagination)
	return ctx.JSON(helpers.ToResponsePagination(responsePagination))
}

func (controller *ProjectControllerImpl) Create(ctx *fiber.Ctx) error {
	projectCreateRequest := new(presenter.ProjectCreateRequest)
	err := ctx.BodyParser(projectCreateRequest)
	if err != nil {
		panic(exception.NewErrorBodyException("Malformed request, please check your request"))
	}

	responseCreateRequest := controller.ProjectService.Create(projectCreateRequest)
	return ctx.JSON(helpers.ToCreatedResponse(responseCreateRequest))
}
func (controller *ProjectControllerImpl) Detail(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		panic(exception.NewErrorBodyException("Malformed request, please check your request"))
	}

	responseCreateRequest := controller.ProjectService.FindById(id)
	return ctx.JSON(helpers.ToReadResponse(responseCreateRequest))
}
