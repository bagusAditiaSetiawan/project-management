package entities

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name        string `gorm:"size:255" json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}
