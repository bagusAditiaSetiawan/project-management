package entities

import "gorm.io/gorm"

type StatusTasks string

const (
	BACKLOG        StatusTasks = "BACKLOG"
	READY_TO_START StatusTasks = "READY_TO_START"
	COMPLETED      StatusTasks = "COMPLETED"
	PLANNING       StatusTasks = "PLANNING"
	STUCK          StatusTasks = "STUCK"
	IN_PROGRESS    StatusTasks = "IN_PROGRESS"
	REVIEW         StatusTasks = "REVIEW"
	DONE           StatusTasks = "DONE"
)

type Task struct {
	gorm.Model
	ProjectID   int         `json:"project_id"`
	Project     Project     `json:"project"`
	Name        string      `gorm:"size:255" json:"name"`
	Description string      `json:"description"`
	Status      StatusTasks `gorm:"size:100" json:"status"`
	TaskPeoples []TaskPeople
}
