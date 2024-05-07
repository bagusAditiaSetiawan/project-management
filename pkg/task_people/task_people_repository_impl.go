package task_people

import (
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
