package presenter

import "github.com/bagusAditiaSetiawan/project-management/pkg/entities"

type TaskPeopleCreateRequest struct {
	UserID int                     `json:"user_id" validate:"required,number"`
	TaskID int                     `json:"task_id"  validate:"required,number"`
	Role   entities.TaskPeopleRole `json:"role" validate:"required,oneof=OWNER REPORTER ASSIGNEE"`
}
