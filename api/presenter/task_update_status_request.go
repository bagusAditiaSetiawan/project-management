package presenter

import "github.com/bagusAditiaSetiawan/project-management/pkg/entities"

type TaskUpdateStatusRequest struct {
	Id     int                  `json:"id" validate:"required,number"`
	Status entities.StatusTasks `json:"status" validate:"required,oneof=BACKLOG READY_TO_START COMPLETED PLANNING STUCK REVIEW DONE IN_PROGRESS FAILED"`
}
