package handler

import (
	"github.com/bagusAditiaSetiawan/project-management/api/exception"
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"github.com/bagusAditiaSetiawan/project-management/pkg/project"
	"github.com/gofiber/fiber/v2"
)

type ProjectControllerImpl struct {
	ProjectService          project.ProjectService
	ProjectTransformService project.ProjectTransformService
}

func NewProjectControllerImpl(projectService project.ProjectService, projectTransformService project.ProjectTransformService) *ProjectControllerImpl {
	return &ProjectControllerImpl{
		ProjectService:          projectService,
		ProjectTransformService: projectTransformService,
	}
}

func (controller *ProjectControllerImpl) Pagination(ctx *fiber.Ctx) error {
	projectPagination := new(presenter.ProjectPaginationRequest)

	err := ctx.BodyParser(&projectPagination)
	if err != nil {
		panic(exception.NewErrorBodyException("Malformed request, please check your request"))
	}

	responsePagination := controller.ProjectService.Paginate(projectPagination)
	return ctx.JSON(controller.ProjectTransformService.ToPagination(responsePagination))
}

func (controller *ProjectControllerImpl) Create(ctx *fiber.Ctx) error {
	projectCreateRequest := new(presenter.ProjectCreateRequest)
	err := ctx.BodyParser(projectCreateRequest)
	if err != nil {
		panic(exception.NewErrorBodyException("Malformed request, please check your request"))
	}

	projectCreated := controller.ProjectService.Create(projectCreateRequest)
	return ctx.Status(fiber.StatusCreated).JSON(controller.ProjectTransformService.ToCreateProject(projectCreated))
}
func (controller *ProjectControllerImpl) Detail(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		panic(exception.NewErrorBodyException("Malformed request, please check your request"))
	}

	projectDetail := controller.ProjectService.FindById(id)
	return ctx.JSON(controller.ProjectTransformService.ToDetailProject(projectDetail))
}
