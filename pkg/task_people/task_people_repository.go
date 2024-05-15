package task_people

import (
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"github.com/bagusAditiaSetiawan/project-management/pkg/entities"
	"gorm.io/gorm"
)

type TaskPeopleRepository interface {
	Create(tx *gorm.DB, taskPeople entities.TaskPeople) entities.TaskPeople
	Paginate(tx *gorm.DB, request *presenter.TaskPeoplePaginationRequest) presenter.PaginationResponse
}
