package handler

import (
	"github.com/bagusAditiaSetiawan/project-management/api/exception"
	"github.com/bagusAditiaSetiawan/project-management/api/helpers"
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"github.com/bagusAditiaSetiawan/project-management/pkg/task_people"
	"github.com/gofiber/fiber/v2"
)

type TaskPeopleControllerImpl struct {
	TaskPeopleService task_people.TaskPeopleService
}

func NewTaskPeopleController(taskPeopleService task_people.TaskPeopleService) *TaskPeopleControllerImpl {
	return &TaskPeopleControllerImpl{
		TaskPeopleService: taskPeopleService,
	}
}

func (controller TaskPeopleControllerImpl) TaskPeopleCreate(ctx *fiber.Ctx) error {
	request := new(presenter.TaskPeopleCreateRequest)
	err := ctx.BodyParser(&request)
	if err != nil {
		panic(exception.NewErrorBodyException("Malformed request, please check your request"))
	}
	taskPeopleResponse := controller.TaskPeopleService.Create(request)
	return ctx.JSON(helpers.ToCreatedResponse(taskPeopleResponse))
}
