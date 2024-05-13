package handler

import (
	"github.com/bagusAditiaSetiawan/project-management/api/exception"
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"github.com/bagusAditiaSetiawan/project-management/pkg/task"
	"github.com/gofiber/fiber/v2"
)

type TaskControllerImpl struct {
	TaskService          task.TaskService
	TaskTransformService task.TaskTransformService
}

func NewTaskControllerImpl(taskService task.TaskService, taskTransformService task.TaskTransformService) *TaskControllerImpl {
	return &TaskControllerImpl{
		TaskService:          taskService,
		TaskTransformService: taskTransformService,
	}
}

func (controller TaskControllerImpl) TaskPagination(ctx *fiber.Ctx) error {
	taskPagination := new(presenter.TaskPaginationRequest)

	if err := ctx.BodyParser(taskPagination); err != nil {
		panic(exception.NewErrorBodyException("Malformed request, please check your request"))
	}
	responsePagination := controller.TaskService.Paginate(taskPagination)
	return ctx.JSON(controller.TaskTransformService.PaginateTask(responsePagination))
}

func (controller TaskControllerImpl) TaskCreate(ctx *fiber.Ctx) error {
	taskCreateRequest := new(presenter.TaskCreateRequest)
	if err := ctx.BodyParser(taskCreateRequest); err != nil {
		panic(exception.NewErrorBodyException("Malformed request, please check your request"))
	}
	taskResponse := controller.TaskService.Create(taskCreateRequest)
	return ctx.Status(fiber.StatusCreated).JSON(controller.TaskTransformService.CreateTask(taskResponse))
}

func (controller TaskControllerImpl) TaskDetail(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		panic(exception.NewErrorBodyException("Malformed request, please check your request"))
	}
	taskData := controller.TaskService.FindById(id)
	return ctx.JSON(controller.TaskTransformService.DetailTask(taskData))
}
