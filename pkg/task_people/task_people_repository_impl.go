package task_people

import (
	"github.com/bagusAditiaSetiawan/project-management/api/presenter"
	"github.com/bagusAditiaSetiawan/project-management/pkg/entities"
	"gorm.io/gorm"
)

type TaskPeopleRepositoryImpl struct{}

func NewTaskPeopleRepositoryImpl() *TaskPeopleRepositoryImpl {
	return &TaskPeopleRepositoryImpl{}
}

func (repository *TaskPeopleRepositoryImpl) Create(tx *gorm.DB, taskPeople entities.TaskPeople) entities.TaskPeople {
	result := tx.Create(&taskPeople)
	if result.Error != nil {
		panic(result.Error)
	}
	return taskPeople
}

func (repository *TaskPeopleRepositoryImpl) Paginate(tx *gorm.DB, request *presenter.TaskPeoplePaginationRequest) presenter.PaginationResponse {
	taskPeople := []entities.TaskPeople{}
	var total int64
	query := tx.Model(taskPeople)
	if request.UserID > 0 {
		query.Where("user_id = ?", request.UserID)
	}
	if request.TaskID > 0 {
		query.Where("task_id = ?", request.TaskID)
	}
	if len(request.Role) > 0 {
		query.Where("status = ?", request.Role)
	}
	query.Count(&total)
	query.Offset(request.GetOffset()).Limit(request.Limit)
	query.Find(&taskPeople)

	tasksPagination := presenter.PaginationResponse{
		List:  taskPeople,
		Total: int(total),
	}
	return tasksPagination
}
