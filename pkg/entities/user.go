package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"size:255;unique" json:"email"`
	Username string `gorm:"size:255;unique" json:"username"`
	Password string `json:"password"`
}
