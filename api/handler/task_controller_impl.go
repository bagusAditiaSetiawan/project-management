package handler

import (
	"github.com/bagusAditiaSetiawan/project-management/api/exception"
	"github.com/bagusAditiaSetiawan/project-management/api/helpers"
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"github.com/bagusAditiaSetiawan/project-management/pkg/task"
	"github.com/gofiber/fiber/v2"
)

type TaskControllerImpl struct {
	TaskService task.TaskService
}

func NewTaskControllerImpl(taskService task.TaskService) *TaskControllerImpl {
	return &TaskControllerImpl{
		TaskService: taskService,
	}
}

func (controller TaskControllerImpl) TaskPagination(ctx *fiber.Ctx) error {
	taskPagination := new(presenter.TaskPaginationRequest)

	if err := ctx.BodyParser(taskPagination); err != nil {
		panic(exception.NewErrorBodyException("Malformed request, please check your request"))
	}
	responsePagination := controller.TaskService.Paginate(taskPagination)
	return ctx.JSON(helpers.ToResponsePagination(responsePagination))
}

func (controller TaskControllerImpl) TaskCreate(ctx *fiber.Ctx) error {
	taskCreateRequest := new(presenter.TaskCreateRequest)
	if err := ctx.BodyParser(taskCreateRequest); err != nil {
		panic(exception.NewErrorBodyException("Malformed request, please check your request"))
	}
	taskResponse := controller.TaskService.Create(taskCreateRequest)
	return ctx.JSON(helpers.ToCreatedResponse(taskResponse))
}
