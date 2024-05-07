package task_people

import (
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"github.com/bagusAditiaSetiawan/project-management/pkg/entities"
)

type TaskPeopleService interface {
	Create(request *presenter.TaskPeopleCreateRequest) entities.TaskPeople
}
