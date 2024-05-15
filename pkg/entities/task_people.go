package entities

import "gorm.io/gorm"

type TaskPeopleRole string

const (
	OWNER    TaskPeopleRole = "OWNER"
	ASSIGNEE TaskPeopleRole = "ASSIGNEE"
	REPORT   TaskPeopleRole = "REPORT"
)

type TaskPeople struct {
	gorm.Model
	TaskId int            `json:"task_id"`
	Task   Task           `json:"task"`
	UserID int            `json:"user_id"`
	User   User           `json:"user"`
	Role   TaskPeopleRole `gorm:"size:100" json:"role"`
}
