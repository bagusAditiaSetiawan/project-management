package task

import (
	"github.com/bagusAditiaSetiawan/project-management/api/helpers"
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"github.com/bagusAditiaSetiawan/project-management/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

type TaskTransformServiceImpl struct{}

func NewTaskTransformServiceImpl() *TaskTransformServiceImpl {
	return &TaskTransformServiceImpl{}
}

func (service *TaskTransformServiceImpl) PaginateTask(response presenter.PaginationResponse) *presenter.WebResponse {
	taskTransformedPagination := []fiber.Map{}
	tasks := response.List.([]entities.Task)
	for _, task := range tasks {
		transformed := fiber.Map{
			"id":          task.ID,
			"name":        task.Name,
			"description": task.Description,
			"status":      task.Status,
		}
		taskTransformedPagination = append(taskTransformedPagination, transformed)
	}
	response.List = taskTransformedPagination
	return helpers.ToResponsePagination(response)
}

func (service *TaskTransformServiceImpl) DetailTask(task entities.Task) *presenter.WebResponse {
	userTransformed := []fiber.Map{}
	for _, taskPeople := range task.TaskPeoples {
		user := fiber.Map{
			"id":       taskPeople.ID,
			"username": taskPeople.User.Username,
			"role":     taskPeople.Role,
		}
		userTransformed = append(userTransformed, user)
	}

	taskTransformed := fiber.Map{
		"id":          task.ID,
		"name":        task.Name,
		"description": task.Description,
		"status":      task.Status,
		"users":       userTransformed,
	}
	return helpers.ToReadResponse(taskTransformed)
}

func (service *TaskTransformServiceImpl) CreateTask(task entities.Task) *presenter.WebResponse {
	taskTransformed := fiber.Map{
		"id":          task.ID,
		"name":        task.Name,
		"description": task.Description,
		"status":      task.Status,
	}
	return helpers.ToCreatedResponse(taskTransformed)
}
