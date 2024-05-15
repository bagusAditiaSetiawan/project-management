package presenter

import "github.com/bagusAditiaSetiawan/project-management/pkg/entities"

type TaskCreateRequest struct {
	ProjectId   int                  `json:"project_id" validate:"required,number"`
	Name        string               `json:"name" validate:"required,max=255"`
	Description string               `json:"description" validate:"required"`
	Status      entities.StatusTasks `json:"status" validate:"required,oneof=BACKLOG READY_TO_START COMPLETED PLANNING STUCK REVIEW DONE"`
}
